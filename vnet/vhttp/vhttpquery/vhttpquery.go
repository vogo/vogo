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

package vhttpquery

import (
	"net/http"
	"strconv"
)

func String(r *http.Request, name string) (string, bool) {
	query := r.URL.Query()
	if len(query) == 0 {
		return "", false
	}
	param := query.Get(name)
	if param == "" {
		return "", false
	}
	return param, true
}

func Int(r *http.Request, name string) (int, bool) {
	param, ok := String(r, name)
	if !ok {
		return 0, false
	}

	i, err := strconv.Atoi(param)
	if err != nil {
		return 0, false
	}

	return i, true
}

func Float(r *http.Request, name string) (float64, bool) {
	param, ok := String(r, name)
	if !ok {
		return 0, false
	}
	f, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

func Bool(r *http.Request, name string) (bool, bool) {
	param, ok := String(r, name)
	if !ok {
		return false, false
	}
	b, err := strconv.ParseBool(param)
	if err != nil {
		return false, false
	}

	return b, true
}
