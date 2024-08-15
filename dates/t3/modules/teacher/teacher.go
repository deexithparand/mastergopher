package teacher

import (
	"t3/constants"
)

type Teacher struct {
	Name string
	Age  int
	Type string
}

func (t *Teacher) Create(name string, age int) *Teacher {
	t.Name = name
	t.Age = age
	t.Type = constants.GetTypesFromTypeMap("t")

	return t
}
