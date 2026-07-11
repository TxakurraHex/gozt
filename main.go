package main

import (
	"fmt"
	"gozt/source"
	"os"
)

func main() {
	sources := []source.Source {
		source.DpkgSource{},
		// TODO: Add implementations as they're created
	}

	exitCode := 0

	for _, s := range sources {
		entries, err := s.Collect()
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: %s: %v\n", s.Name(), err)
			exitCode = 1 // Partial failure,`` still print whatever succeeds
		}
		for _, e := range entries {
			fmt.Println(e.String())
		}
	}

	os.Exit(exitCode)
}