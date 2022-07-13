/*
 * Copyright 2022 Xiongfa Li.
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
	"github.com/xfali/neve-spring/pkg/stringfunc"
	"io"
	"k8s.io/gengo/namer"
	"net/http"
	"runtime"
	"strings"
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
	BeanMarker       *BeanMarker
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
		template:   getBuildTemplate("core.tmpl"),
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
		Name: t.Name.Name,
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
			if !stringfunc.IsFirstUpper(t.Name.Name) {
				return nil, fmt.Errorf("Type %s is private. ", t.Name)
			}
			beanFound = true
			imports.AddType(t)
			ret.ControllerMarker = &controllerMarker
			ret.TypeName = imports.LocalNameOf(t.Name.Package) + "." + t.Name.Name
			break
		}

		serviceMarker := ServiceMarker{}
		set, err = markerdefs.Parse(c, &serviceMarker)
		if err != nil {
			return nil, err
		} else if set {
			if !stringfunc.IsFirstUpper(t.Name.Name) {
				return nil, fmt.Errorf("Type %s is private ", t.Name)
			}
			beanFound = true
			imports.AddType(t)
			ret.ServiceMarker = &serviceMarker
			ret.TypeName = imports.LocalNameOf(t.Name.Package) + "." + t.Name.Name
			break
		}

		componentMarker := ComponentMarker{}
		set, err = markerdefs.Parse(c, &componentMarker)
		if err != nil {
			return nil, err
		} else if set {
			if !stringfunc.IsFirstUpper(t.Name.Name) {
				return nil, fmt.Errorf("Type %s is private ", t.Name)
			}
			beanFound = true
			imports.AddType(t)
			ret.ComponentMarker = &componentMarker
			ret.TypeName = imports.LocalNameOf(t.Name.Package) + "." + t.Name.Name
			break
		}
		beanMarker := BeanMarker{}
		set, err = markerdefs.Parse(c, &beanMarker)
		if err != nil {
			return nil, err
		} else if set {
			if !stringfunc.IsFirstUpper(t.Name.Name) {
				return nil, fmt.Errorf("Function %s is private ", t.Name)
			}
			beanFound = true
			imports.AddType(t)
			ret.BeanMarker = &beanMarker
			ret.TypeName = imports.LocalNameOf(t.Name.Package) + "." + t.Name.Name
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
				if !stringfunc.IsFirstUpper(mname) {
					return nil, fmt.Errorf("Method %s.%s is private ", t.Name, mname)
				}
				m.BeanMarker = &beanMarker
				ret.Methods = append(ret.Methods, m)
				continue
			}
		}
		m.Returns = findResult(imports, mtype)
	}

	return ret, nil
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
	return "core"
}

func add(a, b int) int {
	return a + b
}

func concat(strs ...string) string {
	buf := strings.Builder{}
	for _, s := range strs {
		buf.WriteString(s)
	}
	return buf.String()
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
	//w = io.MultiWriter(w, os.Stderr)

	funcMap := template.FuncMap{
		"add":    add,
		"concat": concat,
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

func (p *corePlugin) Finalize(ctx *generator.Context, imports namer.ImportTracker, w io.Writer) error {
	return nil
}

func SetAutoFillParamFunc(f AutoFillParamFunc) Opt {
	return func(plugin *corePlugin) {
		plugin.fillFunc = f
	}
}
