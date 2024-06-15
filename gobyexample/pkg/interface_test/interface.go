package interface_test

import "reflect"

// MUST Read for conceptual understanding
// interface acts as a medium through which we define common functionalities of different struct type objects
// .. the instances of the struct need not explicitly implement the interface using a variable
// .. rather when the struct instance implements the public methods of the interface, it automatically gets intanciated
// .. mostly used as a middle man or common gateway for different inputs requesting same functionality

type ServerGateway interface {
	Authenticate() bool // just returns true or false
}

// let us say there are two kinds of server which takes the password
type IntegerServer struct {
	username string
	password int
}

// method signature for integer server
func (s *IntegerServer) Authenticate() bool {
	return reflect.TypeOf(s.password).Kind() == reflect.Int
}

type StringServer struct {
	username string
	password string
}

// method signature for integer server
func (s *StringServer) Authenticate() bool {
	return reflect.TypeOf(s.password).Kind() == reflect.String
}
