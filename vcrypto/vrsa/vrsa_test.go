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

// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vrsa_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vcrypto/vrsa"
	"github.com/vogo/vogo/vlog"
)

func TestGenerateAndParse(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	b := vrsa.Hash([]byte("test"))
	assert.Equal(t, 20, len(b))

	data := "1564108511187|test|id"
	b = vrsa.Hash([]byte(data))

	b64 := base64.StdEncoding.EncodeToString(b)
	assert.Equal(t, "Am6FLIqmJx8C2rSnnKHUPyb6kqU=", b64)
}

func TestRsa(t *testing.T) {
	t.Parallel()

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
	vlog.Infof("public sign | sign: %s", base64.StdEncoding.EncodeToString(pubSign))

	err = vrsa.PrivateVerify(pri, pubSign, b)
	assert.Nil(t, err)

	priSign, err := vrsa.PrivateSign(pri, b)
	assert.Nil(t, err)
	vlog.Infof("private sign | sign: %s", base64.StdEncoding.EncodeToString(priSign))

	err = vrsa.PublicVerify(pub, priSign, b)
	assert.Nil(t, err)
}
