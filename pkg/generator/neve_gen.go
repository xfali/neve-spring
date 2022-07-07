// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package generator

import (
	"fmt"
	"github.com/xfali/neve-spring/pkg/generator/web"
	"github.com/xfali/neve-spring/pkg/generator/web/plugin"
	"path/filepath"
	"strings"

	"k8s.io/gengo/args"
	"k8s.io/gengo/examples/set-gen/sets"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"k8s.io/klog/v2"
)

const (
	defaultAnnotation = "+neve:"
)

var (
	neveImports = []string{"github.com/xfali/neve-core"}
)

type neveAnnotion struct {
	rawTypeName string
	body        string
}

type neveAnnotions map[string]*neveAnnotion

func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public": namer.NewPrivateNamer(0, ""),
		"raw":    namer.NewRawNamer("", nil),
	}
}

func DefaultNameSystem() string {
	return "public"
}

func checkEnable(annotation string, comments []string) bool {
	key := annotation + "enable"
	for _, c := range comments {
		if strings.HasPrefix(c, key) {
			return true
		}
	}
	return false
}

func GenPackages(ctx *generator.Context, args *args.GeneratorArgs) generator.Packages {
	inputs := sets.NewString(ctx.Inputs...)
	pkgs := generator.Packages{}
	annotation := defaultAnnotation
	if args.CustomArgs != nil {
		annotation = args.CustomArgs.(fmt.Stringer).String()
	}
	pluginMgr := plugin.NewWebPluginManager(annotation)
	boilerplate, err := args.LoadGoBoilerplate()
	if err != nil {
		klog.Warningf("LoadGoBoilerplate failed: %v. ", err)
		boilerplate = nil
	}
	header := []byte(fmt.Sprintf("// +build !%s\n\n", args.GeneratedBuildTag))
	if boilerplate != nil {
		header = append(header, boilerplate...)
	}

	for in := range inputs {
		klog.V(5).Infof("Parsing pkg %s\n", in)
		pkg := ctx.Universe[in]
		if pkg == nil {
			continue
		}
		for _, i := range pkg.Imports {
			ctx.AddDirectory(i.Path)
		}

		if !checkEnable(annotation, pkg.Comments) {
			continue
		}

		klog.V(5).Infof("Generating package %s...\n", in)

		pkgs = append(pkgs, &generator.DefaultPackage{
			PackageName: strings.Split(filepath.Base(pkg.Path), ".")[0],
			PackagePath: pkg.Path,
			HeaderText:  header,
			GeneratorFunc: func(context *generator.Context) []generator.Generator {
				return []generator.Generator{
					web.NewWebGenerator(args.OutputFileBaseName, annotation, pkg, pluginMgr),
				}
			},
			FilterFunc: func(context *generator.Context, i *types.Type) bool {
				return i.Name.Package == pkg.Path
			},
		})
	}
	return pkgs
}