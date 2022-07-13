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

package stringfunc

import "strings"

func FirstLower(src string) string {
	if len(src) == 0 {
		return src
	}
	b := strings.Builder{}
	b.Grow(len(src))
	b.WriteString(strings.ToLower(src[:1]))
	b.WriteString(src[1:])
	return b.String()
}

func FirstUpper(src string) string {
	if len(src) == 0 {
		return src
	}
	b := strings.Builder{}
	b.Grow(len(src))
	b.WriteString(strings.ToUpper(src[:1]))
	b.WriteString(src[1:])
	return b.String()
}

func IsFirstUpper(str string) bool {
	return strIn(str[:1], 'A', 'Z')
}

func IsFirstLower(str string) bool {
	return strIn(str[:1], 'a', 'z')
}

func strIn(s string, start, end byte) bool {
	for i := range s {
		if !(start <= s[i] && s[i] <= end) {
			return false
		}
	}
	return true
}
