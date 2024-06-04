package test

import (
	"fmt"
	"gobyexample/pkg/array"
	"testing"
)

func TestArray(t *testing.T) {

	arrOneD := [5]int32{1, 2, 3, 4, 5}
	arrTwoD := [5][5]int32{
		{
			1, 1, 1,
		}, {
			1, 1, 1,
		}, {
			1, 1, 1,
		},
	}

	haveOneD := array.Array_Method_OneD(arrOneD) // expects only array of size 5
	expectedOneD := array.OneDResponse{
		Compute: 15,
		Size:    5,
	}

	fmt.Println("anmma sonnange", haveOneD)

	haveTwoD := array.Array_Method_TwoD(arrTwoD)
	expectedTwoD := array.TwoDResponse{
		DiagonalCompute: 3,
		Compute:         9,
	}

	if haveOneD != expectedOneD {
		t.Fatalf("Unmatching One D Response")
	} else if haveTwoD != expectedTwoD {
		t.Fatalf("Unmatching Two D Response")
	} else {
		t.Logf("Test for arrays passed successfully!")
	}

}
