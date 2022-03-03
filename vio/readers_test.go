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

package vio_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vio"
)

func TestByteReader_Read(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	r := vio.NewBytesReader()
	p := [10]byte{}
	_, err := r.Read(p[:])
	assert.Equal(t, io.EOF, err)
}
