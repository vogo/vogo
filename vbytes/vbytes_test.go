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

package vbytes_test

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vbytes"
)

type treader struct {
	readLimit int
	readCount int
}

func (r *treader) Read(buf []byte) (int, error) {
	r.readCount++
	if r.readCount >= r.readLimit {
		return 0, io.EOF
	}

	fmt.Println("read limit: ", r.readLimit, ", read count: ", r.readCount)
	time.Sleep(time.Second)

	return 1024, nil
}

type twriter struct {
}

func (w *twriter) Write(buf []byte) (int, error) {
	fmt.Println("write ...")
	return 1024, nil
}

func TestTimeoutCopy(t *testing.T) {
	r := &treader{readLimit: 5}
	w := &twriter{}

	err := vbytes.TimeoutCopy(w, r, 3*time.Second)
	assert.Equal(t, vbytes.ErrTimeout, err)

	r.readCount = 0
	err = vbytes.TimeoutCopy(w, r, 10*time.Second)
	assert.Nil(t, err)
}
