package control

type ResponseDatatype struct {
	Compute     uint32
	SliceLength string
	VowelCheck  string
}

func ForPositiveCompute(arr []uint32) uint32 {
	var compute uint32 = 0
	for _, v := range arr {
		compute += v
	}
	return compute
}

func CheckLengthSlice(arr []uint32) string {
	if len(arr) < 10 {
		return "too small"
	} else if len(arr) >= 10 {
		return "ideal"
	} else {
		return "too big"
	}
}

func CheckVowels(char byte) string {
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		return "lowercase vowel"
	case 'A', 'E', 'I', 'O', 'U':
		return "uppercase vowel"
	default:
		return "consonant"
	}
}

func Control_method(slice []uint32, char byte) ResponseDatatype {

	// for loop , sum of elements in array
	a := ForPositiveCompute(slice)

	// if else multiple cases
	b := CheckLengthSlice(slice)

	// switch
	c := CheckVowels(char)

	response := ResponseDatatype{
		Compute:     a,
		SliceLength: b,
		VowelCheck:  c,
	}

	return response
}
