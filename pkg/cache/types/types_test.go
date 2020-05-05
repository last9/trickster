/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"testing"
)

func TestCacheTypeString(t *testing.T) {

	t1 := CacheTypeMemory
	t2 := CacheTypeFilesystem
	var t3 CacheType = 13

	if t1.String() != "memory" {
		t.Errorf("expected %s got %s", "memory", t1.String())
	}

	if t2.String() != "filesystem" {
		t.Errorf("expected %s got %s", "filesystem", t2.String())
	}

	if t3.String() != "13" {
		t.Errorf("expected %s got %s", "13", t3.String())
	}

}
