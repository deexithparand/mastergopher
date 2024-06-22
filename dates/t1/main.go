package main

import (
	"fmt"
	"t1/pkg/sender"
	"t1/pkg/validate"
)

func main() {
	// to bring back the usage of type assertion and interface usage also jsonseal package demo too
	// .. I can define a transaction method
	// .. using type assertion I can find the endpoint from where I get the payment
	// .. -> pay from source 1 and source 2 are of different types -> validated with jsonseal -> source found and printed
	// validate - jsonseal, transaction will choose the source from the two, assertion will validate the source too

	// create two senders
	s1 := sender.Salaried{
		Name:    "gopi",
		Balance: 100,
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

	var payment1 validate.Validator = &s1
	fmt.Println(validate.ValidatePayment(payment1))

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
