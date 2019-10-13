package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "No file(s) specified")
		os.Exit(1)
	}

	for _, filename := range os.Args[1:] {
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling lstat on %v: %v\n", filename, err)
			continue
		}

		if !(fileInfo.Mode()&os.ModeSymlink != 0) {
			fmt.Fprintf(os.Stderr, "%v is not a symlink; skipping\n", filename)
			continue
		}
	}
}