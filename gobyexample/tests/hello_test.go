package hello_testing

import (
	"errors"
	"gobyexample/pkg/hello"
	"testing"
)

// string comparison implementation
func strcmp(a, b string) error {
	for i, _ := range a {
		if a[i] != b[i] {
			return errors.New("to break the scope")
		}
	}

	for i, _ := range b {
		if a[i] != b[i] {
			return errors.New("to break the scope")
		}
	}

	return nil
}

func TestHelloWorld(t *testing.T) {
	got := hello.Hello_world_method()
	want := "hellco"
	err := strcmp(got, want)
	if err != nil {
		t.Fatalf("Strings do not match")
	}
}
