package linescanner

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNewSize1(t *testing.T) {
	var (
		ls        *LineScanner = nil
		r         io.Reader    = nil
		gotErr    error        = nil
		hadFailed bool
		s         string
	)

	ls, gotErr = NewSize(r, 16, 16, 16)

	hadFailed, s = tstCheckErr("'returned error'", gotErr, ErrNilIoReader)
	if hadFailed == true {
		t.Error(s)
	}
	if ls != nil {
		t.Error("Got a *LineScanner, which is != nil, but want a *LineScanner, which is == nil.")
	}
}

func TestNewSize2(t *testing.T) {
	var (
		ls        *LineScanner = nil
		input                  = []byte("abcdefghi\n")
		r                      = io.Reader(bytes.NewReader(input))
		gotErr    error        = nil
		hadFailed bool
		s         string
	)

	ls, gotErr = NewSize(r, 0, 16, 16)

	hadFailed, s = tstCheckErr("'returned error'", gotErr, ErrReaderCapacityLessThanOne)
	if hadFailed == true {
		t.Error(s)
	}
	if ls != nil {
		t.Error("Got a *LineScanner, which is != nil, but want a *LineScanner, which is == nil.")
	}
}

func TestNewSize3(t *testing.T) {
	var (
		ls        *LineScanner = nil
		input                  = []byte("abcdefghi\n")
		r                      = io.Reader(bytes.NewReader(input))
		gotErr    error        = nil
		hadFailed bool
		s         string
	)

	ls, gotErr = NewSize(r, 16, 0, 16)

	hadFailed, s = tstCheckErr("'returned error'", gotErr, ErrTokenCapacityLessThanOne)
	if hadFailed == true {
		t.Error(s)
	}
	if ls != nil {
		t.Error("Got a *LineScanner, which is != nil, but want a *LineScanner, which is == nil.")
	}
}

func TestNewSize4(t *testing.T) {
	var (
		ls        *LineScanner = nil
		input                  = []byte("abcdefghi\n")
		r                      = io.Reader(bytes.NewReader(input))
		gotErr    error        = nil
		hadFailed bool
		s         string
	)

	ls, gotErr = NewSize(r, 16, 16, 0)

	hadFailed, s = tstCheckErr("'returned error'", gotErr, ErrBufferCapacityLessThanOne)
	if hadFailed == true {
		t.Error(s)
	}
	if ls != nil {
		t.Error("Got a *LineScanner, which is != nil, but want a *LineScanner, which is == nil.")
	}
}

func TestNewSize5(t *testing.T) {
	var (
		ls        *LineScanner = nil
		input                  = []byte("abcdefghi\n")
		r                      = io.Reader(bytes.NewReader(input))
		gotErr    error        = nil
		hadFailed bool
		s         string
	)

	ls, gotErr = NewSize(r, 16, 32, 64)

	hadFailed, s = tstCheckErr("'returned error'", gotErr, error(nil))
	if hadFailed == true {
		t.Error(s)
	}
	if ls == nil {
		t.Error("Got a *LineScanner, which is == nil, but want a *LineScanner, which is != nil.")
		t.Skip("Because the *LineScanner is nil, the follwing tests will " +
			"crash the test binary, so the rest of this test is aborted.")
	}
	if ls.r == nil {
		t.Error("Got a *LineScanner.r (of type *bufio.Reader), which is == nil, " +
			"but want a *LineScanner.r, which is != nil.")
	}
	if ls.token == nil {
		t.Error("Got a *LineScanner.token (of type []byte), which is == nil, " +
			"but want a *LineScanner.token, which is != nil.")
	} else {
		if len(ls.token) > 0 {
			t.Error("Got a *LineScanner.token (of type []byte), which has " +
				"a length > 0, but want a length equal to 0 (zero).")
		}
		if cap(ls.token) != 32 {
			t.Error("Got a *LineScanner.token (of type []byte), which has " +
				"a capacity != 32, but want a capacity equal to 32.")
		}
	}
	if ls.buf == nil {
		t.Error("Got a *LineScanner.buf (of type []byte), which is == nil, " +
			"but want a *LineScanner.buf, which is != nil.")
	} else {
		if len(ls.buf) > 0 {
			t.Error("Got a *LineScanner.buf (of type []byte), which has " +
				"a length > 0, but want a length equal to 0 (zero).")
		}
		if cap(ls.buf) != 64 {
			t.Error("Got a *LineScanner.buf (of type []byte), which has " +
				"a capacity != 64, but want a capacity equal to 64.")
		}
	}
}

