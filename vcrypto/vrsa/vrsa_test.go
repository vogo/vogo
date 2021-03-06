// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vrsa_test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vcrypto/vrsa"
)

func TestGenerateAndParse(t *testing.T) {
	pri64, pub64, err := vrsa.GenerateKey64()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	_, err = vrsa.PrivateKeyFrom64(pri64)
	assert.Nil(t, err)

	_, err = vrsa.PublicKeyFrom64(pub64)
	assert.Nil(t, err)
}

func TestHash(t *testing.T) {
	b := vrsa.Hash([]byte("test"))
	assert.Equal(t, 20, len(b))

	data := "1564108511187|test|id"
	b = vrsa.Hash([]byte(data))

	b64 := base64.StdEncoding.EncodeToString(b)
	assert.Equal(t, "Am6FLIqmJx8C2rSnnKHUPyb6kqU=", b64)
}

func TestRsa(t *testing.T) {
	pri, pub, err := vrsa.GenerateKey()
	assert.Nil(t, err)

	b := []byte("hello world")

	// ---------------encrypt/decrypt--------------------------
	pubEnc, err := vrsa.PublicEncrypt(pub, b)
	assert.Nil(t, err)

	priDec, err := vrsa.PrivateDecrypt(pri, pubEnc)
	assert.Nil(t, err)

	if !bytes.Equal(b, priDec) {
		assert.FailNow(t, "private decrypt not the same as source")
	}

	// ---------------sign/verify--------------------------
	pubSign, err := vrsa.PublicSign(pub, b)
	assert.Nil(t, err)
	fmt.Println("public sign:", base64.StdEncoding.EncodeToString(pubSign))

	err = vrsa.PrivateVerify(pri, pubSign, b)
	assert.Nil(t, err)

	priSign, err := vrsa.PrivateSign(pri, b)
	assert.Nil(t, err)
	fmt.Println("private sign:", base64.StdEncoding.EncodeToString(priSign))

	err = vrsa.PublicVerify(pub, priSign, b)
	assert.Nil(t, err)
}
