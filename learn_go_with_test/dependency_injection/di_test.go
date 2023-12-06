package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	want := "Hello, Chris"
	got := buffer.String()

	if got != want {
		t.Errorf("want %v but got %v", want, got)
	}
}
