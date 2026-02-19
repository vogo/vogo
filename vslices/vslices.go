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

package vslices

import "slices"

// AppendIfNotExist appends elem to slice if elem does not exist in slice.
func AppendIfNotExist[T comparable](slice []T, elem T) []T {
	if slices.Contains(slice, elem) {
		return slice
	}
	return append(slice, elem)
}

// AppendIfCheckPass appends elem to slice if checker returns true.
func AppendIfCheckPass[T any](slice []T, elem T, checker func([]T, T) bool) []T {
	if checker(slice, elem) {
		return append(slice, elem)
	}
	return slice
}

// MapTo maps the slice from From to To.
func MapTo[From, To any](from []From, mapper func(From) To) []To {
	to := make([]To, 0, len(from))
	for _, v := range from {
		to = append(to, mapper(v))
	}
	return to
}
