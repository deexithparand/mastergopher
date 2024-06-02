package test

import (
	"gobyexample/pkg/values"
	r "reflect"
	"testing"
)

func evalute(dataset values.GrpDatatype) bool {
	return r.TypeOf(dataset.S).Kind() == r.String && r.TypeOf(dataset.I).Kind() == r.Int32 && r.TypeOf(dataset.F).Kind() == r.Float32
}

func TestValues(t *testing.T) {
	have := values.Values_method()
	want := evalute(have)
	// check the type of each var in the struct type

	if !want {
		t.Fatalf("Mismatched group data type")
	}
}
