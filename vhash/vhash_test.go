// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vhash_test

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/vogo/vogo/vhash"
)

func TestMd5(t *testing.T) {
	b := make([]byte, 100)
	_, _ = rand.Reader.Read(b)
	s := base64.StdEncoding.EncodeToString(b)
	m := vhash.Md5(s)
	fmt.Println(m)
}

func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vhash.Md5("hello")
	}
}

func BenchmarkSha1String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vhash.Sha1String("hello")
	}
}

func BenchmarkSha1StringsOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vhash.Sha1Strings("hello")
	}
}

func BenchmarkSha1StringsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vhash.Sha1Strings("hello", "world")
	}
}
