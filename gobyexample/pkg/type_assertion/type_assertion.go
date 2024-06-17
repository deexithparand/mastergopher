package tyassert

import "errors"

// purpose of type assertion is to check if the variable of the interface belongs to a specific type
// .. lets say a struct variable of a type implements a struct
// .. then the interface variable to which the type is embedded is checked for its type
// .. type assertion of the interface variable returns true =>

// syntax : interface_variable.(type_to_check) => valid and returns value => type of variable to which assertion is done is true
// .. if type error => panic throws error
type Appender interface {
	Append(any) error
}

// interface method
func AppendMethod(a Appender, v any) error {
	return a.Append(v) // just calling the func
}

// type assertion for slice of int and string seperately
type Intstore struct {
	Intarr []int
}

func (i *Intstore) Append(v any) error {
	// very good usecase of type assertion
	// .. asserts the interface variable v -> to be of a specific type
	if v, ok := v.(int); ok {
		i.Intarr = append(i.Intarr, v)
	} else {
		return errors.New("Invalid Type of input to be appended")
	}
	return nil
}

type Stringstore struct {
	Stringarr []string
}

func (s *Stringstore) Append(v any) error {
	// very good usecase of type assertion
	// .. asserts the interface variable v -> to be of a specific type
	if v, ok := v.(string); ok {
		s.Stringarr = append(s.Stringarr, v)
	} else {
		return errors.New("Invalid input type to be appended")
	}
	return nil
}
