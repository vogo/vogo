// Copyright 2019 The vogo Authors. All rights reserved.

package vhash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5(key ...string) string {
	md5Ctx := md5.New()

	for _, s := range key {
		_, err := md5Ctx.Write([]byte(s))
		if err != nil {
			panic(err)
		}
	}

	cipherStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func Sha1(data []byte) []byte {
	s := sha1.Sum(data)
	return s[:]
}

func Sha1String(data string) string {
	return hex.EncodeToString(Sha1([]byte(data)))
}
