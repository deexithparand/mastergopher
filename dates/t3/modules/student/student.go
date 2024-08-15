package student

import (
	"t3/constants"
)

type Student struct {
	Name string
	Age  int
	Type string
}

func (s *Student) Create(name string, age int) *Student {
	s.Name = name
	s.Age = age
	s.Type = constants.GetTypesFromTypeMap("s")

	return s
}
