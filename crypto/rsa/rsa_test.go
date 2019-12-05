// Copyright 2019 The vogo Authors. All rights reserved.

package vrsa

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndParse(t *testing.T) {
	pri64, pub64, err := GenerateKey64()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	_, err = PrivateKeyFrom64(pri64)
	assert.Nil(t, err)

	_, err = PublicKeyFrom64(pub64)
	assert.Nil(t, err)
}

func TestHash(t *testing.T) {
	b := hash([]byte("test"))
	assert.Equal(t, 20, len(b))

	data := "1564108511187|test|id"
	b = hash([]byte(data))

	b64 := base64.StdEncoding.EncodeToString(b)
	assert.Equal(t, "Am6FLIqmJx8C2rSnnKHUPyb6kqU=", b64)
}

func TestRsa(t *testing.T) {
	pri, pub, err := GenerateKey()
	assert.Nil(t, err)

	b := []byte("hello world")

	// ---------------encrypt/decrypt--------------------------
	pubEnc, err := PublicEncrypt(pub, b)
	assert.Nil(t, err)

	priDec, err := PrivateDecrypt(pri, pubEnc)
	assert.Nil(t, err)

	if !bytes.Equal(b, priDec) {
		assert.FailNow(t, "private decrypt not the same as source")
	}

	// ---------------sign/verify--------------------------
	pubSign, err := PublicSign(pub, b)
	assert.Nil(t, err)
	fmt.Println("public sign:", base64.StdEncoding.EncodeToString(pubSign))

	err = PrivateVerify(pri, pubSign, b)
	assert.Nil(t, err)

	priSign, err := PrivateSign(pri, b)
	assert.Nil(t, err)
	fmt.Println("private sign:", base64.StdEncoding.EncodeToString(priSign))

	err = PublicVerify(pub, priSign, b)
	assert.Nil(t, err)
}
