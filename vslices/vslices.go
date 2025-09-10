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

// AppendIfNotExist appends elem to slice if elem does not exist in slice.
func AppendIfNotExist[T comparable](slice []T, elem T) []T {
	for _, v := range slice {
		if v == elem {
			return slice
		}
	}
	return append(slice, elem)
}

// AppendIfCheckPass appends elem to slice if checker returns true.
func AppendIfCheckPass[T any](slice []T, elem T, checker func([]T) bool) []T {
	if checker(slice) {
		return append(slice, elem)
	}
	return slice
}
