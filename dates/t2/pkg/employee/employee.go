package employee

import (
	"encoding/json"
	"fmt"
	"log"
	"t2/pkg/person"
)

type Employee struct {
	EmployeeID string
	person.Person
}

func GetEmployeeDetails(e *Employee) string {

	// map and return as json string
	var employeeMap map[string]string = make(map[string]string)
	var errId int32
	var jsonString string

	employeeMap["Name of the Employee"] = e.Name
	employeeMap["Age of the Employee"] = fmt.Sprintf("%d", e.Age)
	employeeMap["Employee ID of Employee"] = e.EmployeeID

	ok, err := json.Marshal(employeeMap)
	if err != nil {
		log.Fatalf("Error marshalling map to json : %d", errId)
	}

	fmt.Println("Output json string : ", string(ok))

	jsonString = string(ok)

	return jsonString
}

func (e *Employee) Writer() string {
	return GetEmployeeDetails(e)
}
