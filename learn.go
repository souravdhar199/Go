package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	test1(19, "WWW")
	fmt.Println(fucnThatreturnint("String 1", "String 2"))

	value1, value2 := swap("Hello", "name")

	fmt.Println(value2, value1)

	// This function from another file
	arrays()
	learnMoreArray()
	//max := max(12, 192738127398)
	//fmt.Println("Max number is --> ", max)
	hashMap()
	library()

}

func test1(value int, value2 string) {
	fmt.Println("This is a random int from 0 to 100 -->", rand.Intn(100))
	fmt.Println(value, value2)
}

func fucnThatreturnint(v, v1 string) int {
	fmt.Println("Passing data -->", v, v1)
	return 12
}

/* A function can return any number of result */
func swap(string1, string2 string) (string, string) {
	return string1, string2
}

func learnMoreArray() {
	array1 := make([]string, 10)
	array2 := make([]string, 10)

	for index := range array1 {
		array1 = append(array1, "Came here")
		array1[index] = "Hello => " + strconv.Itoa(index)
	}

	for _, value := range array1 {
		fmt.Println(value)
	}

	for i := 0; i < len(array2); i++ {
		array2[i] = "Hello => " + strconv.Itoa(i)
	}

	fmt.Println("These two array are equal?", reflect.DeepEqual(array1, array2))
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func hashMap() {
	hashMap := make(map[int]int)

	value := 1

	for value != 10 {
		hashMap[value] = value * 2
		value++
	}

	//this for loop will print key --> value, only printing the odd keys
	for key, value := range hashMap {
		if key%2 == 0 {
			delete(hashMap, key)
		} else {
			fmt.Println(key, "--> ", value)
		}
	}

	stringListHashMap := make(map[string][]string)

	// This will return a list so we will have the append()
	stringListHashMap["hey"] = append(stringListHashMap["hey"], "not like us not with us --> ")
	stringListHashMap["hey"] = append(stringListHashMap["hey"], "not like us not with uss")

	for _, value := range stringListHashMap {
		fmt.Println(value)
	}

	// Count of character in a String using hashMap
	var data string = "abcdef"

	// byte in go represent ASCII value
	charMap := make(map[byte]int)
	for i := 0; i < len(data); i++ {
		charMap[data[i]]++
	}

	fmt.Println("Char Ascii value array ", charMap)

	/*
		rune is a data type in go that represent unicode code point

		data[index] --> returns byte

	*/
	charMap2 := make(map[rune]int)

	for i := 0; i < len(data); i++ {
		charMap2[rune(data[i])]++
	}

}

func library() {
	// Max Integer
	maxInt := math.MaxInt64
	fmt.Println("Maximum Int--> ", maxInt)
	minInt := -math.MaxInt64
	fmt.Println("Minimum Int--> ", minInt)

	// Max double
	maxDouble := math.MaxFloat64
	fmt.Println("MinDouble --> ", maxDouble)
	minDouble := math.MinInt64
	fmt.Println("MinDouble --> ", minDouble)

	//Sort a string
	var letters string = "bcda"
	arrayLetter := strings.Split(letters, "")
	sort.Strings(arrayLetter)
	fmt.Println(letters, "-->", arrayLetter)

	//now we can join the string
	sortedString := strings.Join(arrayLetter, "")
	fmt.Println("Here is the sorted string --> ", sortedString)

}
