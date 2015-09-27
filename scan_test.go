package linescanner

import (
	"bytes"
	"io"
	"strconv"
	"testing"
)

func TestScan1(t *testing.T) {
	var (
		ls        *LineScanner
		gotBool   bool
		hadFailed bool
		s         string
	)
	ls = &LineScanner{}
	gotBool = ls.Scan()
	hadFailed, s = tstCheckBool("'returned bool'", gotBool, false) //false is the want bool
	if hadFailed == true {
		t.Error(s)
	}
}

type tstTScanWant struct {
	Bool  bool
	Err   error
	LsErr error
	Token []byte
	Buf   []byte
}

type tstTScan struct {
	Input               []byte
	WithRead            bool
	BufIoReaderCapacity int
	Want                tstTScanWant
}

var tdScanWant2 = []*tstTScan{
	&tstTScan{
		Input:               []byte{1},
		WithRead:            true,
		BufIoReaderCapacity: 16,
		Want: tstTScanWant{
			Bool:  false,
			LsErr: io.EOF,
			Err:   nil,
			Token: make([]byte, 0, 1),
			Buf:   make([]byte, 0, 1),
		},
	},
	&tstTScan{
		Input:               []byte("abcdefghijklmnopqrstuvxyz"),
		WithRead:            false,
		BufIoReaderCapacity: 16,
		Want: tstTScanWant{
			Bool:  false,
			LsErr: nil,
			Err:   nil,
			Token: make([]byte, 0, 1),
			Buf:   []byte("abcdefghijklmnop"),
		},
	},
	&tstTScan{
		Input:               []byte("abcdefghijklmnopqrst\nuvxyz"),
		WithRead:            false,
		BufIoReaderCapacity: 4096,
		Want: tstTScanWant{
			Bool:  true,
			LsErr: nil,
			Err:   nil,
			Token: []byte("abcdefghijklmnopqrst"),
			Buf:   make([]byte, 0, 1),
		},
	},
}

func TestScan2(t *testing.T) {
	var (
		p         []byte = make([]byte, 1)
		r         io.Reader
		testItem  *tstTScan
		ls        *LineScanner
		gotBool   bool
		hadFailed bool
		s         string
	)
	for _, testItem = range tdScanWant2 {
		r = bytes.NewReader(testItem.Input)
		if testItem.WithRead {
			_, _ = r.Read(p) // read one=len(p) byte
		}
		//Note that the ls.r of type *bufio.Reader will use size=16 as a minimum.
		ls, _ = NewSize(r, testItem.BufIoReaderCapacity, 4096, 4096)
		gotBool = ls.Scan()
		hadFailed, s = tstCheckBool("'returned bool'",
			gotBool, testItem.Want.Bool)
		if hadFailed == true {
			t.Error(s)
		}
		hadFailed, s = tstCheckErr("'*LineScanner.err'",
			ls.err, testItem.Want.LsErr)
		if hadFailed == true {
			t.Error(s)
		}
		hadFailed, s = tstCheckErr("'*LineScanner.err'",
			ls.Err(), testItem.Want.Err)
		if hadFailed == true {
			t.Error(s)
		}
		hadFailed, s = tstCheckByteSlice("'*LineScanner.token'",
			ls.token, testItem.Want.Token)
		if hadFailed == true {
			t.Error(s)
		}
		hadFailed, s = tstCheckByteSlice("'*LineScanner.buf'",
			ls.buf, testItem.Want.Buf)
		if hadFailed == true {
			t.Error(s)
		}
	}
}

