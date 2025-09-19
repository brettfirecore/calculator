package calculator_test

import (
	"testing"

	"github.com/brettfirecore/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel() // run test in parallel with other tests

	got := calculator.Add(2, 3) // this is float64
	want := 5.0                 // must be float64 too

	if got != want { // if got is not equal to want then run t.Fatalf error
		t.Fatalf("got %f want %f", got, want)
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
