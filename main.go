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

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Fprintf(os.Stderr, "%v: removes target symbolic links and the files they point to\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Will do nothing if used on a file that is not a symbolic link\n")
		os.Exit(0)
	}

	for _, filename := range os.Args[1:] {
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if fileInfo.Mode()&os.ModeSymlink == 0 {
			fmt.Fprintf(os.Stderr, "%v is not a symlink; skipping\n", filename)
			continue
		}

		realName, err := os.Readlink(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		err = os.Remove(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		err = os.Remove(realName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}
