package main

import (
	"errors"
	"fmt"
	"os"
)

// writeFile writes data and ensures any error from f.Close()
// can update the returned error. Named result: err.
func writeFile(path string, data []byte) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return // early return if file couldn't be created
	}

	// Defer a closure that captures `err`.
	defer func() {
		// Force an artificial close error for demonstration:
		// (In real life, f.Close() may fail e.g. on full disk)
		closeErr := errors.New("simulated close failure")

		// If f.Close() fails and no prior error, overwrite err.
		if closeErr != nil {
			if err == nil {
				err = closeErr
			} else {
				// If both write and close failed, join them (Go 1.20+).
				err = errors.Join(err, closeErr)
			}
		}
	}()

	// Do the actual write
	if _, werr := f.Write(data); werr != nil {
		err = werr // assign to the named result, not shadowed
		return
	}

	return // naked return: err may still be updated by the defer
}

func main() {
	if err := writeFile("testdata/out.txt", []byte("hello\n")); err != nil {
		fmt.Println("writeFile error:", err)
		return
	}
	fmt.Println("writeFile succeeded")
}
