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

package example

import (
	"errors"
	"github.com/xfali/neve-spring/example/gen/entitiy"
	"log"
)

// +neve:service
type UserService struct {
	// +neve:autowired:name="userHandler",required=false
	UserHandler *UserHandler
}

// +neve:service:value="test"
type UserService2 struct {
	// +neve:autowired:name="userHandler",required=true
	UserHandler *UserHandler
}

// +neve:postconstruct
func (s *UserService2) Construct() {
	log.Println("UserService2 Construct")
}

// +neve:service:value="test"
type UserService3 struct {
	// +neve:autowired:name="userHandler",required=true
	UserHandler *UserHandler
}

// +neve:postconstruct
func (s *UserService3) Construct() {
	log.Println("UserService3 Construct")
}

func (s *UserService3) BeanAfterSet() error {
	log.Println("UserService3 BeanAfterSet")
	return nil
}

// +neve:predestroy
func (s *UserService3) Destroy() error {
	log.Println("UserService3 Destroy")
	return nil
}

type UserHandler struct {
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

// +neve:bean:value="userHandler",destroymethod="Purge"
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Purge() {
}

// +neve:scope:value=prototype
// +neve:bean:initmethod="Init"
func (s *UserService) GetAuth(h *UserHandler) Auth {
	v := AuthImpl("test")
	return &v
}

func (s *UserService) Init() error {
	return nil
}
