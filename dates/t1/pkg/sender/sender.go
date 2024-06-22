package sender

import "fmt"

type Salaried struct {
	Name              string
	Balance           int64
	AllowedCurrencies []string
	EmpId             string
}

func (s *Salaried) Validate() bool {

	fmt.Println("Vera maari")

	if len(s.Name) > 3 {
		return true
	}

	return false
}

type Entreprenuer struct {
	Name              string
	Balance           int64
	AllowedCurrencies []string
	BusinessId        string
}
