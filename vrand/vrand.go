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

// Intn returns a non-negative pseudo-random int in [0,n).
func Intn(n int) int {
	return int(Intn64(int64(n)))
}

// RandomString return a random string with given length, and all characters are from the source string.
func RandomString(src string, length int) string {
	srcLen := len(src)

	buf := make([]byte, length)

	for i := 0; i < length; i++ {
		buf[i] = src[Intn(srcLen)]
	}

	mathRand.Shuffle(length, func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}

// RandomSeedString return random string as function RandomString, but set seed first.
// Deprecated: use RandomString instead.
func RandomSeedString(seed int64, src string, length int) string {
	return RandomString(src, length)
}
