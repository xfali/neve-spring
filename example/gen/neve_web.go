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
	"github.com/xfali/neve-spring/example/gen/entitiy"
)

// +neve:controller:value="users"
// +neve:requestmapping:value="/users"
type UserController struct {
}

// +neve:swagger:apioperation:value="create user"
// +neve:requestmapping:value="",method="POST"
// +neve:requestparam:name="projectId",default="-1",required=true
// +neve:requestbody:name="user",required=false
// +neve:requestheader:name="client",default="1234129040912",required=false
// +neve:loghttp
func (c *UserController) Create(projectId string, client string, user entitiy.User) entitiy.Response {
	// Business codes...
	return entitiy.Response{}
}

// +neve:swagger:apioperation:value="get user list"
// +neve:requestmapping:method="GET"
// +neve:requestparam:name="projectId",default="-1",required=true
// +neve:requestparam:name="page",default="0"
// +neve:requestparam:name="pageSize",default="20"
// +neve:loghttp
func (c *UserController) Get(projectId string, page int64, pageSize int64, orderby string) entitiy.Response {
	// Business codes...
	return entitiy.Response{Data: []entitiy.User{{}}}
}

// +neve:swagger:apioperation:value="get user detail"
// +neve:requestmapping:value="/:userId",method="GET"
// +neve:pathvariable:name="userId",required=true
// +neve:requestparam:name="projectId",default="-1",required=true
func (c *UserController) Detail(projectId string, userId int64) entitiy.Response {
	// Business codes...
	return entitiy.Response{Data: entitiy.User{}}
}

// +neve:swagger:apioperation:value="delete user by id"
// +neve:requestmapping:value="/project/:userId",method="DELETE"
// +neve:pathvariable:name="userId",required=true
// +neve:requestparam:name="projectId",default="-1",required=true
func (c *UserController) Delete(projectId string, userId int64) entitiy.Response {
	// Business codes...
	return entitiy.Response{Data: entitiy.User{}}
}
