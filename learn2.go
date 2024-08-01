package main

import (
	"fmt"
)

func arrays() {
	fmt.Println("Example of Array in Go")
	arrayType1()
	arratTyepe2()
	loopingArray()
	arrayLooping()

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

func arrayLooping() int {

	//Example of initializing an array
	money := make([]int, 4)

	sum := 0

	for i := 0; i < len(money); i++ {
		money[i] = i * 2
	}
	// _ means ignore the return value
	for _, value := range money {
		sum += value
	}

	fmt.Println("Sum -->", sum)
	return sum
}
