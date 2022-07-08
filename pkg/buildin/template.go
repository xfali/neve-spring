// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

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

func WriteBuildinTemplate(pkg string, tmplRoot string, target string) error {
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
	buf.WriteString(fmt.Sprintf("package %s\n\n", pkg))
	buf.WriteString(`import "encoding/base64"` + "\n\n")

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
	buf.WriteString("}\n\n")

	buf.WriteString(`func GetBuildTemplate(name string) string {
	d, err := base64.StdEncoding.DecodeString(buildinTemplate[name])
	if err != nil {
		return ""
	}
	return string(d)
}`)
	return ioutil.WriteFile(target, buf.Bytes(), os.ModePerm)
}


