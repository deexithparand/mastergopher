package constants

func GetTypesFromTypeMap(label string) string {

	typemap := map[string]string{
		"s":  "student",
		"t":  "teacher",
		"et": "external staff",
		"o":  "others",
		"n":  "nil",
	}

	return typemap[label]
}
