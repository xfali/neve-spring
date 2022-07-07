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
	"net/http"
	"os"
	"sigs.k8s.io/controller-tools/pkg/markers"
	"strings"

	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
)

type Method struct {
	Name           string
	RequestMapping RequestMappingMarker
	RequestParam   RequestParamMarker
	RequestHeader  RequestHeaderMarker
	RequestBody    RequestBodyMarker
}

type GinMetadata struct {
	Name       string
	Controller ControllerMarker
	Methods    []*Method
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

func (p *ginPlugin) parseType(t *types.Type) (*GinMetadata, error) {
	method := t.Methods
	if len(method) == 0 {
		return nil, fmt.Errorf("Type: %s without method. ", t.Name)
	}

	ret := &GinMetadata{
		Name: t.Name.Name,
	}
	comments := MergeComments(t)
	for _, c := range comments {
		if c == "" {
			continue
		}
		set, err := markerdefs.Parse(c, &ret.Controller)
		if err != nil {
			return nil, err
		}
		if !set {
			return nil, nil
		} else {
			break
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

			if !m.RequestParam.IsSet() {
				set, err := markerdefs.Parse(c, &m.RequestParam)
				if err != nil {
					return nil, err
				} else if set {
					m.RequestParam.Set(true)
					continue
				}
			}
			if !m.RequestHeader.IsSet() {
				set, err := markerdefs.Parse(c, &m.RequestHeader)
				if err != nil {
					return nil, err
				} else if set {
					m.RequestHeader.Set(true)
					continue
				}
			}

			if !m.RequestBody.IsSet() {
				set, err := markerdefs.Parse(c, &m.RequestBody)
				if err != nil {
					return nil, err
				} else if set {
					m.RequestBody.Set(true)
					continue
				}
			}
		}
	}

	return ret, nil
}

func (p *ginPlugin) Name() string {
	return "web:gin"
}

func (p *ginPlugin) Generate(ctx *generator.Context, w io.Writer, t *types.Type) (err error) {
	meta, err := p.parseType(t)
	if err != nil {
		return err
	}
	if meta == nil {
		// not set
		return nil
	}
	w = io.MultiWriter(w, os.Stderr)
	sw := generator.NewSnippetWriter(w, ctx, delimiterLeft, delimiterRight)
	sw.Do(p.template, meta)
	return sw.Error()
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
	Method          string `marker:"method"`
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
	markerdefs.Flag `marker:",optional"`
	Name            string `marker:"name"`
	Default         string `marker:"default,optional"`
	Require         bool   `marker:"require,optional"`
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
	markerdefs.Flag `marker:",optional"`
	Name            string `marker:"name"`
	Default         string `marker:"default,optional"`
	Require         bool   `marker:"require,optional"`
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
