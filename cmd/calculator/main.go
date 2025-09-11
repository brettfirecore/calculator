package main

import (
	"fmt"
	"os"
	// "github.com/brettfirecore/calculator/calculator"
)

// build-time fields (must come AFTER the import block)
var (
	version   = "dev" // overridden by -ldflags: -X 'main.version=...'
	buildDate = ""    // overridden by -ldflags: -X 'main.buildDate=...'
)

func main() {
	// Version/build info
	fmt.Printf("calculator version=%s buildDate=%s\n", version, buildDate)

	// 👉 Uncomment one block at a time to experiment

	// --- 1. Function declaration + call
	// fmt.Println("add(2,3):", add(2, 3))

	// --- 2. Early return
	// if res, err := safeDivide(10, 0); err != nil {
	// 	fmt.Println("safeDivide error:", err)
	// } else {
	// 	fmt.Println("safeDivide(10,2):", res)
	// }

	// --- 3. Functions as values
	// result := operate(2, 3, add)
	// fmt.Println("operate(2,3,add):", result)

	// --- 4. Closure
	// double := makeMultiplier(2)
	// fmt.Println("double(5):", double(5))

	// --- 5. Defer cleanup
	// if err := writeFileDemo(); err != nil {
	// 	fmt.Println("writeFileDemo error:", err)
	// } else {
	// 	fmt.Println("writeFileDemo wrote demo.txt")
	// }

	// --- 6. Variadic function
	// fmt.Println("sumAll(1,2,3,4):", sumAll(1, 2, 3, 4))

	// --- 7. Using math (remember to import "math" if you uncomment)
	// res := math.Sqrt(16)
	// fmt.Println("math.Sqrt(16):", res)
}

// 1. Simple function declaration
func add(a, b int) int {
	return a + b
}

// 2. Early return with error
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 3. Functions as values
func operate(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

// 4. Closure
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 5. Defer for cleanup (needs import "os")
func writeFileDemo() (err error) {
	f, err := os.Create("demo.txt")
	if err != nil {
		return
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			if err == nil {
				err = cerr
			}
		}
	}()
	_, err = f.WriteString("Hello, Go!\n")
	return
}

// 6. Variadic function
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
