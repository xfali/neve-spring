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

package typeutil

import (
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

func TypeName(imports namer.ImportTracker, v *types.Type) string {
	pt := ""
	if v.Kind == types.Pointer {
		v = v.Elem
		pt = "*"
	}
	imports.AddType(v)
	pkg := imports.LocalNameOf(v.Name.Package)
	if pkg != "" {
		pt = pt + pkg + "." + v.Name.Name
	} else {
		pt = pt + v.Name.Name
	}
	return pt
}
