package main

import "t2/pkg/employee"

func main() {
	var jman employee.Employee

	jman.Name = "deexithparand"
	jman.Age = 22
	jman.EmployeeID = "JMD275"

	jman.Writer()
}
