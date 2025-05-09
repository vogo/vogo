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

package vjson

import (
	"encoding/json"
	"io"
)

// EnsureUnmarshal unmarshal data and panic if has error.
func EnsureUnmarshal(data []byte, v any) {
	if err := json.Unmarshal(data, v); err != nil {
		panic(err)
	}
}

// EnsureMarshal marshal interface and panic if has error.
func EnsureMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}

// UnmarshalStream unmarshal stream data.
func UnmarshalStream(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

// MarshalStream marshal interface to stream.
func MarshalStream(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}
