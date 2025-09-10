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

import (
	"strings"
	"unicode/utf8"
)

// ContainsIn checks if a string is in a string slice.
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

// ContainsAny checks if a string contains any of the test strings.
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
func AfterFirst(s, sep string) string {
	index := strings.Index(s, sep)
	if index < 0 {
		return s
	}

	return s[index+len(sep):]
}

// AfterLast returns the substring after the last occurrence of sep in s.
// It returns s if sep is not present in s.
func AfterLast(s, sep string) string {
	index := strings.LastIndex(s, sep)
	if index < 0 {
		return s
	}

	return s[index+len(sep):]
}

// BeforeFirst returns the substring before the first occurrence of sep in s.
// It returns s if sep is not present in s.
func BeforeFirst(s, sep string) string {
	index := strings.Index(s, sep)
	if index < 0 {
		return s
	}

	return s[:index]
}

// BeforeFirstInclude returns the substring before the first occurrence of sep in s, include sep.
// It returns s if sep is not present in s.
func BeforeFirstInclude(s, sep string) string {
	index := strings.Index(s, sep)
	if index < 0 {
		return s
	}
	return s[:index+len(sep)]
}

// BeforeLast returns the substring before the last occurrence of sep in s.
// It returns s if sep is not present in s.
func BeforeLast(s, sep string) string {
	index := strings.LastIndex(s, sep)
	if index < 0 {
		return s
	}

	return s[:index]
}

// BeforeLastInclude returns the substring before the last occurrence of sep in s, include sep.
// It returns s if sep is not present in s.
func BeforeLastInclude(s, sep string) string {
	index := strings.LastIndex(s, sep)
	if index < 0 {
		return s
	}
	return s[:index+len(sep)]
}

// RuneCut cuts the string by rune length.
func RuneCut(s string, max int) string {
	if max <= 0 {
		return ""
	}
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	return string([]rune(s)[:max])
}

// RuneCutLast cuts the string by rune length from the end.
func RuneCutLast(s string, max int) string {
	if max <= 0 {
		return ""
	}
	length := utf8.RuneCountInString(s)
	if length <= max {
		return s
	}
	return string([]rune(s)[length-max:])
}

const (
	Ellipsis = "..."
)

// RuneCutWithEllipsis cuts the string by rune length with ellipsis.
func RuneCutWithEllipsis(s string, max int) string {
	if max <= 0 {
		return ""
	}

	length := utf8.RuneCountInString(s)

	if length <= max {
		return s
	}

	r := []rune(s)

	if max <= len(Ellipsis) {
		return Ellipsis
	}

	return string(r[:max-len(Ellipsis)]) + Ellipsis
}

// RuneCutLastWithEllipsis cuts the string by rune length from the end with ellipsis.
func RuneCutLastWithEllipsis(s string, max int) string {
	if max <= 0 {
		return ""
	}

	length := utf8.RuneCountInString(s)
	if length <= max {
		return s
	}

	r := []rune(s)

	if max <= len(Ellipsis) {
		return Ellipsis
	}

	return Ellipsis + string(r[length-max+len(Ellipsis):])
}
