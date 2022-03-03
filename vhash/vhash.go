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

package vhash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// Md5 calculate md5 for strings.
func Md5(data ...string) string {
	md5Ctx := md5.New()

	for _, s := range data {
		_, err := md5Ctx.Write([]byte(s))
		if err != nil {
			panic(err)
		}
	}

	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// Sha1 calculate sha1 for given bytes.
func Sha1(data []byte) []byte {
	s := sha1.Sum(data)

	return s[:]
}

// Sha1String calculate sha1 for a single string.
func Sha1String(data string) string {
	return hex.EncodeToString(Sha1([]byte(data)))
}

// Sha1Strings calculate sha1 for multiple strings.
func Sha1Strings(data ...string) string {
	var b []byte
	for _, s := range data {
		b = append(b, []byte(s)...)
	}

	sum := sha1.Sum(b)

	return hex.EncodeToString(sum[:])
}
