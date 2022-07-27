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

package project

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Module     string
	Output     string
	MainOutPut string
	GoVersion  string
	NoSwagger  bool
	NoWeb      bool
	NoDatabase bool
}

var config Config

var cmd = cobra.Command{
	Use:   "init",
	Short: "Generate project init code and files",
	Long:  "Generate project init code and files",
	Run: func(cmd *cobra.Command, args []string) {
		err := generate(config.Output, "go.mod", "mod.tmpl")
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}

		if config.MainOutPut == "" {
			config.MainOutPut = filepath.Join(config.Output, "cmd", filepath.Dir(config.Module))
		}
		err = generate(config.MainOutPut, "main.go", "main.tmpl")
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}

	},
}

func generate(output string, filename string, tpl string) error {
	info, err := os.Stat(output)
	if err != nil {
		fmt.Printf("Output %s is not exists.\n", output)
		err = os.MkdirAll(output, os.ModePerm)
		if err != nil {
			return fmt.Errorf("%s not exist and create failed: %v ", output, err)
		}
	} else {
		if !info.IsDir() {
			return fmt.Errorf("Output %s is not a directory.\n", output)
		}
	}

	buf := &bytes.Buffer{}
	_, file, line, _ := runtime.Caller(1)
	tmpl, err := template.
		New(fmt.Sprintf("%s:%d", file, line)).
		Parse(getBuildTemplate(tpl))
	if err != nil {
		return err
	}
	err = tmpl.Execute(buf, config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(output, filename), buf.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Get() *cobra.Command {
	return &cmd
}

func init() {
	fs := cmd.Flags()
	fs.StringVar(&config.Module, "module", "awesomeProject", "Name of Module")
	fs.StringVarP(&config.Output, "output", "o", config.Output, "Output directory")
	fs.StringVar(&config.MainOutPut, "main-output", config.MainOutPut, "Main Output directory")
	fs.StringVar(&config.GoVersion, "go-version", "1.18", "Init with go version")
	fs.BoolVar(&config.NoSwagger, "no-swagger", config.NoSwagger, "Generate without swagger processor")
	fs.BoolVar(&config.NoWeb, "no-web", config.NoWeb, "Generate without web processor")
	fs.BoolVar(&config.NoDatabase, "no-database", config.NoDatabase, "Generate without database processor")
}
