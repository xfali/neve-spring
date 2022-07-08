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

package example

import (
	"errors"
	"github.com/xfali/neve-spring/example/gen/entitiy"
)

// +neve:service
type UserService struct {
}

type userHandler struct {
}

type Auth interface {
	Verify(user entitiy.User) error
}

type AuthImpl string

func (a *AuthImpl) Verify(user entitiy.User) error {
	if string(*a) == user.Tel {
		return nil
	}
	return errors.New("Failed. ")
}

// +neve:component:value="userHandler"
func NewUserHandler() *userHandler {
	return &userHandler{}
}

// +neve:bean
func (s *UserService) GetAuth() Auth {
	v := AuthImpl("test")
	return &v
}
