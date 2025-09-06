package calculator_test

import (
	"math"
	"testing"
	"github.com/brettfirecore/calculator/calculator"
)

func closeEnough(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestAdd(t *testing.T) {
	got := Add(2, 2) // no prefix needed
	want := 4.0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(5, 3)
	want := 2.0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(3, 4)
	want := 12.0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatal(err)
	}
	want := 5.0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Fatal("expected error for division by zero, got nil")
	}
}

func TestSqrt(t *testing.T) {
	got, err := Sqrt(16)
	if err != nil {
		t.Fatal(err)
	}
	want := 4.0
	if !closeEnough(got, want, 1e-9) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSqrtNegative(t *testing.T) {
	_, err := Sqrt(-1)
	if err == nil {
		t.Fatal("expected error for negative input, got nil")
	}
}
