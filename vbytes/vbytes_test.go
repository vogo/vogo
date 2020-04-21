// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vbytes

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	err := TimeoutCopy(w, r, 3*time.Second)
	assert.Equal(t, ErrTimeout, err)

	r.readCount = 0
	err = TimeoutCopy(w, r, 10*time.Second)
	assert.Nil(t, err)
}
