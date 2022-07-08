/*
 * Copyright (C) 2022, Xiongfa Li.
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

type ComponentMarker struct {
	Value string `marker:"value,optional"`
}

func (ComponentMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "Component",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define Component.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

type ServiceMarker struct {
	Value string `marker:"value,optional"`
}

func (ServiceMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "Service",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define Service.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

type BeanMarker struct {
	Value string `marker:"value,optional"`
}

func (BeanMarker) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "Bean",
		DetailedHelp: markers.DetailedHelp{
			Summary: "Define Bean.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func init() {
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:controller", markers.DescribesType, ControllerMarker{})).
		WithHelp(ControllerMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:component", markers.DescribesType, ComponentMarker{})).
		WithHelp(ComponentMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:service", markers.DescribesType, ServiceMarker{})).
		WithHelp(ServiceMarker{}.Help()))
	markerdefs.Register(markerdefs.Must(markers.MakeDefinition("neve:bean", markers.DescribesType, BeanMarker{})).
		WithHelp(BeanMarker{}.Help()))
}