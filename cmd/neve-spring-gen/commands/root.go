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

package commands

import (
	"flag"
	"github.com/spf13/cobra"
	"github.com/xfali/neve-spring/cmd/neve-spring-gen/commands/project"
	"github.com/xfali/neve-spring/cmd/neve-spring-gen/customargs"
	"github.com/xfali/neve-spring/pkg/generator"
	"k8s.io/gengo/args"
	"k8s.io/klog/v2"
)

var (
	gArgs   *args.GeneratorArgs
	cusArgs *customargs.NeveArgs
)

var root = cobra.Command{
	Use:   "neve-spring-gen",
	Short: "Generate code with neve spring like annotations",
	Long:  "Generate code with neve spring like annotations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := customargs.Validate(gArgs); err != nil {
			klog.Fatalln(err)
		}
		if err := gArgs.Execute(generator.NameSystems(), generator.DefaultNameSystem(), generator.GenPackages); err != nil {
			klog.Fatalln(err)
		}
		klog.V(2).Info("Success")
	},
}

func Execute() {
	cobra.CheckErr(root.Execute())
}

func init() {
	gArgs, cusArgs = customargs.NewDefault()

	fs := root.Flags()
	gArgs.AddFlags(fs)
	cusArgs.AddFlags(fs)

	_ = flag.Set("logtostderr", "true")
	fs.AddGoFlagSet(flag.CommandLine)
	//_ = fs.Parse(os.Args[1:])
	root.AddCommand(project.Get())
}
