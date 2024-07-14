package person

import "fmt"

type Person struct {
	Name string
	Age  uint8 // 0 to 2^7-1
}

func (p *Person) Writer() string {
	return fmt.Sprintf("Person writes : ", p.Name, " with an age ", p.Age)
}
