package calculator_test

import (
	"testing"

	"github.com/brettfirecore/calculator/calculator"
)

func TestAddMany(t *testing.T) {
	t.Parallel()
	if got, want := calculator.AddMany(), 0.0; got != want {
		t.Errorf("AddMany() = %v, want %v", got, want)
	}
	if got, want := calculator.AddMany(1, 2, 3.5), 6.5; got != want {
		t.Errorf("AddMany(1,2,3.5) = %v, want %v", got, want)
	}
}

func TestMultiplyMany(t *testing.T) {
	t.Parallel()
	if got, want := calculator.MultiplyMany(), 1.0; got != want {
		t.Errorf("MultiplyMany() = %v, want %v", got, want)
	}
	if got, want := calculator.MultiplyMany(2, 3, 4), 24.0; got != want {
		t.Errorf("MultiplyMany(2,3,4) = %v, want %v", got, want)
	}
}

func TestSubtractMany(t *testing.T) {
	t.Parallel()
	if got, want := calculator.SubtractMany(), 0.0; got != want {
		t.Errorf("SubtractMany() = %v, want %v", got, want)
	}
	if got, want := calculator.SubtractMany(10), 10.0; got != want {
		t.Errorf("SubtractMany(10) = %v, want %v", got, want)
	}
	if got, want := calculator.SubtractMany(10, 1, 2, 3), 4.0; got != want {
		t.Errorf("SubtractMany(10,1,2,3) = %v, want %v", got, want)
	}
}

func TestDivideMany(t *testing.T) {
	t.Parallel()
	if _, err := calculator.DivideMany(); err == nil {
		t.Fatalf("DivideMany() want error, got nil")
	}
	if got, err := calculator.DivideMany(12); err != nil || got != 12 {
		t.Fatalf("DivideMany(12) got (%v,%v), want (12,nil)", got, err)
	}
	if got, err := calculator.DivideMany(12, 4, 3); err != nil || got != 1 {
		t.Fatalf("DivideMany(12,4,3) got (%v,%v), want (1,nil)", got, err)
	}
	if _, err := calculator.DivideMany(1, 0); err == nil {
		t.Fatalf("DivideMany(1,0) want divide-by-zero error, got nil")
	}
}
