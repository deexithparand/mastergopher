package test

import (
	"gobyexample/pkg/control"
	"testing"
)

func TestControl(t *testing.T) {
	response := control.Control_method([]uint32{1, 2, 3, 4, 5}, 'a')
	expected := control.ResponseDatatype{
		Compute:     15,
		SliceLength: "too small",
		VowelCheck:  "lowercase vowel",
	}

	if response != expected {
		t.Fatalf("FAIL : Mismatching Response")
	}
}
