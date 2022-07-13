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

package plugin

import "testing"

func TestSwagger(t *testing.T) {
	if swaggerRouter("") != "" {
		t.Fatal("expect empty but get: ", swaggerRouter(""))
	}
	if swaggerRouter("/") != "/" {
		t.Fatal("expect / but get: ", swaggerRouter("/"))
	}
	if swaggerRouter("/:id") != "/{id}" {
		t.Fatal("expect {id} but get: ", swaggerRouter("/:id"))
	}

	if swaggerRouter("/test/:id") != "/test/{id}" {
		t.Fatal("expect /test/{id} but get: ", swaggerRouter("/test/:id"))
	}
	if swaggerRouter("/test/:id/") != "/test/{id}/" {
		t.Fatal("expect /test/{id}/ but get: ", swaggerRouter("/test/:id/"))
	}
	if swaggerRouter("/test/:id/test") != "/test/{id}/test" {
		t.Fatal("expect /test/{id}/test but get: ", swaggerRouter("/test/:id/test"))
	}
}
