package test

import (
	"gobyexample/pkg/jsonify"
	"testing"
)

// setup a json variable
var JsonString string = `{
	"company":"jp morgan & chase",
	"ceo":"jamie dimon"
}`

func match(a map[string]any, b map[string]any) bool {
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}

func TestJsonify(t *testing.T) {
	have := map[string]any{
		"company": "jp morgan & chase",
		"ceo":     "jamie dimon",
		// "secret sauce": []any{
		// 	"idk",
		// 	"yes",
		// 	"ikr",
		// },
	}

	expected, err := jsonify.MyUnmarshall(JsonString)
	if err != nil {
		t.Fatalf("error while my unmarshall %s", err)
	}

	if match(have, expected) {
		t.Logf("unmarshalling tests passed successfully")
	} else {
		t.Fatalf("Unmatched or unexpected map after un marshalling")
	}

	// try using struct
	// if reflect.DeepEqual(have, expected) {
	// 	t.Logf("unmarshalling tests passed successfully")
	// } else {
	// 	fmt.Println(have, "ddd", expected)
	// 	t.Fatalf("Unmatched or unexpected map after un marshalling")
	// }
}
