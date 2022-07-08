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
	"k8s.io/gengo/namer"
	"net/http"
	"os"
	"runtime"
	"text/template"

	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
)

type RequestType string

const (
	RequestTypePath   = "path"
	RequestTypeQuery  = "query"
	RequestTypeHeader = "header"
	RequestTypeBody   = "body"
)

var templateMap = map[string]string{
	http.MethodGet:    "gin_get.tmpl",
	http.MethodPost:   "gin_post.tmpl",
	http.MethodDelete: "gin_delete.tmpl",
	http.MethodPut:    "gin_put.tmpl",
}

type TypeMeta struct {
	Name        string
	TypeName    string
	Default     string
	Required    bool
	RequestType RequestType
}

type Method struct {
	Name       string
	BeanMarker *BeanMarker
	Params     []*TypeMeta
	Returns    []*TypeMeta
}

type CoreMetadata struct {
	Name             string
	TypeName         string
	ControllerMarker *ControllerMarker
	ServiceMarker    *ServiceMarker
	ComponentMarker  *ComponentMarker
	Methods          []*Method
}

type corePlugin struct {
	template   string
	annotation string
	fillFunc   AutoFillParamFunc
}

type AutoFillParamFunc func(imports namer.ImportTracker, name string, param *types.Type) (*TypeMeta, error)

type Opt func(*corePlugin)

func NewCorePlugin(annotation string, opts ...Opt) *corePlugin {
	ret := &corePlugin{
		template:   "", //getBuildTemplate("coregin.tmpl"),
		annotation: annotation,
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

func (p *corePlugin) Annotation() string {
	return p.annotation
}

func MergeComments(t *types.Type) []string {
	ret := make([]string, 0, len(t.CommentLines)+len(t.SecondClosestCommentLines))
	ret = append(ret, t.CommentLines...)
	ret = append(ret, t.SecondClosestCommentLines...)
	return ret
}

func (p *corePlugin) CouldHandle(t *types.Type) bool {
	switch t.Kind {
	case types.Struct, types.DeclarationOf:
		return true
	}
	return false
}

func (p *corePlugin) parseType(imports namer.ImportTracker, t *types.Type) (*CoreMetadata, error) {
	ret := &CoreMetadata{
		Name:     t.Name.Name,
		TypeName: imports.LocalNameOf(t.Name.Package) + "." + t.Name.Name,
	}
	comments := MergeComments(t)
	beanFound := false
	for _, c := range comments {
		if c == "" {
			continue
		}
		controllerMarker := ControllerMarker{}
		set, err := markerdefs.Parse(c, &controllerMarker)
		if err != nil {
			return nil, err
		} else if set {
			beanFound = true
			imports.AddType(t)
			ret.ControllerMarker = &controllerMarker
			break
		}

		serviceMarker := ServiceMarker{}
		set, err = markerdefs.Parse(c, &serviceMarker)
		if err != nil {
			return nil, err
		} else if set {
			beanFound = true
			imports.AddType(t)
			ret.ServiceMarker = &serviceMarker
			break
		}

		componentMarker := ComponentMarker{}
		set, err = markerdefs.Parse(c, &componentMarker)
		if err != nil {
			return nil, err
		} else if set {
			beanFound = true
			imports.AddType(t)
			ret.ComponentMarker = &componentMarker
			break
		}
	}
	if !beanFound {
		return nil, nil
	}
	for mname, mtype := range t.Methods {
		m := &Method{}
		m.Name = mname
		comments := MergeComments(mtype)
		for _, c := range comments {
			if c == "" {
				continue
			}

			beanMarker := BeanMarker{}
			set, err := markerdefs.Parse(c, &beanMarker)
			if err != nil {
				return nil, err
			} else if set {
				ret.Methods = append(ret.Methods, m)
				continue
			}
		}
		m.Returns = findResult(imports, mtype)
	}

	return ret, nil
}

func findParam(imports namer.ImportTracker, t *types.Type, name string) (*TypeMeta, bool) {
	ret := &TypeMeta{}
	for i, v := range t.Signature.ParameterNames {
		if v == name {
			param := t.Signature.Parameters[i]
			ret.Name = name
			if param.Kind == types.Struct || param.Kind == types.Interface {
				imports.AddType(param)
				ret.TypeName = imports.LocalNameOf(param.Name.Package) + "." + param.Name.Name
			} else {
				ret.TypeName = param.Name.Name
			}

			return ret, true
		}
	}
	return ret, false
}

func findResult(imports namer.ImportTracker, t *types.Type) []*TypeMeta {
	ret := make([]*TypeMeta, len(t.Signature.Results))
	for i, v := range t.Signature.Results {
		meta := &TypeMeta{}
		if v.Kind == types.Struct || v.Kind == types.Interface {
			imports.AddType(v)
			meta.TypeName = imports.LocalNameOf(v.Name.Package) + "." + v.Name.Name
		} else {
			meta.TypeName = v.Name.Name
		}
		ret[i] = meta
	}
	return ret
}

func (p *corePlugin) Name() string {
	return "web:gin"
}

func add(a, b int) int {
	return a + b
}

func (p *corePlugin) Generate(ctx *generator.Context, imports namer.ImportTracker, w io.Writer, t *types.Type) (err error) {
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
		"add": add,
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

func SetAutoFillParamFunc(f AutoFillParamFunc) Opt {
	return func(plugin *corePlugin) {
		plugin.fillFunc = f
	}
}