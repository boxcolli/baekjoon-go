package main

import (
	"os"
)

// Returns input writer and defer function
func mockStdin() (*os.File, func()) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdin := os.Stdin
	os.Stdin = r
	
	return w, func() { os.Stdin = stdin }
}

// Returns output reader and defer function
func mockStdout() (*os.File, *os.File, func()) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w
	
	return r, w, func() {
		os.Stdout = stdout
	}
}
