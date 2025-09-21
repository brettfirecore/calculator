// Declare that this file belongs to the "calculator_test" package.
// This convention keeps the tests in a separate package from the code under test.
package calculator_test

import (
	"testing"

	"github.com/brettfirecore/calculator" // the package we’re testing
	// Go's standard testing framework
)

// TestAdd checks the Add function with multiple cases
func TestAdd(t *testing.T) {
	t.Parallel() // Allow this test to run in parallel with others

	// Define a helper struct to represent one test case
	type testCase struct {
		a, b float64 // inputs
		want float64 // expected output
	}

	// Create a slice (dynamic array) of testCase values.
	// []testCase means "a slice of testCase".
	// := assigns it to the variable testCases.
	testCases := []testCase{ // [] <--slice
		// Each element is a struct literal with fields a, b, and want.
		{a: 2, b: 2, want: 4}, // 2 + 2 = 4
		{a: 1, b: 1, want: 2}, // 1 + 1 = 2
		{a: 5, b: 0, want: 5}, // 5 + 0 = 5
	}

	// Loop over every element in the slice testCases
	// range testCases gives us two values each time:
	//  1. the index (0, 1, 2, …)
	//  2. the element itself (a testCase struct)
	//
	// The underscore (_) means "ignore this value" —
	// here, we don’t care about the index, only the element.
	// tc will hold each testCase as we go through the slice.
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b) // call Add with inputs from tc
		if tc.want != got {               // compare expected vs actual
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got) // fail the test with details
		}
	}
}

// TestSubtract checks the Subtract function
func TestSubtract(t *testing.T) {
	t.Parallel() // run in parallel with other tests

	// Define a helper struct for test cases
	type testCase struct {
		a, b float64 // inputs
		want float64 // expected output
	}

	// Slice of test cases for Subtract
	testCases := []testCase{
		{a: 2, b: 2, want: 0},  // 2 - 2 = 0
		{a: 2, b: 1, want: 1},  // 2 - 1 = 1
		{a: 5, b: -4, want: 9}, // 5 - (-4) = 9
	}

	// Loop through each testCase
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b) // call Subtract
		if tc.want != got {                    // compare want vs got
			t.Errorf("Subtract(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got) // report mismatch
		}
	}
}

// TestMultiply checks the Multiply function
func TestMultiply(t *testing.T) {
	t.Parallel() // run in parallel with other tests

	// Define a helper struct for test cases
	type testCase struct {
		a, b float64 // inputs
		want float64 // expected output
	}

	// Slice of test cases for Multiply
	testCases := []testCase{
		{a: 2, b: 2, want: 4},   // 2 × 2 = 4
		{a: -1, b: -1, want: 1}, // (-1) × (-1) = 1
		{a: 5, b: 0, want: 0},   // 5 × 0 = 0
	}

	// Loop through each testCase
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b) // call Multiply
		if tc.want != got {                    // compare want vs got
			t.Errorf("Multiply(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got) // report mismatch
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel() // Allow this test to run in parallel with others

	// Define a helper struct to represent one test case
	type testCase struct {
		a, b float64 // inputs
		want float64 // expected result of a / b
	}

	// Table of valid test cases (no divide-by-zero here)
	testCases := []testCase{
		{a: 2, b: 2, want: 1},   // 2 / 2 = 1
		{a: -1, b: -1, want: 1}, // (-1) / (-1) = 1
		{a: 10, b: 2, want: 5},  // 10 / 2 = 5
	}

	// Loop over every test case
	for _, tc := range testCases {
		// Call Divide: got is the numeric result, err is the error (should be nil here)
		got, err := calculator.Divide(tc.a, tc.b)
		// First check: make sure no error occurred for valid inputs
		if err != nil {
			t.Fatalf("Divide(%f, %f) returned unexpected error: %v",
				tc.a, tc.b, err)
		}

		// Second check: compare the result with what we expected
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

// TestDivideInvalid checks the Divide function with invalid input behaviour
func TestDivideInvalid(t *testing.T) {
	t.Parallel() // run this test in parallel

	// Attempt to divide by zero: this is invalid input.
	// We don’t care about the numeric result (so we use the blank identifier _),
	// but we *do* expect a non-nil error.
	_, err := calculator.Divide(1, 0)

	// If err is nil, that means Divide failed to signal an error,
	// so the test fails.
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}
