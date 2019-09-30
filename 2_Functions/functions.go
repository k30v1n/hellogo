package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string
	var age string

	name, age = format("John", 40)

	fmt.Println(name + ", " + age)
}

// Receives two parameters and returns two parameters
func format(name string, age int64) (string, string) {
	return "Name : " + name, "Age : " + strconv.FormatInt(age, 10)
}
