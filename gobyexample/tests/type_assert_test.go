package test

import (
	tyassert "gobyexample/pkg/type_assertion"
	"testing"
)

func TestTypeAssert(t *testing.T) {
	istore := tyassert.Intstore{
		Intarr: []int{1, 2, 3, 4, 5},
	}
	sstore := tyassert.Stringstore{
		Stringarr: []string{"abc", "abc", "abc"},
	}

	ierr := tyassert.AppendMethod(&istore, 5)
	serr := tyassert.AppendMethod(&sstore, "dsds")
	if ierr == nil && serr == nil {
		t.Logf("Test passed, int arr : %v and string arr : %v\n", istore.Intarr, sstore.Stringarr)
	} else {
		t.Fatalf("pochu")
	}
}
