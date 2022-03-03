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

package vio

import (
	"bytes"
	"io"
)

// queueReader queue reader read from multiple readers.
type queueReader struct {
	idx     int
	total   int
	readers []io.Reader
}

// Read read bytes from multiple readers.
func (b *queueReader) Read(p []byte) (n int, err error) {
	var (
		count   int
		read    int
		readErr error
	)

	required := len(p)

	for count < required && b.idx < b.total {
		r := b.readers[b.idx]

		read, readErr = r.Read(p[count:])
		count += read

		if readErr != nil {
			if readErr != io.EOF {
				return count, readErr
			}

			b.idx++

			continue
		}
	}

	if count == 0 {
		return 0, io.EOF
	}

	return count, nil
}

// NewBytesReader new multiple bytes reader.
func NewBytesReader(src ...[]byte) io.Reader {
	readers := make([]io.Reader, len(src))

	for i, b := range src {
		readers[i] = bytes.NewReader(b)
	}

	return NewQueueReader(readers...)
}

// NewQueueReader new queue reader.
func NewQueueReader(src ...io.Reader) io.Reader {
	return &queueReader{
		idx:     0,
		total:   len(src),
		readers: src,
	}
}
