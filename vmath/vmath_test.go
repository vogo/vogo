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

// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vmath_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vmath"
)

func TestRoundFloat64(t *testing.T) {
	assert.Equal(t, 12.23, vmath.RoundFloat64(12.2345678, 2))
	assert.Equal(t, 12.235, vmath.RoundFloat64(12.2345678, 3))
	assert.Equal(t, 12.2346, vmath.RoundFloat64(12.2345678, 4))
	assert.Equal(t, 12.23457, vmath.RoundFloat64(12.2345678, 5))
	assert.Equal(t, 12.23456, vmath.RoundFloat64(12.2345632, 5))
	assert.Equal(t, 12.234563, vmath.RoundFloat64(12.2345632, 6))
	assert.Equal(t, 12.2345632, vmath.RoundFloat64(12.2345632, 7))
	assert.Equal(t, 12.2345632, vmath.RoundFloat64(12.2345632, 8))
	assert.Equal(t, 12.2345632, vmath.RoundFloat64(12.2345632, 9))
}
