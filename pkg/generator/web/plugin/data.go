/*
 * Copyright (C) 2019-2022, Xiongfa Li.
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package plugin

import (
	"fmt"
	"github.com/xfali/neve-spring/pkg/generator/markerdefs"
	"io"
	"io/ioutil"
	"k8s.io/gengo/namer"
	"net/http"
	"os"
	"runtime"
	"sigs.k8s.io/controller-tools/pkg/markers"
	"strings"
	"text/template"

	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
)

type RequestType string

const (
	RequestTypeQuery  = "query"
	RequestTypeHeader = "header"
	RequestTypeBody   = "body"
)

type TypeMeta struct {
	Name        string
	Type        string
	Default     string
	Require     bool
	RequestType RequestType
}

type Method struct {
	Name           string
	Params         []TypeMeta
	Return         []TypeMeta
	RequestMapping RequestMappingMarker
}

type GinMetadata struct {
	Name           string
	TypeName       string
	Controller     ControllerMarker
	RequestMapping RequestMappingMarker
	Methods        []*Method
}

type ginPlugin struct {
	template         string
	annotation       string
	structAnnotation string
	methodAnnotation []string
}

func NewGinPlugin(annotation string) *ginPlugin {
	tmpl, err := ioutil.ReadFile("../pkg/generator/web/template/gin.tmpl")
	if err != nil {
		panic(err)
	}
	return &ginPlugin{
		template:         string(tmpl),
		annotation:       annotation,
		structAnnotation: annotation + "controller",
		methodAnnotation: []string{
			annotation + "requestmapping",
		},
	}
}

func (p *ginPlugin) Annotation() string {
	return p.annotation
}

func MergeComments(t *types.Type) []string {
	ret := make([]string, 0, len(t.CommentLines)+len(t.SecondClosestCommentLines))
	ret = append(ret, t.CommentLines...)
	ret = append(ret, t.SecondClosestCommentLines...)
	return ret
}

func (p *ginPlugin) CouldHandle(t *types.Type) bool {
	if t.Kind == types.Struct {
		comments := MergeComments(t)
		for _, c := range comments {
			i := strings.Index(c, p.structAnnotation)
			if i != -1 {
				return true
			}
		}
	}
	return false
}

func (p *ginPlugin) parseType(imports namer.ImportTracker, t *types.Type) (*GinMetadata, error) {
	method := t.Methods
	if len(method) == 0 {
		return nil, fmt.Errorf("Type: %s without method. ", t.Name)
	}
	imports.AddType(t)
	ret := &GinMetadata{
		Name:     t.Name.Name,
		TypeName: imports.LocalNameOf(t.Name.Package) + "." + t.Name.Name,
	}
	comments := MergeComments(t)
	for _, c := range comments {
		if c == "" {
			continue
		}
		set, err := markerdefs.Parse(c, &ret.Controller)
		if err != nil {
			return nil, err
		} else if set {
			continue
		}
		set, err = markerdefs.Parse(c, &ret.RequestMapping)
		if err != nil {
			return nil, err
		} else if set {
			continue
		}
	}
	for mname, mtype := range t.Methods {
		m := &Method{}
		m.Name = mname
		comments := MergeComments(mtype)
		for _, c := range comments {
			if c == "" {
				continue
			}
			if !m.RequestMapping.IsSet() {
				set, err := markerdefs.Parse(c, &m.RequestMapping)
				if err != nil {
					return nil, err
				} else if set {
					m.RequestMapping.Set(true)
					if m.RequestMapping.Method == "" {
						m.RequestMapping.Method = http.MethodGet
					}
					ret.Methods = append(ret.Methods, m)
					continue
				}
			}

			paramMarker := RequestParamMarker{}
			set, err := markerdefs.Parse(c, &paramMarker)
			if err != nil {
				return nil, err
			} else if set {
				meta, have := findParam(imports, mtype, paramMarker.Name)
				if have {
					meta.Default = paramMarker.Default
					meta.Require = paramMarker.Require
					meta.RequestType = RequestTypeQuery
					m.Params = append(m.Params, meta)
				}
				continue
			}

			headerMarker := RequestHeaderMarker{}
			set, err = markerdefs.Parse(c, &headerMarker)
			if err != nil {
				return nil, err
			} else if set {
				meta, have := findParam(imports, mtype, headerMarker.Name)
				if have {
					meta.Default = headerMarker.Default
					meta.Require = headerMarker.Require
					meta.RequestType = RequestTypeHeader
					m.Params = append(m.Params, meta)
				}
				continue
			}

			bodyMarker := RequestBodyMarker{}
			set, err = markerdefs.Parse(c, &bodyMarker)
			if err != nil {
				return nil, err
			} else if set {
				meta, have := findParam(imports, mtype, bodyMarker.Name)
				if have {
					meta.Require = bodyMarker.Require
					meta.RequestType = RequestTypeBody
					m.Params = append(m.Params, meta)
				}
				continue
			}
		}
	}

	return ret, nil
}

func findParam(imports namer.ImportTracker, t *types.Type, name string) (TypeMeta, bool) {
	ret := TypeMeta{}
	for i, v := range t.Signature.ParameterNames {
		if v == name {
			param := t.Signature.Parameters[i]
			if param.Kind == types.Struct {
				imports.AddType(param)
			}
			ret.Name = name
			ret.Type = param.Name.Name
			return ret, true
		}
	}
	return ret, false
}

func (p *ginPlugin) Name() string {
	return "web:gin"
}

func concatUrl(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	if s2 == "" {
		return s1
	}
	if s1[len(s1)-1:] != "/" && s2[:1] != "/" {
		return s1 + "/" + s2
	}
	return s1 + s2
}

func (p *ginPlugin) Generate(ctx *generator.Context, imports namer.ImportTracker, w io.Writer, t *types.Type) (err error) {
	meta, err := p.parseType(imports, t)
	if err != nil {
		return err
	}
	if meta == nil {
		// not set
		return nil
	}
	w = io.MultiWriter(w, os.Stderr)

	funcMap := template.FuncMap{
		"concatUrl": concatUrl,
	}
	for name, namer := range ctx.Namers {
		funcMap[name] = namer.Name
	}
	_, file, line, _ := runtime.Caller(1)
	tmpl, err := template.
		New(fmt.Sprintf("%s:%d", file, line)).
		Funcs(funcMap).
		Parse(p.template)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, meta)
	//sw := generator.NewSnippetWriter(w, ctx, delimiterLeft, delimiterRight)
	//sw.Do(p.template, meta)
	//return sw.Error()
}

type ControllerMarker struct {
	Value string `marker:"value"`
}

func (ControllerMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "Controller",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Enable type controller function.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

type RequestMappingMarker struct {
	markerdefs.Flag `marker:",optional"`
	Value           string `marker:"value"`
	Method          string `marker:"method,optional"`
}

func (RequestMappingMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "RequestMapping",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define the controller request router path and method.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

type RequestParamMarker struct {
	Name    string `marker:"name"`
	Default string `marker:"default,optional"`
	Require bool   `marker:"require,optional"`
}

func (RequestParamMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "RequestParam",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define RequestParam.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

type RequestHeaderMarker struct {
	Name    string `marker:"name"`
	Default string `marker:"default,optional"`
	Require bool   `marker:"require,optional"`
}

func (RequestHeaderMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "RequestHeader",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define RequestHeader.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

type RequestBodyMarker struct {
	markerdefs.Flag `marker:",optional"`
	Name            string `marker:"name"`
	Require         bool   `marker:"require,optional"`
}

func (RequestBodyMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "RequestBody",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define RequestBody.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func init() {
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:controller", markers.DescribesType, ControllerMarker{})).
		WithHelp(ControllerMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:requestmapping", markers.DescribesType, RequestMappingMarker{})).
		WithHelp(RequestMappingMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:requestparam", markers.DescribesType, RequestParamMarker{})).
		WithHelp(RequestParamMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:requestheader", markers.DescribesType, RequestHeaderMarker{})).
		WithHelp(RequestHeaderMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:requestbody", markers.DescribesType, RequestBodyMarker{})).
		WithHelp(RequestBodyMarker{}.Help()))

}
