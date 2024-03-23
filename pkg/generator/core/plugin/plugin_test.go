/*
 * Copyright (C) 2024, Xiongfa Li.
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
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestTpl(t *testing.T) {
	v, _ := base64.StdEncoding.DecodeString(buildinTemplate["core.tmpl"])
	ioutil.WriteFile("core.tmpl", v, 0666)
	t.Log(string(v))
}
