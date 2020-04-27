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
	pri, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return nil, nil, err
	}

	return pri, &pri.PublicKey, nil
}

func GenerateKeyBytes() (privateBytes, publicBytes []byte, err error) {
	pri, pub, err := GenerateKey()
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
	pri, pub, err := GenerateKeyBytes()
	if err != nil {
		return "", "", nil
	}

	return base64.StdEncoding.EncodeToString(pri),
		base64.StdEncoding.EncodeToString(pub),
		nil
}

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
		return nil, errors.New("invalid public key")
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
		return nil, errors.New("invalid private key")
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
