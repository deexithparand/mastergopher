package maps

import (
	"encoding/csv"
	"log"
	"os"
)

// Create a map to store the names and ages of several people. Write a function to add a new person to the map and another function to retrieve a person's age by their name.
type CityData struct {
	Name    string
	Age     string
	City    string
	Country string
}

type KeyStoreDT struct {
	Key map[string]string // more than one key-value pairs
}

func CreateNewKeyStore(name string, age string, city string, country string) KeyStoreDT {
	var keystore KeyStoreDT

	key := map[string]string{
		"name-age":     name + "-" + age,
		"name-city":    name + "-" + city,
		"name-country": name + "-" + country,
	}

	keystore = KeyStoreDT{
		Key: key,
	}

	return keystore
}

func MapsMethod(path string) []KeyStoreDT {

	var response []KeyStoreDT

	// read from a csv file
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Error opening file from csv")
	}

	defer file.Close()

	fileReader := csv.NewReader(file)

	records, err := fileReader.ReadAll()

	if err != nil {
		log.Fatalf("Error opening csv file")
	}

	for index, eachrecord := range records {
		if index > 0 {
			response = append(response, CreateNewKeyStore(eachrecord[0], eachrecord[1], eachrecord[2], eachrecord[3]))
		}
	}

	return response
}
