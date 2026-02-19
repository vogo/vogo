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

package vrand

import (
	"crypto/rand"
	"math/big"
	mathRand "math/rand"
)

// Intn64 returns a non-negative pseudo-random int64 in [0,n).
func Intn64(n int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		return mathRand.Int63n(n)
	}

	return nBig.Int64()
}

// Intn64Range returns a non-negative pseudo-random int64 in [min,max].
func Intn64Range(min, max int64) int64 {
	if min >= max {
		min, max = max, min
	}
	return min + Intn64(max-min+1)
}

// Intn returns a non-negative pseudo-random int in [0,n).
func Intn(n int) int {
	return int(Intn64(int64(n)))
}

// IntnRange returns a non-negative pseudo-random int in [min,max).
func IntnRange(min, max int) int {
	if min >= max {
		min, max = max, min
	}
	return min + Intn(max-min)
}

// Float64 returns a pseudo-random float64 in [0.0,1.0).
func Float64() float64 {
	return mathRand.Float64()
}

// Float64Range returns a pseudo-random float64 in [min,max).
func Float64Range(min, max float64) float64 {
	if min >= max {
		min, max = max, min
	}
	return min + Float64()*(max-min)
}

// RandomBytes return a random bytes with given length, and all bytes are from the source strings.
func RandomBytes(src string, length int) []byte {
	srcLen := len(src)

	buf := make([]byte, length)

	for i := range length {
		buf[i] = src[Intn(srcLen)]
	}

	mathRand.Shuffle(length, func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return buf
}

// RandomString return a random string with given length, and all characters are from the source string.
func RandomString(src string, length int) string {
	return string(RandomBytes(src, length))
}

// RandomSeedString return random string as function RandomString, but set seed first.
// Deprecated: use RandomString instead.
func RandomSeedString(seed int64, src string, length int) string {
	return RandomString(src, length)
}
