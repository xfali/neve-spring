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

package test

import (
	"github.com/xfali/neve-spring/pkg/buildin"
	"testing"
)

func TestUpdateAllTemplate(t *testing.T) {
	err := buildin.WriteBuildinTemplate("plugin", "../pkg/generator/web/template", "../pkg/generator/web/plugin/gin_tpl.go")
	if err != nil {
		t.Fatal(err)
	}
	err = buildin.WriteBuildinTemplate("plugin", "../pkg/generator/core/template", "../pkg/generator/core/plugin/core_tpl.go")
	if err != nil {
		t.Fatal(err)
	}
	err = buildin.WriteBuildinTemplate("plugin", "../pkg/generator/restclient/template", "../pkg/generator/restclient/plugin/restclient_tpl.go")
	if err != nil {
		t.Fatal(err)
	}
	err = buildin.WriteBuildinTemplate("project", "../cmd/neve-spring-gen/commands/project/template", "../cmd/neve-spring-gen/commands/project/project_tpl.go")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdatCoreTemplate(t *testing.T) {
	err := buildin.WriteBuildinTemplate("plugin", "../pkg/generator/core/template", "../pkg/generator/core/plugin/core_tpl.go")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdatGinTemplate(t *testing.T) {
	err := buildin.WriteBuildinTemplate("plugin", "../pkg/generator/web/plugin/template", "../pkg/generator/web/plugin/gin_tpl.go")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateProjectTemplate(t *testing.T) {
	err := buildin.WriteBuildinTemplate("project", "../cmd/neve-spring-gen/commands/project/template", "../cmd/neve-spring-gen/commands/project/project_tpl.go")
	if err != nil {
		t.Fatal(err)
	}
}
