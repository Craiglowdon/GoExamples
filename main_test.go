package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	expected := "Hello World!"
	if ret := HelloWorld(); ret != expected {
		t.Errorf("main() = %q, want %q", ret, expected)
	}
}