func TestNew1(t *testing.T) {
	var (
		ls        *LineScanner = nil
		r         io.Reader    = nil
		gotErr    error        = nil
		hadFailed bool
		s         string
	)

	ls, gotErr = New(r)

	hadFailed, s = tstCheckErr("'returned error'", gotErr, ErrNilIoReader)
	if hadFailed == true {
		t.Error(s)
	}
	if ls != nil {
		t.Error("Got a *LineScanner, which is != nil, but want a *LineScanner, which is == nil.")
	}
}

func TestNew5(t *testing.T) {
	var (
		ls        *LineScanner = nil
		input                  = []byte("abcdefghi\n")
		r                      = io.Reader(bytes.NewReader(input))
		gotErr    error        = nil
		hadFailed bool
		s         string
	)

	ls, gotErr = New(r)

	hadFailed, s = tstCheckErr("'returned error'", gotErr, error(nil))
	if hadFailed == true {
		t.Error(s)
	}
	if ls == nil {
		t.Error("Got a *LineScanner, which is == nil, but want a *LineScanner, which is != nil.")
		t.Skip("Because the *LineScanner is nil, the follwing tests will " +
			"crash the test binary, so the rest of this test is aborted.")
	}
	if ls.r == nil {
		t.Error("Got a *LineScanner.r (of type *bufio.Reader), which is == nil, " +
			"but want a *LineScanner.r, which is != nil.")
	}
	if ls.token == nil {
		t.Error("Got a *LineScanner.token (of type []byte), which is == nil, " +
			"but want a *LineScanner.token, which is != nil.")
	} else {
		if len(ls.token) > 0 {
			t.Error("Got a *LineScanner.token (of type []byte), which has " +
				"a length > 0, but want a length equal to 0 (zero).")
		}
		if cap(ls.token) != os.Getpagesize() {
			t.Error("Got a *LineScanner.token (of type []byte), which has " +
				"a capacity != (os.Getpagesize()), but want a capacity equal " +
				"to (os.Getpagesize()).")
		}
	}
	if ls.buf == nil {
		t.Error("Got a *LineScanner.buf (of type []byte), which is == nil, " +
			"but want a *LineScanner.buf, which is != nil.")
	} else {
		if len(ls.buf) > 0 {
			t.Error("Got a *LineScanner.buf (of type []byte), which has " +
				"a length > 0, but want a length equal to 0 (zero).")
		}
		if cap(ls.buf) != os.Getpagesize() {
			t.Error("Got a *LineScanner.buf (of type []byte), which has " +
				"a capacity != (os.Getpagesize()), but want a capacity equal " +
				"to (os.Getpagesize()).")
		}
	}
}

func BenchmarkNew(b *testing.B) {
	var (
		input  = []byte("Hello World\r\nHej Verden\n")
		reader = bytes.NewReader(input)
		r      = io.Reader(reader)
	)
	for i := 0; i < b.N; i++ {
		_, _ = New(r)
	}
}

func BenchmarkNewSize(b *testing.B) {
	var (
		input  = []byte("Hello World\r\nHej Verden\n")
		reader = bytes.NewReader(input)
		r      = io.Reader(reader)
	)
	for i := 0; i < b.N; i++ {
		_, _ = NewSize(r, 16, 16, 16)
	}
}
