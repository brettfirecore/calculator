package main

import (
	"bytes"
	"testing"
)

func TestRun_Add(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer

	if err := run([]string{"-op", "add", "2", "3", "4"}, &buf); err != nil {
		t.Fatal(err)
	}

	got := buf.String()
	want := "9\n"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
