package calculator_test

import (
    "testing"

    "github.com/brettfirecore/calculator/calculator"
)

func TestAddMany(t *testing.T) {
    t.Parallel()
    type testCase struct {
        inputs []float64
        want   float64
    }
    cases := []testCase{
        {[]float64{}, 0},
        {[]float64{2}, 2},
        {[]float64{2, 2}, 4},
        {[]float64{1, 2, 3}, 6},
    }
    for _, tc := range cases {
        got := calculator.AddMany(tc.inputs...)
        if tc.want != got {
            t.Errorf("AddMany(%v): want %f, got %f", tc.inputs, tc.want, got)
        }
    }
}

func TestSubtractMany(t *testing.T) {
    t.Parallel()
    type testCase struct {
        inputs []float64
        want   float64
    }
    cases := []testCase{
        {[]float64{}, 0},
        {[]float64{2}, 2},
        {[]float64{2, 2}, 0},
        {[]float64{1, 2, 3}, -4},
    }
    for _, tc := range cases {
        got := calculator.SubtractMany(tc.inputs...)
        if tc.want != got {
            t.Errorf("SubtractMany(%v): want %f, got %f", tc.inputs, tc.want, got)
        }
    }
}

func TestMultiplyMany(t *testing.T) {
    t.Parallel()
    type testCase struct {
        inputs []float64
        want   float64
    }
    cases := []testCase{
        {[]float64{}, 0},
        {[]float64{2}, 2},
        {[]float64{2, 2}, 4},
        {[]float64{1, 2, 3, 4}, 24},
    }
    for _, tc := range cases {
        got := calculator.MultiplyMany(tc.inputs...)
        if tc.want != got {
            t.Errorf("MultiplyMany(%v): want %f, got %f", tc.inputs, tc.want, got)
        }
    }
}

func TestDivideMany(t *testing.T) {
    t.Parallel()
    type testCase struct {
        inputs []float64
        want   float64
    }
    cases := []testCase{
        {[]float64{}, 0},
        {[]float64{2}, 2},
        {[]float64{2, 2}, 1},
        {[]float64{-1, -1}, 1},
        {[]float64{100, 5, 2}, 10},
    }
    for _, tc := range cases {
        got, err := calculator.DivideMany(tc.inputs...)
        if err != nil {
            t.Fatalf("DivideMany(%v): want no error for valid input, got %v", tc.inputs, err)
        }
        if tc.want != got {
            t.Errorf("DivideMany(%v): want %f, got %f", tc.inputs, tc.want, got)
        }
    }
}

func TestDivideManyInvalid(t *testing.T) {
    t.Parallel()
    _, err := calculator.DivideMany(10, 5, 0, 7)
    if err == nil {
        t.Error("DivideMany(10, 5, 0, 7): want error for zero divisor, got nil")
    }
}
