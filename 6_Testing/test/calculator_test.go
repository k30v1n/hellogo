package test

import (
	"testing"

	"../calculator"
)

// There are few specifications for adding a test file in Go. First, file name should end
// with *_test.go and second, test function name should start with Test*. The test function
// takes an object of type *testing.T, which comes from the testing package.

// run tests on this folder
// go test
// run tests on sub directories
// go test ./...

func TestAdd(t *testing.T) {
	sum := calculator.Add(1, 2)
	if sum != 3 {
		t.Error("Sum Did Not Match!!")
		t.Fail()
		return
	}
	t.Log("Add Test Passed")
}
