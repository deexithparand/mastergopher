package test

import (
	"errors"
	"gobyexample/pkg/hello"
	"testing"
)

// string comparison implementation
func strcmp(a, b string) error {

	if len(a) != len(b) {
		return errors.New("Different length strings")
	}

	for i := range a {
		if a[i] != b[i] {
			return errors.New("Mismatching strings")
		}
	}

	return nil
}

func TestHello(t *testing.T) {
	have := hello.Hello_method()
	want := "hello"

	err := strcmp(have, want)
	if err != nil {
		t.Fatalf("Strings do not match")
	}
}
