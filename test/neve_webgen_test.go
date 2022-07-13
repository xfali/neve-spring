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
	"fmt"
	generator2 "github.com/xfali/neve-spring/pkg/generator"
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/parser"
	"strings"
	"testing"
)

func NewBuilder(g *args.GeneratorArgs) (*parser.Builder, error) {
	b := parser.New()

	// flag for including *_test.go
	b.IncludeTestFiles = g.IncludeTestFiles

	// Ignore all auto-generated files.
	b.AddBuildTags(g.GeneratedBuildTag)

	d := "github.com/xfali/neve-spring/example/gen"
	var err error
	if strings.HasSuffix(d, "/...") {
		err = b.AddDirRecursive(strings.TrimSuffix(d, "/..."))
	} else {
		err = b.AddDir(d)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to add directory %q: %v", d, err)
	}
	return b, nil
}

type cArgs struct {
	Prefix string
}

func (arg *cArgs) String() string {
	return arg.Prefix
}

func TestWebGen(t *testing.T) {
	g := args.Default().WithoutDefaultFlagParsing()
	g.GoHeaderFilePath = "../hack/boilerplate.txt"
	g.OutputBase = "."
	g.OutputFileBaseName = "zz_generated"
	b, err := NewBuilder(g)
	if err != nil {
		t.Fatal(err)
	}
	c, err := generator.NewContext(b, generator2.NameSystems(), generator2.DefaultNameSystem())
	if err != nil {
		t.Fatal(err)
	}

	c.TrimPathPrefix = ""

	c.Verify = false

	packages := generator2.GenPackages(c, g)
	if err := c.ExecutePackages(g.OutputBase, packages); err != nil {
		t.Fatal("Failed executing generator ", err)
	}
}
