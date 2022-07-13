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

package web

import (
	"fmt"
	plugin2 "github.com/xfali/neve-spring/pkg/generator/plugin"
	"io"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"k8s.io/klog/v2"
)

var (
	neveWebImports = []string{
		"github.com/xfali/xlog",
		"github.com/gin-gonic/gin",
		"github.com/xfali/neve-web/gineve/midware/loghttp",
		"github.com/xfali/neve-core/boot",
		"github.com/xfali/neve-utils/neverror",
		"net/http",
		"fmt",
	}
)

type neveGen struct {
	name       string
	annotation string
	targetPkg  string
	pkg        *types.Package
	imports    namer.ImportTracker
	pluginMgr  plugin2.Manager
	plugins    []plugin2.Plugin
}

func NewWebGenerator(name, annotation string, pkg *types.Package, manager plugin2.Manager) *neveGen {
	ret := &neveGen{
		name:       name,
		annotation: annotation,
		pkg:        pkg,
		targetPkg:  pkg.Path,
		imports:    generator.NewImportTracker(),
		pluginMgr:  manager,
	}

	return ret
}

// The name of this generator. Will be included in generated comments.
func (g *neveGen) Name() string {
	return g.name
}

// Filter should return true if this generator cares about this type.
// (otherwise, GenerateType will not be called.)
//
// Filter is called before any of the generator's other functions;
// subsequent calls will get a context with only the types that passed
// this filter.
func (g *neveGen) Filter(ctx *generator.Context, t *types.Type) bool {
	return true
}

// If this generator needs special namers, return them here. These will
// override the original namers in the context if there is a collision.
// You may return nil if you don't need special names. These names will
// be available in the context passed to the rest of the generator's
// functions.
//
// A use case for this is to return a namer that tracks imports.
func (g *neveGen) Namers(ctx *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.targetPkg, g.imports),
	}
}

// Init should write an init function, and any other content that's not
// generated per-type. (It's not intended for generator specific
// initialization! Do that when your Package constructs the
// Generators.)
func (g *neveGen) Init(ctx *generator.Context, w io.Writer) error {
	g.plugins = nil
	return nil
}

// Finalize should write finish up functions, and any other content that's not
// generated per-type.
func (g *neveGen) Finalize(ctx *generator.Context, w io.Writer) error {
	for _, p := range g.plugins {
		err := p.Finalize(ctx, g.imports, w)
		if err != nil {
			return err
		}
	}
	return nil
}

// PackageVars should emit an array of variable lines. They will be
// placed in a var ( ... ) block. There's no need to include a leading
// \t or trailing \n.
func (g *neveGen) PackageVars(ctx *generator.Context) []string {
	return nil
}

// PackageConsts should emit an array of constant lines. They will be
// placed in a const ( ... ) block. There's no need to include a leading
// \t or trailing \n.
func (g *neveGen) PackageConsts(ctx *generator.Context) []string {
	return nil
}

// GenerateType should emit the code for a particular type.
func (g *neveGen) GenerateType(ctx *generator.Context, t *types.Type, w io.Writer) error {
	p := g.pluginMgr.FindPlugin(t)
	if p == nil {
		klog.V(4).Infof("Cannot handle type: %s. ", t.String())
		return nil
	}
	err := p.Generate(ctx, g.imports, w, t)
	if err != nil {
		err = fmt.Errorf("Generate by plugin: %s failed, pkg: %s type %s, err: %v. ", p.Name(), g.pkg.Path, t.Name, err)
	}
	g.plugins = append(g.plugins, p)
	return err
}

// Imports should return a list of necessary imports. They will be
// formatted correctly. You do not need to include quotation marks,
// return only the package name; alternatively, you can also return
// imports in the format `name "path/to/pkg"`. Imports will be called
// after Init, PackageVars, PackageConsts, and GenerateType, to allow
// you to keep track of what imports you actually need.
func (g *neveGen) Imports(ctx *generator.Context) []string {
	imports := g.imports.ImportLines()
	imports = append(imports, neveWebImports...)
	return imports
}

// Preferred file name of this generator, not including a path. It is
// allowed for multiple generators to use the same filename, but it's
// up to you to make sure they don't have colliding import names.
// TODO: provide per-file import tracking, removing the requirement
// that generators coordinate..
func (g *neveGen) Filename() string {
	return g.name + "_gin.go"
}

// A registered file type in the context to generate this file with. If
// the FileType is not found in the context, execution will stop.
func (g *neveGen) FileType() string {
	return generator.GolangFileType
}
