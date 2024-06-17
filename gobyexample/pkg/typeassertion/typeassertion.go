package tassert

// purpose of type assertion is to check if the variable of the interface belongs to a specific type
// .. lets say a struct variable of a type implements a struct
// .. then the interface variable to which the type is embedded is checked for its type
// .. type assertion of the interface variable returns true =>

// syntax : interface_variable.(type_to_check) => valid and returns value => type of variable to which assertion is done is true
// .. if type error => panic throws error
