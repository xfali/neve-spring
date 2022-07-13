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
	"github.com/xfali/neve-spring/pkg/generator/markerdefs"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

type ControllerMarker struct {
	Value string `marker:"value,optional"`
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
	Value           string `marker:"value,optional"`
	Method          string `marker:"method,optional"`
	Consumes        string `marker:"consumes,optional"`
	Produces        string `marker:"produces,optional"`
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
	Name     string `marker:"name"`
	Default  string `marker:"default,optional"`
	Required bool   `marker:"required,optional"`
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

type PathVariableMarker struct {
	Name     string `marker:"name"`
	Default  string `marker:"default,optional"`
	Required bool   `marker:"required,optional"`
}

func (PathVariableMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "PathVariable",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define PathVariable.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

type RequestHeaderMarker struct {
	Name     string `marker:"name"`
	Default  string `marker:"default,optional"`
	Required bool   `marker:"required,optional"`
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
	Required        bool   `marker:"required,optional"`
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

type LogHttpMarker struct {
	NoRequestHeader  bool   `marker:"norequestheader,optional"`
	NoRequestBody    bool   `marker:"norequestbody,optional"`
	NoResponseHeader bool   `marker:"noresponseheader,optional"`
	NoResponseBody   bool   `marker:"noresponsebody,optional"`
	Level            string `marker:"level,optional"` // debug | info | warn
}

func NewLogHttpMarker() *LogHttpMarker {
	return &LogHttpMarker{
		NoRequestHeader:  false,
		NoRequestBody:    false,
		NoResponseHeader: false,
		NoResponseBody:   false,
		Level:            "info",
	}
}

func (m *LogHttpMarker) Fill() *LogHttpMarker {
	if m.Level == "" {
		m.Level = "info"
	}
	return m
}

func (LogHttpMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "LogHttp",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define LogHttp.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

type ApiOperationMarker struct {
	Value  string `marker:"value"`
	Notes  string `marker:"notes,optional"`
	Tags   string `marker:"tags,optional"`
	Router string `marker:"router,optional"`
}

func (ApiOperationMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "ApiOperation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define ApiOperation.",
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
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:pathvariable", markers.DescribesType, PathVariableMarker{})).
		WithHelp(PathVariableMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:requestparam", markers.DescribesType, RequestParamMarker{})).
		WithHelp(RequestParamMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:requestheader", markers.DescribesType, RequestHeaderMarker{})).
		WithHelp(RequestHeaderMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:requestbody", markers.DescribesType, RequestBodyMarker{})).
		WithHelp(RequestBodyMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:loghttp", markers.DescribesType, LogHttpMarker{})).
		WithHelp(LogHttpMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:swagger:apioperation", markers.DescribesType, ApiOperationMarker{})).
		WithHelp(ApiOperationMarker{}.Help()))
}
