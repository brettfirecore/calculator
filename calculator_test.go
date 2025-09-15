package calculator_test

import (
	"testing"

	"github.com/brettfirecore/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	got := calculator.Add(2, 3) // this is float64
	want := 5.0                 // must be float64 too

	if got != want {
		t.Fatalf("got %f want %f", got, want)
	}
}
