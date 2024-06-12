package utils

import "fmt"

// concept type polymorphism
// empty interface that accepts anything
type writerinput interface{}

// function to avoid writing console statements everywhere
func Writer(wi writerinput) {
	fmt.Println(wi)
}
