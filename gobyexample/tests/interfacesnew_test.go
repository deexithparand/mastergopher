package test

import (
	"fmt"
	"gobyexample/pkg/interfacesnew"
	"testing"
)

func TestInterfacesNew(t *testing.T) {

	// define two struct object that implement the interface
	// .. tryout : passing the struct object instance to the method signature of the interface directly
	var userOne *interfacesnew.User = &interfacesnew.User{
		Username: "joe",
		Pswd:     "ds21!#!#$$@!#@!d",
	}

	fmt.Printf("true : %s", userOne)

	err := interfacesnew.ValidateGatewayMethod(userOne)
	if err != nil {
		t.Fatalf("Error : %s", err)
	}

	t.Logf("Test Passed validation successfull")
}
