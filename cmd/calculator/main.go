package main

import (
	"errors"
	"fmt"
	"os"
)

// writeFile writes some data to a file, and ensures that
// any error from closing the file is also captured and returned.
func writeFile(path string, data []byte) (err error) {
	// Create (or truncate) the file
	f, err := os.Create(path)
	if err != nil {
		return // early return if file can't be created
	}

	// Defer a cleanup function that closes the file.
	// This can *modify* the named result variable `err`
	// before the function actually returns.
	defer func() {
		if cerr := f.Close(); cerr != nil {
			if err == nil {
				// If no error so far, return the close error
				err = cerr
			} else {
				// If both writing and closing failed, join errors (Go 1.20+)
				err = errors.Join(err, cerr)
				// For older Go versions, use fmt.Errorf:
				// err = fmt.Errorf("%w; also close error: %v", err, cerr)
			}
		}
	}()

	// Try writing the data to the file
	if _, werr := f.Write(data); werr != nil {
		// Assign to named result `err` (not := to avoid shadowing!)
		err = werr
		return
	}

	// Naked return — returns current value of `err`
	// (nil if no error happened, or a real error if set above).
	return
}

func main() {
	// Call writeFile and check for error
	if err := writeFile("testdata/out.txt", []byte("hello\n")); err != nil {
		fmt.Println("writeFile error:", err)
		return
	}
	fmt.Println("writeFile succeeded")
}
