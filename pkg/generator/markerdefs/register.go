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

package markerdefs

import (
	"fmt"
	"reflect"
	"strings"
)

var definitions = map[reflect.Type]*DefinitionWithHelp{}

type Setter interface {
	IsSet() bool
	Set(bool)
}

type Flag bool

func (s *Flag) IsSet() bool {
	return bool(*s)
}
func (s *Flag) Set(v bool) {
	*s = Flag(v)
}

func Register(defs ...*DefinitionWithHelp) {
	for _, d := range defs {
		definitions[d.Output] = d
	}
}

func Parse(comment string, o interface{}) (set bool, err error) {
	t := reflect.TypeOf(o)
	if t.Kind() != reflect.Ptr {
		return false, fmt.Errorf("Result object type %s is not ptr. ", t.String())
	}
	v := reflect.ValueOf(o).Elem()
	t = t.Elem()
	d := definitions[t]
	if d == nil {
		return false, fmt.Errorf("Result object type %s is not been registered. ", t.String())
	}
	if i := strings.Index(comment, d.Name); i != -1 {
		ret, err := d.Parse(comment)
		if err != nil {
			return false, err
		}
		vv := reflect.ValueOf(ret)
		v.Set(vv)
		return true, nil
	}
	// Comment not invalid
	return false, nil
}
