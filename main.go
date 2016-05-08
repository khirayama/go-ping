package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "arg error")
		os.Exit(1)
	}
	// host := os.Args[1]
}
