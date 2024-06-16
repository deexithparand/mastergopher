package interfacesnew

import "errors"

type Validator interface {
	Validate() error
}

// passing validate function by value
// Passing an interface by value still effectively passes a reference to the underlying data. (V.V.Important)
func ValidateGatewayMethod(v Validator) error {
	// inherently pass by value works while passing by value to interface
	return v.Validate()
}

type User struct {
	Username string
	Pswd     string
}

// just for test validate will help
func (u *User) Validate() error {
	if u.Username == "" && u.Pswd == "" {
		return errors.New("invalid username and password")
	} else if u.Username == "" {
		return errors.New("invalid username")
	} else if u.Pswd == "" {
		return errors.New("invalid password")
	} else {
		return nil
	}
}
