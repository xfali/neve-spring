/*
 * Copyright (C) 2019-2022, Xiongfa Li.
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
	"github.com/xfali/neve-spring/pkg/generator/plugin"
	"k8s.io/gengo/types"
)

type pluginManager struct {
	annotation string
	plugins    []plugin.Plugin
}

func NewWebPluginManager(annotation string, opts ...Opt) *pluginManager {
	ret := &pluginManager{
		annotation: annotation,
	}
	ret.RegisterPlugin(NewGinPlugin(annotation, opts...))
	return ret
}

func (m *pluginManager) FindPlugin(t *types.Type) plugin.Plugin {
	for i := len(m.plugins) - 1; i >= 0; i-- {
		v := m.plugins[i]
		if v.CouldHandle(t) {
			return v
		}
	}
	return nil
}

func (m *pluginManager) RegisterPlugin(plugin plugin.Plugin) {
	m.plugins = append(m.plugins, plugin)
}
