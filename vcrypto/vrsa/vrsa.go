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

package vrsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

const (
	rsaKeySize = 2048
)

func Hash(data []byte) []byte {
	s := sha1.Sum(data)

	return s[:]
}

func GenerateKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	return GenerateSizedKey(rsaKeySize)
}

func GenerateSizedKey(size int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	pri, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, nil, err
	}

	return pri, &pri.PublicKey, nil
}

func GenerateKeyBytes() (privateBytes, publicBytes []byte, err error) {
	return GenerateSizedKeyBytes(rsaKeySize)
}

func GenerateSizedKeyBytes(size int) (privateBytes, publicBytes []byte, err error) {
	pri, pub, err := GenerateSizedKey(size)
	if err != nil {
		return nil, nil, err
	}

	priBytes, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		return nil, nil, err
	}

	pubBytes := x509.MarshalPKCS1PublicKey(pub)

	return priBytes, pubBytes, nil
}

func GenerateKey64() (pri64, pub64 string, err error) {
	return GenerateSizedKey64(rsaKeySize)
}

func GenerateSizedKey64(size int) (pri64, pub64 string, err error) {
	pri, pub, err := GenerateSizedKeyBytes(size)
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(pri),
		base64.StdEncoding.EncodeToString(pub),
		nil
}

var (
	ErrPrivateKeyInvalid = errors.New("private key invalid")
	ErrPubKeyInvalid     = errors.New("public key invalid")
)

func PublicKeyFrom(key []byte) (*rsa.PublicKey, error) {
	var (
		pubInterface interface{}
		err          error
	)

	pubInterface, err = x509.ParsePKCS1PublicKey(key)
	if err != nil {
		pubInterface, err = x509.ParsePKIXPublicKey(key)
		if err != nil {
			return nil, err
		}
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, ErrPubKeyInvalid
	}

	return pub, nil
}

func PublicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return PublicKeyFrom(b)
}

func PrivateKeyFrom(key []byte) (*rsa.PrivateKey, error) {
	pri, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}

	p, ok := pri.(*rsa.PrivateKey)

	if !ok {
		return nil, ErrPrivateKeyInvalid
	}

	return p, nil
}

func PrivateKeyFrom64(key string) (*rsa.PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return PrivateKeyFrom(b)
}

func PublicEncrypt(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, key, data)
}

func PublicSign(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return PublicEncrypt(key, Hash(data))
}

func PublicVerify(key *rsa.PublicKey, sign, data []byte) error {
	return rsa.VerifyPKCS1v15(key, crypto.SHA1, Hash(data), sign)
}

func PrivateDecrypt(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, key, data)
}

func PrivateSign(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, Hash(data))
}

func PrivateVerify(key *rsa.PrivateKey, sign, data []byte) error {
	h, err := PrivateDecrypt(key, sign)
	if err != nil {
		return err
	}

	if !bytes.Equal(h, Hash(data)) {
		return rsa.ErrVerification
	}

	return nil
}
