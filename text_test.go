package linescanner

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestLineScannerText(t *testing.T) {
	var (
		slice   = []byte("Hello World")
		r       = io.Reader(bytes.NewReader(slice))
		ls      *LineScanner
		gotTxt  string
		wantTxt string = "Hello World"
	)
	ls, _ = New(r)
	ls.token = append(ls.token, []byte("Hello World")...)

	gotTxt = ls.Text()
	if strings.Compare(wantTxt, gotTxt) != 0 {
		t.Errorf("Got text '%s'\nWant text '%s'", gotTxt, wantTxt)
	}
}

func BenchmarkText(b *testing.B) {
	var (
		ls     *LineScanner
		input  = []byte("Hello World\r\nHej Verden\n")
		reader = bytes.NewReader(input)
		r      = io.Reader(reader)
	)
	ls, _ = New(r)
	_ = ls.Scan()
	for i := 0; i < b.N; i++ {
		_ = ls.Text()
	}
}
