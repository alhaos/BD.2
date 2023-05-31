package fs

import (
	"bufio"
	"bytes"
	"os"
)

// ReadLines return line array from file
func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)

	s.Split(scanLinesWithCR)

	var lines []string

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines, nil
}

// scanLinesWithCR handle lines that end in 0x0d, you can write your own split
// function based on bufio.ScanLines, which will also treat a carriage return as
// the end of a line:
func scanLinesWithCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}

	if i := bytes.IndexByte(data, '\r'); i >= 0 {
		// We have a line terminated by a carriage return.
		return i + 1, data[0:i], nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}
