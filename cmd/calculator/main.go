package main

import (
	"errors"
	"fmt"
	"os"
)

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
	// Instead of hardcoding testdata/out.txt:
	path := "out.txt"

	if err := writeFile(path, []byte("Hello from Go!\n")); err != nil {
		fmt.Println("writeFile error:", err)
		return
	}
	fmt.Println("writeFile succeeded, wrote to", path)
}
