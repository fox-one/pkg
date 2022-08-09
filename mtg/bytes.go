package mtg

import (
	"errors"
	"io"
)

func NewReader(b []byte) *Reader {
	return &Reader{b: b}
}

type Reader struct {
	b      []byte
	offset int
}

func (r *Reader) ReadByte() (byte, error) {
	if r.offset >= len(r.b) {
		return 0, io.EOF
	}

	b := r.b[r.offset]
	r.offset++
	return b, nil
}

func (r *Reader) Read(n int) ([]byte, error) {
	if r.offset >= len(r.b) {
		return nil, io.EOF
	}

	if r.offset+n > len(r.b) {
		return nil, errors.New("cannot read more than available")
	}

	b := r.b[r.offset : r.offset+n]
	r.offset += n

	return b, nil
}

func (r *Reader) ReadAll() ([]byte, error) {
	b := r.b[r.offset:]
	r.offset = len(r.b)
	return b, nil
}
