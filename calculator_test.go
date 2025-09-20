package calculator_test

import (
	"testing"

	"github.com/brettfirecore/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel() // run test in parallel with other tests
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	var want float64 = 2
	got := calculator.Subtract(4, 2) // 4->(b) minus from 2->(a) is -2 throws an error
									 // refer calculator.go line 13
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	var want float64 = 9
	got := calculator.Multiply(3, 3) // 3 times 3 equal 9
	
	if got != want {
		t.Fatalf("got %f want %f", got, want)
	}
}
