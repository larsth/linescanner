package linescanner

import "io"

//Scan reads a line, and returns true when a complete line had been found.
//If Scan did't find the end of line characters ('\n' or '\r\n'), and Scan
//returns false.
//Any previous token (line), and error can be read as long as Scan is returning
// false.
//A token is read via the Bytes() and Text methods, and the error is read via
//the Err() method.
//Scan supports reading sequential lines, but no indication or error is given
//if the input ends without a final line end.
func (ls *LineScanner) Scan() (hasToken bool) {
	var (
		line     []byte
		isPrefix bool
		err      error
	)

	ls.mutex.Lock()
	defer ls.mutex.Unlock()

	if ls.r == nil {
		return false
	}

	line, isPrefix, err = ls.r.ReadLine()
	//From the documentation about bufio.Reader.ReadLine:
	//ReadLine tries to return a single line, not including the end-of-line
	//bytes.
	//The text returned from ReadLine does not include the line end ("\r\n" or
	// "\n").
	//If the line was too long for the buffer then isPrefix is set and the
	//beginning of the line is returned.
	//The rest of the line will be returned from future calls.
	//isPrefix will be false when returning the last fragment of the line.
	//The returned buffer is only valid until the next call to ReadLine.
	//ReadLine either returns a non-nil line or it returns an error, never both.
	//No indication or error is given if the input ends without a final line end.

	hasToken = false
	if line != nil {
		ls.err = nil
		if isPrefix == false {
			hasToken = true
			ls.readCount = uint(0)
			//clear the token slice:
			ls.token = ls.token[0:0:cap(ls.token)]
			//Append the buffer slice, then the line slice, to the token slice:
			ls.token = append(ls.token, ls.buf...)
			ls.token = append(ls.token, line...)
			//clear the buffer slice:
			ls.buf = ls.buf[0:0:cap(ls.buf)]
		} else {
			//Append the line slice to the buffer slice:
			ls.buf = append(ls.buf, line...)
		}
	}
	if err != nil {
		ls.err = err
	}
	ls.hasEof = (err == io.EOF)
	return
}
