package validate

// validate must be an interface because it must validate the sender and the payment

// GENERAL VALIDATION JSON
// ... currency must not be below 0

// BASED ON OUR CUSTOM VALIDATION

// ... First level : allowed only to send mony in allowed currencies
// ... optional Second Level : allowed to send money only less than balance amount
type Validator interface {
	Validate() ([]int, []string, error) // returns (valid transaction indices, alert strings for invalidity)
}

func ValidatePayment(v Validator) ([]int, []string, error) {
	return v.Validate()
}

// custom validations based on the sender will be implemented in the struct level of the sender
