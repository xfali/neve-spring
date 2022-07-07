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

package buildin

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/xfali/neve-spring/pkg/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteBuildinTemplate(tmplRoot string, target string) error {
	err := utils.Mkdir(filepath.Dir(target))
	if err != nil {
		return err
	}

	f, err := os.Open(tmplRoot)
	if err != nil {
		return fmt.Errorf("Open template director %s failed: %v ", tmplRoot, err)
	}
	defer f.Close()
	fis, err := f.Readdir(-1)
	if err != nil {
		return fmt.Errorf("Read template director %s failed: %v ", tmplRoot, err)
	}

	buf := bytes.Buffer{}
	buf.Grow(4 * 1024 * 1024)
	buf.WriteString("package buildin\n\n")
	buf.WriteString("var buildinTemplate = map[string]string{}\n\nfunc init() {\n")
	for _, fi := range fis {
		// Skip dir
		if fi.IsDir() {
			continue
		}

		name := filepath.Base(fi.Name())
		buf.WriteString(fmt.Sprintf("	buildinTemplate[\"%s\"] = `", name))
		content, err := ioutil.ReadFile(filepath.Join(tmplRoot, fi.Name()))
		if err != nil {
			return err
		}
		buf.WriteString(base64.StdEncoding.EncodeToString(content))
		buf.WriteString("`\n")
	}
	buf.WriteString("}\n")

	return ioutil.WriteFile(target, buf.Bytes(), os.ModePerm)
}

func GetBuildTemplate(name string) (s string) {
	//d, err := base64.StdEncoding.DecodeString(buildinTemplate[name])
	//if err != nil {
	//	return ""
	//}
	//return string(d)
	return
}
