package test

import (
	"gobyexample/pkg/anonymous"
	"testing"
)

func TestAnonymous(t *testing.T) {
	_, err := anonymous.StringModifiability()
	if err != nil {
		t.Fatalf("Error in package %v", err)
	}
}
