//author: wongoo
//date: 20190814

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

var (
	// ErrTimeout is returned from timed out calls.
	ErrTimeout = errors.New("timeout")
)

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
			if err == io.EOF {
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

// TimeoutCopy copy in a timeout duration
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
