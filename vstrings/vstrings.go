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

package vstrings

import "strings"

func ContainsIn(items []string, item string) bool {
	if len(items) == 0 {
		return false
	}

	for _, n := range items {
		if item == n {
			return true
		}
	}

	return false
}

func ContainsAny(s string, test ...string) bool {
	for _, t := range test {
		if strings.Contains(s, t) {
			return true
		}
	}

	return false
}

// AfterFirst returns the substring after the first occurrence of sep in s.
// It returns s if sep is not present in s.
func AfterFirst(s string, sep string) string {
	index := strings.Index(s, sep)
	if index < 0 {
		return s
	}

	return s[index+len(sep):]
}

// AfterLast returns the substring after the last occurrence of sep in s.
// It returns s if sep is not present in s.
func AfterLast(s string, sep string) string {
	index := strings.LastIndex(s, sep)
	if index < 0 {
		return s
	}

	return s[index+1:]
}

// BeforeFirst returns the substring before the first occurrence of sep in s.
// It returns s if sep is not present in s.
func BeforeFirst(s string, sep string) string {
	index := strings.Index(s, sep)
	if index < 0 {
		return s
	}

	return s[:index]
}

// BeforeLast returns the substring before the last occurrence of sep in s.
// It returns s if sep is not present in s.
func BeforeLast(s string, sep string) string {
	index := strings.LastIndex(s, sep)
	if index < 0 {
		return s
	}

	return s[:index]
}
