package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "testdata/somefile.txt"

	// Step 1: Read and print the file line by line
	fmt.Println("Current file contents:")
	if err := printFile(path); err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	// Step 2: Append a new line
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening file for append:", err)
		return
	}
	defer f.Close()

	newLine := "Appended from Go!\n"
	if _, err := f.WriteString(newLine); err != nil {
		fmt.Println("error writing to file:", err)
		return
	}
	fmt.Println("✔ Wrote new line:", newLine)

	// Step 3: Show the updated file
	fmt.Println("\nUpdated file contents:")
	if err := printFile(path); err != nil {
		fmt.Println("error reading file:", err)
	}
}

// Helper function: read a file line by line and print it
func printFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}
