package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"io"
	"github.com/brettfirecore/calculator/calculator"
)

var (
	version   = "dev"
	buildDate = ""
)

func run(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("calculator", flag.ContinueOnError)
	showVersion := fs.Bool("version", false, "print version and exit")
	op := fs.String("op", "", "operation: add|sub|mul|div")
	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "usage: calculator -op <add|sub|mul|div> <numbers...>")
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *showVersion {
		fmt.Printf("calculator version=%s buildDate=%s\n", version, buildDate)
		return nil
	}
	if *op == "" {
		return fmt.Errorf("missing -op (add|sub|mul|div)")
	}
	if fs.NArg() == 0 {
		return fmt.Errorf("provide at least one number")
	}

	// parse numbers
	vals := make([]float64, 0, fs.NArg())
	for _, s := range fs.Args() {
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return fmt.Errorf("invalid number %q: %v", s, err)
		}
		vals = append(vals, n)
	}

	switch *op {
	case "add":
    		fmt.Fprintln(out, calculator.AddMany(vals...))
	case "sub":
    		fmt.Fprintln(out, calculator.SubtractMany(vals...))
	case "mul":
    		fmt.Fprintln(out, calculator.MultiplyMany(vals...))
	case "div":
    		res, err := calculator.DivideMany(vals...)
    		if err != nil {
        		return err
    		}
    		fmt.Fprintln(out, res)

	default:
		return fmt.Errorf("unknown op %q (use add|sub|mul|div)", *op)
	}
	return nil
}

func main() {
    if err := run(os.Args[1:], os.Stdout); err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
    os.Exit(0)
}

