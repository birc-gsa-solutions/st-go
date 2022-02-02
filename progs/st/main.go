package main

import (
	"fmt"
	"os"

	// Directories in the root of the repo can be imported
	// as long as we pretend that they sit relative to the
	// url birc.au.dk/gsa, like this for the example 'shared':
	"birc.au.dk/gsa/shared"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	genome := os.Args[1]
	reads := os.Args[2]
	fmt.Println(shared.Todo(genome, reads))
}
