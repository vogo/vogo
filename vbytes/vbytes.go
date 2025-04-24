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

package vbytes

import (
	"bytes"
	"errors"
	"io"
	"time"
)

const (
	DefaultBufferSize = 32 * 1024
)

// ErrTimeout is returned from timed out calls.
var ErrTimeout = errors.New("timeout")

func CopyFilterBytes(from io.Reader, target io.Writer, filter []byte) error {
	var (
		index int
		err   error
	)

	data := make([]byte, DefaultBufferSize)

	for {
		data = data[:cap(data)]
		index, err = from.Read(data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		data = data[:index]

		if _, err = target.Write(bytes.ReplaceAll(data, filter, nil)); err != nil {
			return err
		}
	}

	return nil
}

// TimeoutCopy copy in a timeout duration.
func TimeoutCopy(dst io.Writer, src io.Reader, timeout time.Duration) error {
	buf := make([]byte, DefaultBufferSize)
	timer := time.NewTimer(timeout)

	for {
		select {
		case <-timer.C:
			return ErrTimeout
		default:
			readSize, err := src.Read(buf)

			if readSize > 0 {
				writeSize, writeErr := dst.Write(buf[0:readSize])
				if writeErr != nil {
					return writeErr
				}

				if readSize != writeSize {
					return io.ErrShortWrite
				}
			}

			if err != nil {
				if err != io.EOF {
					return err
				}

				return nil
			}
		}
	}
}
