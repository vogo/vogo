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

package vregexp

import (
	"fmt"
	"strconv"
)

// RegexGroupRender returns a function that renders an expression for a match regex group.
// The expression is a string of the form "$1", "$2","name is $1", etc.
func RegexGroupRender(renderExpression string) func([][]byte) []byte {
	var (
		err   error
		list  []func([][]byte) []byte
		start int
	)

	length := len(renderExpression)
	for i := 0; i < length; i++ {
		v := renderExpression[i]

		if v == '$' {
			if i > start {
				sub := []byte(renderExpression[start:i])

				list = append(list, func([][]byte) []byte {
					return sub
				})
			}

			j := i + 1
			for ; j < length && (renderExpression[j] >= '0' && renderExpression[j] <= '9'); j++ {
			}

			var groupIndex int
			groupIndex, err = strconv.Atoi(renderExpression[i+1 : j])
			if err != nil {
				panic(fmt.Errorf("vregexp: invalid group index: %s", renderExpression[i+1:j]))
			}

			list = append(list, func(groups [][]byte) []byte {
				return groups[groupIndex]
			})

			start = j
		}
	}

	if start < length {
		sub := []byte(renderExpression[start:length])

		list = append(list, func([][]byte) []byte {
			return sub
		})
	}

	return func(groups [][]byte) []byte {
		var buf []byte

		for _, f := range list {
			buf = append(buf, f(groups)...)
		}

		return buf
	}
}
