package sender

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"t1/pkg/payment"

	"github.com/scriptnull/jsonseal"
)

type Salaried struct {
	Name              string
	Balance           int64
	AllowedCurrencies []string
	EmpId             string
}

func checkAllowedCurrencies(s *Salaried, cur string) bool {
	fmt.Println("Entered Currencies : ", cur)
	for _, v := range s.AllowedCurrencies {
		if v == cur {
			return true
		}
	}

	return false
}

func ValidateEmpTransaction(s *Salaried, v payment.TransactionDetails) error {
	var salariedCheckGroup jsonseal.CheckGroup

	salariedCheckGroup.Check(func() error {

		if s.Balance < int64(v.Amount) {
			return errors.New("low balance in bank")
		}

		if !checkAllowedCurrencies(s, v.Currency) {
			return errors.New("employees allowed only these currencies")
		}

		return nil
	})

	return salariedCheckGroup.Validate()
}

func (s *Salaried) Validate() ([]int, []string, error) {

	var (
		validIndices   []int
		invalidStrings []string
	)

	fileout, err := os.ReadFile("C:\\Users\\deexi\\OneDrive\\Desktop\\Deexi\\GoExtended\\gophercises\\dates\\t1\\data\\emp.json")
	if err != nil {
		return validIndices, invalidStrings, fmt.Errorf("file error : %v", err)
	}

	var empPayments payment.PaymentStruct

	err = json.Unmarshal(fileout, &empPayments)
	if err != nil {
		return validIndices, invalidStrings, fmt.Errorf("unmarshalling error : %v", err)
	}

	for i, v := range empPayments.Transactions {
		transactionErr := ValidateEmpTransaction(s, v)
		if transactionErr != nil {
			invalidStrings = append(invalidStrings, transactionErr.Error())
		} else {
			validIndices = append(validIndices, i)
		}
	}

	return validIndices, invalidStrings, nil
}

type Entreprenuer struct {
	Name              string
	Balance           int64
	AllowedCurrencies []string
	BusinessId        string
}
