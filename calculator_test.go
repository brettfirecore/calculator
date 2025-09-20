// Declare that this file belongs to the "calculator_test" package.
// This convention keeps the tests in a separate package from the code under test.
package calculator_test

import (
	"testing" // Go's standard testing framework
	"github.com/brettfirecore/calculator" // the package we’re testing
)

func TestAdd(t *testing.T) {
	t.Parallel() // Allow this test to run in parallel with others

	// Define a helper struct to represent one test case
	type testCase struct {
		a, b float64 // inputs
		want float64 // expected output
	}

	// A slice of test cases to try
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	// Loop through the test cases
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b) // Call the function under test
		if tc.want != got {                // Compare expected vs actual
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
			// t.Errorf marks the test as failed but continues execution
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
