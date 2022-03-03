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

package vhash_test

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/vogo/vogo/vhash"
)

func TestMd5(t *testing.T) {
	b := make([]byte, 100)
	_, _ = rand.Reader.Read(b)
	s := base64.StdEncoding.EncodeToString(b)
	m := vhash.Md5(s)
	fmt.Println(m)
}

func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vhash.Md5("hello")
	}
}

func BenchmarkSha1String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vhash.Sha1String("hello")
	}
}

func BenchmarkSha1StringsOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vhash.Sha1Strings("hello")
	}
}

func BenchmarkSha1StringsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vhash.Sha1Strings("hello", "world")
	}
}
