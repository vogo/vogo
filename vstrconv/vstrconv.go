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

package vstrconv

import (
	"log"
	"strconv"
)

func EnsureInt(s string) int {
	return int(EnsureInt64(s))
}

func EnsureInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panicf("parse int64 error: %v", err)
	}
	return i
}

func EnsureInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Panicf("parse int32 error: %v", err)
	}
	return int32(i)
}

func EnsureUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Panicf("parse uint error: %v", err)
	}
	return uint(i)
}

func EnsureBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Panicf("parse bool error: %v", err)
	}
	return b
}

func EnsureFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Panicf("parse float64 error: %v", err)
	}
	return f
}

func EnsureFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Panicf("parse float32 error: %v", err)
	}
	return float32(f)
}

func EnsureUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Panicf("parse uint64 error: %v", err)
	}
	return i
}

func Int(s string) int {
	if s == "" {
		return 0
	}
	return EnsureInt(s)
}

func Int64(s string) int64 {
	if s == "" {
		return 0
	}
	return EnsureInt64(s)
}

func Int32(s string) int32 {
	if s == "" {
		return 0
	}
	return EnsureInt32(s)
}

func Uint(s string) uint {
	if s == "" {
		return 0
	}
	return EnsureUint(s)
}

func Bool(s string) bool {
	if s == "" {
		return false
	}
	return EnsureBool(s)
}

func Float64(s string) float64 {
	if s == "" {
		return 0
	}
	return EnsureFloat64(s)
}

func Float32(s string) float32 {
	if s == "" {
		return 0
	}
	return EnsureFloat32(s)
}

func Uint64(s string) uint64 {
	if s == "" {
		return 0
	}
	return EnsureUint64(s)
}

func I64toa(i int64) string {
	return strconv.FormatInt(i, 10)
}

func I32toa(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Ftoa(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func F32toa(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}
