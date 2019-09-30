package main

import (
	"fmt"
)

const (
	CONSTANT = "This is a constant variable"
)

// Multiple struct
type (
	sampleStruct struct {
		Name string
	}

	sampleStruct2 struct {
		Name string
	}
)

func main() {

	// inline struct
	type sampleStruct3 struct {
		// Public to "main" package
		Name string
		// Private to "main" package
		lastName string
	}

	var variable string = "I can define types"
	var variable2 = " I can infer types"
	// variable2 = 2 //error

	fmt.Println("Hello Golang!! " + variable + variable2)
}
