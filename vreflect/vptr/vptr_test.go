/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		v := 1
		p := Of(v)
		assert.NotNil(t, p)
		assert.Equal(t, 1, *p)
	})

	t.Run("string", func(t *testing.T) {
		v := "hello"
		p := Of(v)
		assert.NotNil(t, p)
		assert.Equal(t, "hello", *p)
	})

	t.Run("struct", func(t *testing.T) {
		type S struct {
			ID int
		}
		v := S{ID: 10}
		p := Of(v)
		assert.NotNil(t, p)
		assert.Equal(t, 10, p.ID)
	})

	t.Run("nil interface", func(t *testing.T) {
		var v any = nil
		p := Of(v)
		assert.NotNil(t, p)
		assert.Nil(t, *p)
	})

	t.Run("nil pointer", func(t *testing.T) {
		var v *int = nil
		p := Of(v)
		assert.NotNil(t, p)
		assert.Nil(t, *p)
	})
}
