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
	"bytes"
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
	Name               string
	Params             []*TypeMeta
	Returns            []*TypeMeta
	RequestMapping     RequestMappingMarker
	LogHttpMarker      *LogHttpMarker
	ApiOperationMarker *ApiOperationMarker
}

type GinMetadata struct {
	Name           string
	TypeName       string
	Controller     ControllerMarker
	RequestMapping RequestMappingMarker
	LogHttpMarker  *LogHttpMarker
	Methods        []*Method
}

type ginPlugin struct {
	template         string
	annotation       string
	structAnnotation string
	methodAnnotation []string
	fillFunc         AutoFillParamFunc

	metas []*GinMetadata
}

type AutoFillParamFunc func(imports namer.ImportTracker, name string, param *types.Type) (*TypeMeta, error)

type Opt func(*ginPlugin)

func NewGinPlugin(annotation string, opts ...Opt) *ginPlugin {
	ret := &ginPlugin{
		template:         getBuildTemplate("webgin.tmpl"),
		annotation:       annotation,
		structAnnotation: annotation + "controller",
		methodAnnotation: []string{
			annotation + "requestmapping",
		},
		fillFunc: autoFillQuery,
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
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
		return nil, fmt.Errorf("Type: %s without method ", t.Name)
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
			if !stringfunc.IsFirstUpper(t.Name.Name) {
				return nil, fmt.Errorf("Type %s is private ", t.Name)
			}
			continue
		}
		set, err = markerdefs.Parse(c, &ret.RequestMapping)
		if err != nil {
			return nil, err
		} else if set {
			continue
		}

		logHttp := NewLogHttpMarker()
		set, err = markerdefs.Parse(c, logHttp)
		if err != nil {
			return nil, err
		} else if set {
			ret.LogHttpMarker = logHttp.Fill()
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
					if !stringfunc.IsFirstUpper(mname) {
						return nil, fmt.Errorf("Method %s.%s is private ", t.Name, mname)
					}
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
					meta.Required = paramMarker.Required
					meta.RequestType = RequestTypeQuery
					m.Params = append(m.Params, meta)
					continue
				} else {
					return nil, fmt.Errorf("Found annotaion: requestparam with name: %s but method not exist ", paramMarker.Name)
				}
			}

			pathMarker := PathVariableMarker{}
			set, err = markerdefs.Parse(c, &pathMarker)
			if err != nil {
				return nil, err
			} else if set {
				meta, have := findParam(imports, mtype, pathMarker.Name)
				if have {
					meta.Default = pathMarker.Default
					meta.Required = pathMarker.Required
					meta.RequestType = RequestTypePath
					m.Params = append(m.Params, meta)
					continue
				} else {
					return nil, fmt.Errorf("Found annotaion: pathvariable with name: %s but method not exist ", pathMarker.Name)
				}
			}

			headerMarker := RequestHeaderMarker{}
			set, err = markerdefs.Parse(c, &headerMarker)
			if err != nil {
				return nil, err
			} else if set {
				meta, have := findParam(imports, mtype, headerMarker.Name)
				if have {
					meta.Default = headerMarker.Default
					meta.Required = headerMarker.Required
					meta.RequestType = RequestTypeHeader
					m.Params = append(m.Params, meta)
					continue
				} else {
					return nil, fmt.Errorf("Found annotaion: requestheader with name: %s but method not exist ", headerMarker.Name)
				}
			}

			bodyMarker := RequestBodyMarker{}
			set, err = markerdefs.Parse(c, &bodyMarker)
			if err != nil {
				return nil, err
			} else if set {
				meta, have := findParam(imports, mtype, bodyMarker.Name)
				if have {
					meta.Required = bodyMarker.Required
					meta.RequestType = RequestTypeBody
					m.Params = append(m.Params, meta)
				}
				continue
			}

			logHttp := NewLogHttpMarker()
			set, err = markerdefs.Parse(c, logHttp)
			if err != nil {
				return nil, err
			} else if set {
				m.LogHttpMarker = logHttp.Fill()
				continue
			}

			apiOptMarker := &ApiOperationMarker{}
			set, err = markerdefs.Parse(c, apiOptMarker)
			if err != nil {
				return nil, err
			} else if set {
				if apiOptMarker.Notes == "" {
					apiOptMarker.Notes = apiOptMarker.Value
				}
				m.ApiOperationMarker = apiOptMarker
				continue
			}
		}
		if m.LogHttpMarker == nil && ret.LogHttpMarker != nil {
			m.LogHttpMarker = ret.LogHttpMarker
		}
		err := p.validateParams(imports, mtype, m)
		if err != nil {
			return nil, err
		}
		m.Returns = findResult(imports, mtype)
	}

	return ret, nil
}

