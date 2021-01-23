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
