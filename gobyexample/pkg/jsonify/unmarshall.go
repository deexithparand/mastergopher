package jsonify

import (
	"encoding/json"
)

// marshall and unmarshall json
// validate and print a json

// why empty interface because any value can be the unmarshalled output of a json string
type JsonOut map[string]interface{}

// unmarshall this json string
func MyUnmarshall(jsonstring string) (JsonOut, error) {

	inputByteArray := []byte(jsonstring)

	var JsonMap JsonOut
	err := json.Unmarshal(inputByteArray, &JsonMap)
	if err != nil {
		return JsonMap, err
	}

	return JsonMap, nil
}
