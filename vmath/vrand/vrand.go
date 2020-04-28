// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vrand

import (
	"math/rand"
)

// RandomString return a random string with given length, and all characters are from the source string.
func RandomString(src string, length int) string {
	srcLen := len(src)

	buf := make([]byte, length)

	for i := 0; i < length; i++ {
		buf[i] = src[rand.Intn(srcLen)]
	}

	rand.Shuffle(length, func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}

// RandomSeedString return random string as function RandomString, but set seed first.
func RandomSeedString(seed int64, src string, length int) string {
	rand.Seed(seed)
	return RandomString(src, length)
}
