package main

import (
	"fmt"
	"t1/pkg/sender"
	"t1/pkg/validate"
)

func main() {

	// create two senders
	s1 := sender.Salaried{
		Name:    "gopi",
		Balance: 10000,
		AllowedCurrencies: []string{
			"inr", "yen",
		},
		EmpId: "JMD275",
	}

	// s2 := sender.Entreprenuer{
	// 	Name:    "nath",
	// 	Balance: 10000,
	// 	AllowedCurrencies: []string{
	// 		"usd", "inr", "yen", "eur",
	// 	},
	// 	BusinessId: "JMD001",
	// }

	// employee transactions
	var empPayment validate.Validator = &s1
	// var entrePayment validate.Validator = &s2

	validArr, invalidArr, validateErr := validate.ValidatePayment(empPayment)
	if validateErr != nil {
		fmt.Println("Error : ", validateErr.Error())
	} else {
		fmt.Println("Valid Payments List : ", validArr)
		fmt.Println("InValid Payments List : ", invalidArr)
	}

	// jsoninbuilt package tryout
	// jsonstring, err := os.ReadFile("./data/validpay.json")
	// fmt.Println("File contents : ", string(jsonstring))
	// if err != nil {
	// 	fmt.Println("file opening error : ", err)
	// 	return
	// }

	// outvar, err := jsonify.Unmarshall(string(jsonstring))
	// if err != nil {
	// 	fmt.Println("fmt out error : ", err)
	// 	return
	// }

	// fmt.Println("fmt out of unmarshall : ", outvar)
}
