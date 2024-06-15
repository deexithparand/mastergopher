package test

import (
	"gobyexample/pkg/interfaces"
	"testing"
)

func TestInterfaces(t *testing.T) {
	have := true

	is := interfaces.IntegerServer{
		Username: "harsh",
		Password: 34,
	}

	expected := is.Authenticate()

	if have != expected {
		t.Fatalf("Integer Server failed test")
		// func exits here
	}

	t.Logf("Integer Server Test Passed Successfully")
}
