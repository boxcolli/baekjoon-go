package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type addTest struct {
	in, out string
}

var testSuite []addTest

func TestMainProgram(t *testing.T) {
	openAndReadTestSuite()

	for index, testcase := range testSuite {
		t.Run(fmt.Sprintf("test %v", index), func(t *testing.T) {
			assert := assert.New(t)

			inputWriter, inputFunc := mockStdin()
			defer inputFunc()
			outputReader, outputWriter, outputFunc := mockStdout()
			defer outputFunc()

			inputWriter.Write([]byte(testcase.in))
			inputWriter.Close()
			
			main()

			outputWriter.Close()
			out, _ := io.ReadAll(outputReader)

			splitInput := splitLines(testcase.in)
			splitExpected := splitLines(testcase.out)
			splitActual := splitLines(string(out))
			t.Log("input:    ", splitInput)
			t.Log("expected: ", splitExpected)
			t.Log("actual:   ", splitActual)
			assert.Equal(true, reflect.DeepEqual(
				splitExpected,
				splitActual,
			))
		})
	}
}

func openAndReadTestSuite() {
	// Open testcase file
	f, err := os.Open(TestCaseFileName)
	if err != nil {
		panic(err)
	}
	
	// Read test cases
	r := bufio.NewReader(f)
	for {
		in, out, err := readSingleTestCase(r)
		if in != "" {
			testSuite = append(testSuite, addTest{in: in, out: out})
		}

		if err == io.EOF {
			break
		}
	}

	f.Close()
}

func splitLines(s string) ([]string) {
	// this function splits lines and drops every newline.
	return strings.Fields(s)
}

