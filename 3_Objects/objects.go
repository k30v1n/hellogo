package main

import (
	"fmt"
	"strconv"
)

// Person is a simple person obj
type Person struct {
	ID    int
	Name  string
	Age   int
	Email string
}

func main() {
	// Object with no arguments
	//person1 := Person{}

	// Object with all args
	//person2 := Person{ID: 1, Age: 30, Email: "email", Name: "Josh"}

	// Object with partial args
	//person3 := Person{Name: "Josh"}

	//------------------------------------------------
	// Arrays vs Slices: Arrays in Go are fixed length collection of objects of same type,
	// whereas slices are variable length collections. In general, slices are much more
	// powerful than Arrays and are used much more extensively.

	//var arr1 = [5]int{}          // Passing length
	//var arr2 = [...]int{1, 2, 3} // Letting the compiler define length
	//var lenght2 = len(arr2)      // Acessing array`s length

	// Slices in Go, are variable length collections of same data type. Go Slice is a powerful
	// data structure, built over Arrays, keeping ease of use and scalability in mind.
	// Unlike arrays, slices are not of fixed length. While declaring a Slice,
	// you don’t need to provide size.
	// var sl1 = []int{}        // Slice of int
	// var sl2 = []int{1, 2, 3} // Slice with a initialization

	// You can also make a slice with length and capacity options, with an inbuilt function make.
	// var sl3 = make([]int, 3, 5) // 3 is length, 5 is capacity
	// var sl4 = make([]int, 5)    // if capacity is ommited, it uses length to define defaul capacity

	// Access length and capacity of a slice
	// length := len(sl1)
	// capacity := cap(sl1)

	var slice = []int{}
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	// ommiting index variable
	fmt.Println("Printing all slice: ")
	for _, val := range slice {
		fmt.Println(val)
	}
	// for i := 0; i < len(slice); i++ {
	// 	...
	// }

	// MAPS is a key value structure
	// map1 := make(map[int]string)
	// map2 := map[int]string{}
	map3 := map[int64]string{1: "A", 2: "B"}
	//Checking if a key exists there
	if _, ok := map3[2]; ok {
		fmt.Println("Printing all map values: ")
		fmt.Println(len(map3))

		for key, value := range map3 {
			fmt.Println(strconv.FormatInt(key, 10) + ": " + value)
		}

		delete(map3, 2) // deleting key 2 value
		fmt.Print("Map size: ")
		fmt.Println(len(map3))

		// Common for
		// for i := 0; i < 10; i++ {
		// 	val += i
		// }
		// Common while
		// for i < 10 {
		// 	val += i
		// 	i++
		// }
		// Inifinitive loop
		// for {
		// 	val += i
		// }
	}
}
