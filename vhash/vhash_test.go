// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vhash

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	b := make([]byte, 100)
	_, _ = rand.Reader.Read(b)
	s := base64.StdEncoding.EncodeToString(b)
	m := Md5(s)
	fmt.Println(m)
}
