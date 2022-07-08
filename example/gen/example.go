/*
 * Copyright (c) 2022, Xiongfa Li.
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

package example

import (
	"context"
	"github.com/xfali/neve-spring/example/gen/entitiy"
)

// +neve:controller:value="users"
// +neve:requestmapping:value="/users"
type UserController struct {
}

// +neve:swagger:apioperation
// +neve:requestmapping:value="",method="POST"
// +neve:requestparam:name="projectId",default="-1",required=true
// +neve:requestheader:name="client",default="",required=false
// +neve:requestbody:name="user",required=false
func (c *UserController) Create(ctx context.Context, projectId string, client string, user entitiy.User) entitiy.Response {
	// Business codes...
	return entitiy.Response{}
}

// +neve:swagger:apioperation
// +neve:requestmapping:value="",method="GET"
// +neve:requestparam:name="projectId",default="-1",required=true
// +neve:requestparam:name="page",default="0"
// +neve:requestparam:name="pageSize",default="20"
// +neve:requestheader:name="client",default="",required=false
func (c *UserController) Get(ctx context.Context, projectId string, page int64, pageSize int64) entitiy.Response {
	// Business codes...
	return entitiy.Response{Data: []entitiy.User{{}}}
}
