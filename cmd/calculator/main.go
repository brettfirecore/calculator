package main

import "fmt"

func divide(a, b int) (quotient int, remainder int, err error) {
	if b == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return
	}
	quotient = a / b
	remainder = a % b
	return // naked return: returns quotient, remainder, err
}

func main() {
	q, r, err := divide(17, 5)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
}
