package test

import (
	"gobyexample/pkg/slices"
	"testing"
)

func compareSlices(a []int32, b []int32) bool {

	if len(a) == len(b) {

		for i := range a {

			if a[i] != b[i] {
				return false
			}
		}

	}

	return true
}

func TestSlice(t *testing.T) {

	arr := [127]int32{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	have := slices.SliceMethod(arr, 8)
	expected := []int32{
		2, 4, 6, 8,
	}

	if compareSlices(have, expected) {
		t.Logf("Passed the slice test")
	} else {
		t.Fatalf("Mismatching slices")
	}
}
