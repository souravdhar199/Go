package main

import (
	"fmt"
	"time"
	"unicode"
)

func arrays() {
	fmt.Println("Example of Array in Go")
	arrayType1()
	arratTyepe2()
	loopingArray()
	arrayLooping()
	charcterArray()
	checkPalindrome()
	//funCombination()
	timeCheck()

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

func charcterArray() {
	//this is how we initialize an char array
	charArray := []rune{'a', 'b', 'a'}

	p := 0
	p2 := len(charArray) - 1

	for p <= p2 {
		fmt.Println("Hello", rune(charArray[p]))
		p++
	}
}

func checkPalindrome() {
	/*
		How to get a character from a string

		rune means char

		string1[1] ==> returns a string of b

	*/
	string1 := "aba"

	fmt.Println("This is one of the character from string1 --> ", string(string1[1]) == "as")

	fmt.Println(unicode.IsLetter(rune(string1[1])))

	fmt.Println(unicode.IsDigit(rune(string1[1])))

}

func funCombination() {
	chars := []string{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ", // X
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ", // H
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ", // U
		"0123456789",                 // 3
		"-",                          // -
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ", // D
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ", // X
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ", // B
		"0123456789",                 // 3
	}

	var results []string
	generateCombinations("", chars, len(chars), 0, &results)

	// Print the results

}

func generateCombinations(prefix string, chars []string, length int, pos int, results *[]string) {
	if pos == length {
		*results = append(*results, prefix)
		return
	}
	for _, char := range chars[pos] {
		fmt.Println(prefix + string(char))
		generateCombinations(prefix+string(char), chars, length, pos+1, results)
	}
}

func timeCheck() {
	numStrings := 1000000 // Number of strings to print for testing
	start := time.Now()

	for i := 0; i < numStrings; i++ {
		fmt.Println("This is string number", i)
	}

	duration := time.Since(start)
	fmt.Printf("Time taken to print %d strings: %v\n", numStrings, duration)

	// Extrapolate for 30 billion strings
	estimatedTime := duration * 30000
	fmt.Printf("Estimated time to print 30 billion strings: %v\n", estimatedTime)
}
