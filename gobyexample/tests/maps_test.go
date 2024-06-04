package test

import (
	"gobyexample/pkg/maps"
	"testing"
)

func comparemaps(a map[string]string, b map[string]string) bool {
	// Check if maps have the same length. If not, they can't be equal.
	if len(a) != len(b) {
		return false
	}

	// Loop through each key in map 'a'
	for key, valueA := range a {
		// Check if key exists in map 'b'
		valueB, ok := b[key]
		if !ok {
			return false // Key not found in 'b'
		}

		// Check if values for the same key are equal
		if valueA != valueB {
			return false // Values don't match
		}
	}

	// If all keys and values match, return true
	return true
}

func TestMaps(t *testing.T) {
	have := maps.MapsMethod("./maps_test_data.csv")

	expectedKey := maps.KeyStoreDT{
		Key: map[string]string{
			"name-age":     "Alice-30",
			"name-city":    "Alice-New York",
			"name-country": "Alice-USA",
		},
	}

	if comparemaps(have[0].Key, expectedKey.Key) {
		t.Logf("Test passed successfully")
	} else {
		t.Fatalf("KeyStores didn't match")
	}

}
