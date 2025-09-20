// Declare that this file belongs to the "calculator_test" package.
// This convention keeps the tests in a separate package from the code under test.
package calculator_test

import (
	"testing"

	"github.com/brettfirecore/calculator" // the package we’re testing
	// Go's standard testing framework
)

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
		{a: 2, b: 2, want: 4}, // 2 + 2 should equal 4
		{a: 1, b: 1, want: 2}, // 1 + 1 should equal 2
		{a: 5, b: 0, want: 5}, // 5 + 0 should equal 5
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
		got := calculator.Add(tc.a, tc.b) // call function with the testCase inputs
		if tc.want != got {               // compare expected vs actual
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel() // run test in parallel with others

	var want float64 = 2
	got := calculator.Subtract(4, 2) // Call the Subtract function

	if want != got {
		t.Errorf("want %f, got %f", want, got)
		// Reports a failure if the result isn't as expected
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel() // run in parallel with other tests

	var want float64 = 9
	got := calculator.Multiply(3, 3) // 3 × 3 should equal 9

	if got != want {
		t.Fatalf("got %f want %f", got, want)
		// t.Fatalf marks the test as failed and stops it immediately
	}
}
