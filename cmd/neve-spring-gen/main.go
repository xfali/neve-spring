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

package main

import (
	"github.com/xfali/neve-spring/cmd/neve-spring-gen/commands"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	//args, cusArgs := customargs.NewDefault()
	//
	//args.AddFlags(pflag.CommandLine)
	//cusArgs.AddFlags(pflag.CommandLine)
	//
	//_ = flag.Set("logtostderr", "true")
	//pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//pflag.Parse()
	//
	//if err := customargs.Validate(args); err != nil {
	//	klog.Fatalln(err)
	//}
	//
	//if err := args.Execute(generator.NameSystems(), generator.DefaultNameSystem(), generator.GenPackages); err != nil {
	//	klog.Fatalln(err)
	//}
	commands.Execute()
	klog.V(2).Info("Success")
}
