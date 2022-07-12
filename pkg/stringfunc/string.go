// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

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
