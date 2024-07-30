package main

import (
	"fmt"
)

func arrays() {
	fmt.Println("Example of Array in Go")
	arrayType1()
	arratTyepe2()
	loopingArray()

}

func arrayType1() {
	var numbers [1]int
	numbers[0] = 1
	fmt.Println("Array type 1", numbers)
}

func arratTyepe2() {
	/* Define and initialize an array in one step */
	arrayVariable := [2]int{1, 420}
	fmt.Println("Array type 2", arrayVariable)
}

func loopingArray() {
	arrayVariable := []int{1, 2, 3}

	/*

		The range returns index & value

	*/
	for index, value := range arrayVariable {
		fmt.Println("Index", index, " Value is -->", value)
	}
}
