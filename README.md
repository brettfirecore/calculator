# calculator

[![Go CI](https://github.com/brettfirecore/calculator/actions/workflows/ci.yml/badge.svg)](https://github.com/brettfirecore/calculator/actions/workflows/ci.yml)

A simple Go module providing basic arithmetic operations with error handling.

## Features

- `Add(a, b float64) float64`
- `Subtract(a, b float64) float64`
- `Multiply(a, b float64) float64`
- `Divide(a, b float64) (float64, error)`
- `Sqrt(x float64) (float64, error)`

## Installation

```bash
go get github.com/brettfirecore/calculator/calculator

./calculator -op add 1 2 3    # 6
./calculator -op div 10 2     # 5
./calculator -op div 10 0     # Error + exit 1
