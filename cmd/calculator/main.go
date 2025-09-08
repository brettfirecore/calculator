package main

import (
	"errors"
	"fmt"
	"os"
)

// writeFile writes data and lets a real f.Close() error update the returned err.
func writeFile(path string, data []byte) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}

	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err == nil {
				err = closeErr
			} else {
				// Go 1.20+: keep both errors
				err = errors.Join(err, closeErr)
			}
		}
	}()

	if _, werr := f.Write(data); werr != nil {
		err = werr
		return
	}
	return
}

func main() {
	if err := writeFile("testdata/out.txt", []byte("Hello from Go!\n")); err != nil {
		fmt.Println("writeFile error:", err)
		return
	}
	fmt.Println("writeFile succeeded")
}
