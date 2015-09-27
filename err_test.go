package linescanner

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestLineScannerErrNotIoEOF(t *testing.T) {
	var (
		slice   = []byte("Hello World")
		r       = io.Reader(bytes.NewReader(slice))
		ls      *LineScanner
		gotErr  error
		wantErr error = errors.New("test")
	)
	ls, _ = New(r)
	ls.err = wantErr
	gotErr = ls.Err()

	sGotErr := fmt.Sprintf("%s", gotErr)
	sWantErr := fmt.Sprintf("%s", wantErr)
	if strings.Compare(sWantErr, sGotErr) != 0 {
		t.Errorf("Got error '%s'\nWant '%s'", sGotErr, sWantErr)
	}
}

func TestLineScannerErrIsIoEOF(t *testing.T) {
	var (
		slice   = []byte("Hello World")
		r       = io.Reader(bytes.NewReader(slice))
		ls      *LineScanner
		gotErr  error
		wantErr error = nil
	)
	ls, _ = New(r)
	ls.err = io.EOF
	gotErr = ls.Err()

	sGotErr := fmt.Sprintf("%s", gotErr)
	sWantErr := fmt.Sprintf("%s", wantErr)
	if strings.Compare(sWantErr, sGotErr) != 0 {
		t.Errorf("Got error '%s'\nWant '%s'", sGotErr, sWantErr)
	}
}

func TestLineScannerErrIsnil(t *testing.T) {
	var (
		slice   = []byte("Hello World")
		r       = io.Reader(bytes.NewReader(slice))
		ls      *LineScanner
		gotErr  error
		wantErr error = nil
	)
	ls, _ = New(r)
	ls.err = nil
	gotErr = ls.Err()

	sGotErr := fmt.Sprintf("%s", gotErr)
	sWantErr := fmt.Sprintf("%s", wantErr)
	if strings.Compare(sWantErr, sGotErr) != 0 {
		t.Errorf("Got error '%s'\nWant '%s'", sGotErr, sWantErr)
	}
}

func BenchmarkErr(b *testing.B) {
	var (
		ls     *LineScanner
		input  = []byte("Hello World\r\nHej Verden\n")
		reader = bytes.NewReader(input)
		r      = io.Reader(reader)
	)
	ls, _ = New(r)
	for i := 0; i < b.N; i++ {
		_ = ls.Err()
	}
}
