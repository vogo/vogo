package vio_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vio"
)

func TestByteReader_Read(t *testing.T) {
	r := vio.NewBytesReader(
		[]byte("hello world"),
		[]byte("big world"),
		[]byte("small world"),
		[]byte("nice world"),
		[]byte("bad world end"),
	)

	p := [10]byte{}

	n, err := r.Read(p[:])
	assert.Nil(t, err)
	assert.Equal(t, 10, n)
	assert.Equal(t, []byte("hello worl"), p[:])

	n, err = r.Read(p[:])
	assert.Nil(t, err)
	assert.Equal(t, 10, n)
	assert.Equal(t, []byte("dbig world"), p[:])

	n, err = r.Read(p[:])
	assert.Nil(t, err)
	assert.Equal(t, 10, n)
	assert.Equal(t, []byte("small worl"), p[:])

	n, err = r.Read(p[:])
	assert.Nil(t, err)
	assert.Equal(t, 10, n)
	assert.Equal(t, []byte("dnice worl"), p[:])

	n, err = r.Read(p[:])
	assert.Nil(t, err)
	assert.Equal(t, 10, n)
	assert.Equal(t, []byte("dbad world"), p[:])

	n, err = r.Read(p[:])
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, []byte(" end"), p[:n])

	_, err = r.Read(p[:])
	assert.Equal(t, err, io.EOF)
}

func TestByteReader_ReadNil(t *testing.T) {
	r := vio.NewBytesReader()
	p := [10]byte{}
	_, err := r.Read(p[:])
	assert.Equal(t, io.EOF, err)
}
