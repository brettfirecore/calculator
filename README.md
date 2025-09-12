# Calculator

[![CI](https://github.com/brettfirecore/calculator/actions/workflows/test.yml/badge.svg)](https://github.com/brettfirecore/calculator/actions/workflows/test.yml)
[![Release](https://github.com/brettfirecore/calculator/actions/workflows/release.yml/badge.svg)](https://github.com/brettfirecore/calculator/actions/workflows/release.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/brettfirecore/calculator/calculator.svg)](https://pkg.go.dev/github.com/brettfirecore/calculator/calculator)

A simple Go module providing basic arithmetic operations with error handling,  
plus a CLI tool that wraps the library.

---

## Features

### Library (`calculator` package)
- `AddMany(inputs ...float64) float64`
- `SubtractMany(inputs ...float64) float64`
- `MultiplyMany(inputs ...float64) float64`
- `DivideMany(inputs ...float64) (float64, error)`

### CLI (`cmd/calculator`)
- `-op add|sub|mul|div` to pick the operation
- Accepts any number of operands (`calculator -op add 1 2 3 4`)
- `-version` flag shows build version and date

---

## Installation

### As a library
```bash
go get github.com/brettfirecore/calculator/calculator
