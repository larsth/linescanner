package linescanner

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var (
	ErrNilIoReader               = errors.New("Nil io.Reader")
	ErrReaderCapacityLessThanOne = errors.New("Reader capacity < 1")
	ErrTokenCapacityLessThanOne  = errors.New("Token capacity < 1")
	ErrBufferCapacityLessThanOne = errors.New("Buffer capacity < 1")
)

func NewSize(r io.Reader, readerCapacity int, tokenCapacity int, bufferCapacity int) (ls *LineScanner, err error) {
	if r == nil {
		return nil, ErrNilIoReader
	}
	if readerCapacity < 1 {
		return nil, ErrReaderCapacityLessThanOne
	}
	if tokenCapacity < 1 {
		return nil, ErrTokenCapacityLessThanOne
	}
	if bufferCapacity < 1 {
		return nil, ErrBufferCapacityLessThanOne
	}

	ls = new(LineScanner)
	ls.r = bufio.NewReaderSize(r, readerCapacity)
	ls.token = make([]byte, 0, tokenCapacity)
	ls.buf = make([]byte, 0, bufferCapacity)
	ls.err = nil

	return ls, nil
}

func New(r io.Reader) (ls *LineScanner, err error) {
	pageSize := os.Getpagesize()
	return NewSize(r, pageSize, pageSize, pageSize)
}
