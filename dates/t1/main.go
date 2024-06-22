package main

func main() {
	// to bring back the usage of type assertion and interface usage also jsonseal package demo too
	// .. I can define a transaction method
	// .. using type assertion I can find the endpoint from where I get the payment
	// .. -> pay from source 1 and source 2 are of different types -> validated with jsonseal -> source found and printed
	// validate - jsonseal, transaction will choose the source from the two, assertion will validate the source too

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
