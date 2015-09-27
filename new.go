package linescanner

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var (
	//ErrNilIoReader is an error that tells the calling function that it is
	//using an io.Reader, which is nil
	ErrNilIoReader = errors.New("Nil io.Reader")

	//ErrReaderCapacityLessThanOne is an error that tells the calling
	//function that it is using a bufio.Reader capacity less than the
	//bufio.Reader minimum capacity, which is 16 (according to the source code
	//for the bufio.Reader.NewSize function).
	ErrReaderCapacityLessThanOne = errors.New("Reader capacity < 16")

	//ErrTokenCapacityLessThanOne is an error that tells the calling
	//function that it is using a token capacity < 1
	ErrTokenCapacityLessThanOne = errors.New("Token capacity < 1")

	//ErrBufferCapacityLessThanOne is an error that tells the calling
	//function that it is using a buffer capacity < 1
	ErrBufferCapacityLessThanOne = errors.New("Buffer capacity < 1")
)

//NewSize creates a LineScanner and returns a pointer to it, if the provided
//io.Redader is not nil, if the readerCapacity is >= 16, if the token capacity
// is >= 1, and if te buffer capacity is >= 1.
//If the conditions are met, then the returned error is nil.
//If the conditions are not met a nil *LineScanner pointer and an error != nil
//are returned.
func NewSize(r io.Reader, readerCapacity int, tokenCapacity int, bufferCapacity int) (ls *LineScanner, err error) {
	if r == nil {
		return nil, ErrNilIoReader
	}
	if readerCapacity < 16 {
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

//New creates a LineScanner and returns a pointer to it, if the provided
//io.Reader is not nil. If the provided io.Reader is not nil a nil error is also
//returned.
//If the io.Reader is nil, then a nil *Linescanner, and the ErrNilIoReader error
//are returned.
//All capacities (reader, token and buffer) will be set to os.Getpagesize,
//which will usually return 4096 (=a page is 4 KiB).
func New(r io.Reader) (ls *LineScanner, err error) {
	pageSize := os.Getpagesize()
	return NewSize(r, pageSize, pageSize, pageSize)
}
