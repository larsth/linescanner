package linescanner

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestLineScannerBytes(t *testing.T) {
	var (
		slice       = []byte("Hello World")
		r           = io.Reader(bytes.NewReader(slice))
		lc          *LineScanner
		got         []byte
		want        []byte = slice
		sGot, sWant string
	)
	lc, _ = New(r)
	lc.token = append(lc.token, want...)

	got = lc.Bytes()

	if bytes.Compare(want, got) != 0 {
		sGot = fmt.Sprintf("%s", got)
		sWant = fmt.Sprintf("%s", want)
		t.Errorf("Got text '%s'\nWant text '%s'", sGot, sWant)
	}
}

func BenchmarkBytes(b *testing.B) {
	var (
		ls     *LineScanner
		input  = []byte("Hello World\r\nHej Verden\n")
		reader = bytes.NewReader(input)
		r      = io.Reader(reader)
	)
	ls, _ = New(r)
	_ = ls.Scan()
	for i := 0; i < b.N; i++ {
		_ = ls.Bytes()
	}
}
