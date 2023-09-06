package main

import (
	"testing"
	"bufio"
	"io"
	"os"
)

const TestCaseFileName = "testcase.txt"

//	readSingleTestCase() reads exactly single test case.
//	This function should:
//		- completely read until the end of single test case.
//		- always return EOF if met with.
//		- for other errors, do panic.
func readSingleTestCase(r *bufio.Reader) (in, out string, err error) {
	/* You can customize how readSingleTestCase() will scan the testcases. */
	
	for {
		var buf string
		{
			// Using r.ReadString('\n') will include trailing newline character.
			buf, err = r.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return
				} else {
					panic(err)
				}
			}
		}

		// Separation of input and output
		if buf == "/\n" {
			// read output
			out, err = r.ReadString('\n')
			if err != nil && err != io.EOF {
				panic(err)
			}

			// read until the end of this test
			_, err = r.ReadString('\n')
			if err != nil && err != io.EOF {
				panic(err)
			}

			return
		} else {
			in += buf
		}
	}
}

func TestReadSingleTestCase(t *testing.T) {
	/*
		To see the log, execute (only visible when fails):
			$ go test -v -run TestReadSingleTestCase
	*/

	// Open testcase file
	f, err := os.Open(TestCaseFileName)
	if err != nil {
		panic(err)
	}

	// Read single test case
	in, out, err := readSingleTestCase(bufio.NewReader(f))
	if err != nil && err != io.EOF {
		panic(err)
	}

	t.Log("in: ", in)
	t.Log("out: ", out)

	f.Close()
}