var tdScanWant3 = []*tstTScan{
	//[0]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  true,
			LsErr: nil,
			Err:   nil,
			Token: []byte("abc"),
			Buf:   make([]byte, 0, 1),
		},
	},
	//[1]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  true,
			LsErr: nil,
			Err:   nil,
			Token: []byte("defghijk"),
			Buf:   make([]byte, 0, 1),
		},
	},
	//[2]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  true,
			LsErr: nil,
			Err:   nil,
			Token: []byte("lmno\rpqrst"),
			Buf:   make([]byte, 0, 1),
		},
	},
	//[3]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  false,
			LsErr: nil,
			Err:   nil,
			Token: []byte("lmno\rpqrst"),
			Buf:   []byte("uvxyzABCDEFGHIJK"),
		},
	},
	//[4]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  false,
			LsErr: nil,
			Err:   nil,
			Token: []byte("lmno\rpqrst"),
			Buf:   []byte("uvxyzABCDEFGHIJKLMNOPQRSTUVXYZ01"),
		},
	},
	//[5]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  true,
			LsErr: nil,
			Err:   nil,
			Token: []byte("uvxyzABCDEFGHIJKLMNOPQRSTUVXYZ012"),
			Buf:   make([]byte, 0, 1),
		},
	},
	//[6]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  false,
			LsErr: nil,
			Err:   nil,
			Token: []byte("uvxyzABCDEFGHIJKLMNOPQRSTUVXYZ012"),
			Buf:   []byte("34567890aAbBcCdD"),
		},
	},
	//[7]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  true,
			LsErr: nil,
			Err:   nil,
			Token: []byte("34567890aAbBcCdD"),
			Buf:   make([]byte, 0, 1),
		},
	},
	//[8]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  false,
			LsErr: nil,
			Err:   nil,
			Token: []byte("34567890aAbBcCdD"),
			Buf:   []byte("eEfFgGhHiIjJkKlL"),
		},
	},
	//[9]
	&tstTScan{
		Want: tstTScanWant{
			Bool:  true,
			LsErr: nil,
			Err:   nil,
			Token: []byte("eEfFgGhHiIjJkKlLmM"),
			Buf:   make([]byte, 0, 1),
		},
	},
}

func TestScan3(t *testing.T) {
	var (
		input               []byte
		bufIoReaderCapacity int = 16
		r                   io.Reader
		testItem            *tstTScan
		ls                  *LineScanner
		gotBool             bool
		hadFailed           bool
		s                   string
		i                   int = 0
	)
	input = append(input, []byte("abc\r\ndefghijk\r\nlmno\rpqrst\nuvxyz")...)
	input = append(input, []byte("ABCDEFGHIJKLMNOPQRSTUVXYZ")...)
	input = append(input, []byte("012\n34567890aAbBcCdD\r\neEfFgGhHiIjJkKlLmM\n")...)

	r = bytes.NewReader(input)
	//Note that the ls.r of type *bufio.Reader will use size=16 as a minimum.
	ls, _ = NewSize(r, bufIoReaderCapacity, 4096, 4096)

	for _, testItem = range tdScanWant3 {
		gotBool = ls.Scan()
		hadFailed, s = tstCheckBool("'returned bool'",
			gotBool, testItem.Want.Bool)
		if hadFailed == true {
			t.Error(s)
			t.Log("Failed test: 'tdScanWant3[", strconv.Itoa(i), "]'")
		}
		hadFailed, s = tstCheckErr("'*LineScanner.err'",
			ls.err, testItem.Want.LsErr)
		if hadFailed == true {
			t.Error(s)
			t.Log("Failed test: 'tdScanWant3[", strconv.Itoa(i), "]'")
		}
		hadFailed, s = tstCheckErr("'*LineScanner.err'",
			ls.Err(), testItem.Want.Err)
		if hadFailed == true {
			t.Error(s)
			t.Log("Failed test: 'tdScanWant3[", strconv.Itoa(i), "]'")
		}
		hadFailed, s = tstCheckByteSlice("'*LineScanner.token'",
			ls.token, testItem.Want.Token)
		if hadFailed == true {
			t.Error(s)
			t.Log("Failed test: 'tdScanWant3[", strconv.Itoa(i), "]'")
		}
		hadFailed, s = tstCheckByteSlice("'*LineScanner.buf'",
			ls.buf, testItem.Want.Buf)
		if hadFailed == true {
			t.Error(s)
			t.Log("Failed test: 'tdScanWant3[", strconv.Itoa(i), "]'")
		}
		i++
	}
}

func BenchmarkScan(b *testing.B) {
	var (
		ls     *LineScanner
		input  = []byte("Hello World\r\nHej Verden\n")
		reader = bytes.NewReader(input)
		r      = io.Reader(reader)
	)
	for i := 0; i < b.N; i++ {
		ls, _ = NewSize(r, 16, 16, 16)
		_ = ls.Scan()
	}
}

func BenchmarkScanWithOutNew(b *testing.B) {
	var (
		ls     *LineScanner
		input  = []byte("Hello World\r\nHej Verden\n")
		reader = bytes.NewReader(input)
		r      = io.Reader(reader)
	)
	ls, _ = NewSize(r, 16, 16, 16)
	for i := 0; i < b.N; i++ {
		_ = ls.Scan()
	}
}