func autoFillQuery(imports namer.ImportTracker, name string, param *types.Type) (*TypeMeta, error) {
	ret := &TypeMeta{
		Required:    true,
		RequestType: RequestTypeQuery,
	}
	ret.Name = name
	if param.Kind == types.Struct || param.Kind == types.Interface {
		imports.AddType(param)
		ret.TypeName = imports.LocalNameOf(param.Name.Package) + "." + param.Name.Name
	} else {
		ret.TypeName = param.Name.Name
	}
	return ret, nil
}

func (p *ginPlugin) validateParams(imports namer.ImportTracker, t *types.Type, method *Method) error {
	for i, s := range t.Signature.ParameterNames {
		found := false
		for _, d := range method.Params {
			if s == d.Name {
				found = true
				break
			}
		}
		if !found {
			param := t.Signature.Parameters[i]
			ret, err := p.fillFunc(imports, s, param)
			if err != nil {
				return err
			}
			method.Params = append(method.Params, ret)
		}
	}
	for _, p := range method.Params {
		if p.RequestType == RequestTypePath {
			check := ":" + p.Name
			if index := strings.Index(method.RequestMapping.Value, check); index == -1 {
				return fmt.Errorf("Method [%s] RequestMapping missing Path param [%s] add with :%s ",
					method.Name, p.Name, p.Name)
			}
		}
	}
	// Sort by method parameters.
	for i, v := range t.Signature.ParameterNames {
		for j, m := range method.Params {
			if v == m.Name {
				if i != j {
					// swap.
					method.Params[j] = method.Params[i]
					method.Params[i] = m
				}
				break
			}
		}
	}
	return nil
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

func add(a, b int) int {
	return a + b
}

func configLogHttp(log *LogHttpMarker) string {
	if log.NoRequestBody || log.NoRequestHeader || log.NoResponseHeader || log.NoResponseBody || log.Level != "info" {
		buf := strings.Builder{}
		buf.WriteString(fmt.Sprintf(`loghttp.OptLogLevel("%s")`, log.Level))
		if log.NoRequestHeader {
			buf.WriteString(", loghttp.DisableLogReqHeader()")
		}
		if log.NoRequestBody {
			buf.WriteString(", loghttp.DisableLogReqBody()")
		}
		if log.NoResponseHeader {
			buf.WriteString(", loghttp.DisableLogRespHeader()")
		}
		if log.NoResponseBody {
			buf.WriteString(", loghttp.DisableLogRespBody()")
		}
		return fmt.Sprintf("h.HLog.OptLogHttp(%s)", buf.String())
	} else {
		return "h.HLog.LogHttp()"
	}
}

func swaggerRouter(s string) string {
	buf := bytes.NewBuffer(nil)
	buf.Grow(len(s))
	for {
		i := strings.Index(s, "/:")
		if i == -1 {
			buf.WriteString(s)
			return buf.String()
		} else {
			buf.WriteString(s[:i+1])
			// find next node
			n := strings.Index(s[i+2:], "/")
			if n == -1 {
				buf.WriteString(fmt.Sprintf("{%s}", s[i+2:]))
				return buf.String()
			} else {
				n += i + 2
				buf.WriteString(fmt.Sprintf("{%s}", s[i+2:n]))
				s = s[n:]
				continue
			}
		}
	}
}

func generateMethod(method *Method) (string, error) {
	tmplKey := templateMap[method.RequestMapping.Method]
	if tmplKey == "" {
		return "", fmt.Errorf("Method %s not support. ", method.RequestMapping.Method)
	}
	buf := &bytes.Buffer{}
	buf.Grow(1024)
	_, file, line, _ := runtime.Caller(1)
	tmpl, err := template.
		New(fmt.Sprintf("%s:%d", file, line)).
		Parse(getBuildTemplate(tmplKey))
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(buf, method)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
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
	//w = io.MultiWriter(w, os.Stderr)

	funcMap := template.FuncMap{
		"concatUrl":     concatUrl,
		"add":           add,
		"toLower":       strings.ToLower,
		"swaggerRouter": swaggerRouter,
		"configLogHttp": configLogHttp,
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
	p.metas = append(p.metas, meta)
	return tmpl.Execute(w, meta)
	//sw := generator.NewSnippetWriter(w, ctx, delimiterLeft, delimiterRight)
	//sw.Do(p.template, meta)
	//return sw.Error()
}

func (p *ginPlugin) Finalize(ctx *generator.Context, imports namer.ImportTracker, w io.Writer) error {
	_, file, line, _ := runtime.Caller(1)
	tmpl, err := template.
		New(fmt.Sprintf("%s:%d", file, line)).
		Parse(getBuildTemplate("webgin_register.tmpl"))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, p.metas)
}

func SetAutoFillParamFunc(f AutoFillParamFunc) Opt {
	return func(plugin *ginPlugin) {
		plugin.fillFunc = f
	}
}
