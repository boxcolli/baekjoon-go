package main

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Provide example of structured input data */
var readResult = Input{
	name: "World",
}

// Compare read() result with given example
func TestRead(t *testing.T) {
	assert := assert.New(t)

	{
		inputWriter, f := mockStdin()
		defer f()

		in, _ := openAndReadSingleTestCase()
		t.Log("input: ", in)
		inputWriter.Write([]byte(in))
		inputWriter.Close()
	}

	input := read()
	t.Log("input: ", input)

	assert.Equal(true, reflect.DeepEqual(input, readResult))
}

func openAndReadSingleTestCase() (in, out string) {
	// Open testcase file
	f, err := os.Open(TestCaseFileName)
	if err != nil {
		panic(err)
	}
	
	// 
	in, out, err = readSingleTestCase(bufio.NewReader(f))
	if err != nil && err != io.EOF {
		panic(err)
	}

	f.Close()

	return
}