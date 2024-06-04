package anonymous

type GrpDatatype struct {
	s string
	i int32
}

// best use of anonymous function in golang
func StringModifiability() (bool, error) {

	dataset := GrpDatatype{
		"hello",
		45,
	}

	var addon byte = 'g'

	// anonymous function to display string modifiability
	modifiabilityCheck := func(obj GrpDatatype) bool {
		stringbytes := []byte(obj.s)
		stringbytes[0] = addon
		outstring := string(stringbytes)

		// returns true if mutable
		return outstring == obj.s
	}(dataset)

	return modifiabilityCheck, nil
}
