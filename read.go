package main

import "fmt"

type Input struct {
	/* Define customized data */
	name string
}

// Read from stdin and return structured data
func read() (input Input) {
	/* Define how to structure input data */
	fmt.Scan(&input.name)

	return
}