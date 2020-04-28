// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vrand_test

import (
	"testing"
	"time"

	"github.com/vogo/vogo/vmath/vrand"
)

const (
	randomSrc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vrand.RandomString(randomSrc, 16)
	}
}

func BenchmarkRandomSeedString(b *testing.B) {
	seed := time.Now().UnixNano()

	for i := 0; i < b.N; i++ {
		vrand.RandomSeedString(seed, randomSrc, 16)
	}
}
