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
	"github.com/xfali/neve-spring/pkg/typeutil"
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

const (
	ScopeTypePrototype = "prototype"
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
	Name            string
	BeanMarker      *BeanMarker
	ScopeMarker     *ScopeMarker
	AutowiredMarker *AutowiredMarker
	Params          []*TypeMeta
	Returns         []*TypeMeta
}

type Field struct {
	TypeMeta
	AutowiredMarker *AutowiredMarker
}

type CoreMetadata struct {
	Name                string
	TypeName            string
	ControllerMarker    *ControllerMarker
	ServiceMarker       *ServiceMarker
	ComponentMarker     *ComponentMarker
	BeanMarker          *BeanMarker
	ScopeMarker         *ScopeMarker
	PostConstructMarker *PostConstructMarker
	PreDestroyMarker    *PreDestroyMarker
	Fields              []*Field
	Methods             []*Method
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
		scopeMarker := ScopeMarker{}
		set, err := markerdefs.Parse(c, &scopeMarker)
		if err != nil {
			return nil, err
		} else if set {
			ret.ScopeMarker = &scopeMarker
			continue
		}

		if !beanFound {
			controllerMarker := ControllerMarker{}
			set, err = markerdefs.Parse(c, &controllerMarker)
			if err != nil {
				return nil, err
			} else if set {
				if !stringfunc.IsFirstUpper(t.Name.Name) {
					return nil, fmt.Errorf("Type %s is private ", t.Name)
				}
				beanFound = true
				imports.AddType(t)
				ret.ControllerMarker = &controllerMarker
				ret.TypeName = imports.LocalNameOf(t.Name.Package) + "." + t.Name.Name
				continue
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
				continue
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
				continue
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
				continue
			}
		}
	}
	if !beanFound {
		return nil, nil
	}

	for _, member := range t.Members {
		m := &Field{}
		mname, mtype := member.Name, member.Type
		m.Name = mname
		for _, c := range member.CommentLines {
			if c == "" {
				continue
			}

			autowiredMarker := AutowiredMarker{}
			set, err := markerdefs.Parse(c, &autowiredMarker)
			if err != nil {
				return nil, err
			} else if set {
				//if ret.ScopeMarker != nil && ret.ScopeMarker.Value == ScopeTypePrototype {
				//	return nil, fmt.Errorf("Type %s is prototype, field cannot with [autowired] annotation ", t.Name)
				//}
				if !stringfunc.IsFirstUpper(mname) {
					return nil, fmt.Errorf("Type %s field %s is private ", t.Name, mname)
				}
				m.AutowiredMarker = &autowiredMarker
				m.TypeName = typeutil.TypeName(imports, mtype)

				ret.Fields = append(ret.Fields, m)
				continue
			}
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

			postConstructMarker := PostConstructMarker{}
			set, err := markerdefs.Parse(c, &postConstructMarker)
			if err != nil {
				return nil, err
			} else if set {
				if ret.PostConstructMarker != nil {
					return ret, fmt.Errorf("Duplicate error: Type [%s] already has a PostConstruct method [%s] ", ret.TypeName, ret.PostConstructMarker.MethodName)
				}
				if !stringfunc.IsFirstUpper(mname) {
					return nil, fmt.Errorf("PostConstruct Method %s.%s is private ", t.Name, mname)
				}
				postConstructMarker.MethodName = mname
				ret.PostConstructMarker = &postConstructMarker
				continue
			}

			preDestroyMarker := &PreDestroyMarker{}
			set, err = markerdefs.Parse(c, preDestroyMarker)
			if err != nil {
				return nil, err
			} else if set {
				if ret.PreDestroyMarker != nil {
					return ret, fmt.Errorf("Duplicate error: Type [%s] already has a PreDestroy method [%s] ", ret.TypeName, ret.PreDestroyMarker.MethodName)
				}
				if !stringfunc.IsFirstUpper(mname) {
					return nil, fmt.Errorf("PostConstruct Method %s.%s is private ", t.Name, mname)
				}
				preDestroyMarker.MethodName = mname
				ret.PreDestroyMarker = preDestroyMarker
				continue
			}

			scopeMarker := ScopeMarker{}
			set, err = markerdefs.Parse(c, &scopeMarker)
			if err != nil {
				return nil, err
			} else if set {
				if ret.ScopeMarker != nil && ret.ScopeMarker.Value == ScopeTypePrototype {
					return nil, fmt.Errorf("Type %s is prototype, method cannot with [scope] annotation ", t.Name)
				}
				m.ScopeMarker = &scopeMarker
				continue
			}

			//autowiredMarker := AutowiredMarker{}
			//set, err = markerdefs.Parse(c, &autowiredMarker)
			//if err != nil {
			//	return nil, err
			//} else if set {
			//	if ret.ScopeMarker != nil && ret.ScopeMarker.Value == ScopeTypePrototype {
			//		return nil, fmt.Errorf("Type %s is prototype, method cannot with [autowired] annotation ", t.Name)
			//	}
			//	m.AutowiredMarker = &autowiredMarker
			//	continue
			//}

			beanMarker := BeanMarker{}
			set, err = markerdefs.Parse(c, &beanMarker)
			if err != nil {
				return nil, err
			} else if set {
				if ret.ScopeMarker != nil && ret.ScopeMarker.Value == ScopeTypePrototype {
					return nil, fmt.Errorf("Type %s is prototype, method cannot with [bean] annotation ", t.Name)
				}
				if !stringfunc.IsFirstUpper(mname) {
					return nil, fmt.Errorf("Bean Method %s.%s is private ", t.Name, mname)
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
		meta.TypeName = typeutil.TypeName(imports, v)
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

func prototype(scope *ScopeMarker) bool {
	return scope != nil && scope.Value == ScopeTypePrototype
}

func haveAutowired(core *CoreMetadata) bool {
	if len(core.Fields) != 0 {
		return true
	}
	for _, m := range core.Methods {
		if m.AutowiredMarker != nil {
			return true
		}
	}
	return false
}

func haveBean(core *CoreMetadata) bool {
	if core.ComponentMarker != nil {
		return true
	}
	if core.ServiceMarker != nil {
		return true
	}
	if core.ControllerMarker != nil {
		return true
	}
	return false
}

func valueOrName(core *CoreMetadata) string {
	if core.ComponentMarker != nil {
		return core.ComponentMarker.Value
	}
	if core.ServiceMarker != nil {
		return core.ServiceMarker.Value
	}
	if core.ControllerMarker != nil {
		return core.ControllerMarker.Value
	}
	return ""
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
		"add":             add,
		"concat":          concat,
		"prototype":       prototype,
		"firstLower":      stringfunc.FirstLower,
		"haveAutowired":   haveAutowired,
		"haveBean":        haveBean,
		"beanValueOrName": valueOrName,
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
func SetTemplate(tmpl string) Opt {
	return func(plugin *corePlugin) {
		plugin.template = getBuildTemplate(tmpl)
	}
}
