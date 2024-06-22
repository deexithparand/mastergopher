package jsonify // works with slice of bytes []byte and errors only
import (
	"encoding/json"
	"fmt"
)

// input : json string (refer gobyexample for json standards and rules) -> variable
func Unmarshall(s string) (any, error) {

	var v interface{}

	// give var of any type
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		return v, fmt.Errorf("unmarshalling error : %v", err)
	}

	return v, err
}

// input : variable -> jsonstring
func Marshall(v any) (string, error) {

	var (
		out []byte
		err error
	)

	out, err = json.Marshal(v)
	if err != nil {
		return string(out), fmt.Errorf("marshalling error : %v", err)
	}

	return string(out), nil
}
