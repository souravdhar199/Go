package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("Hello world")
	test1(19, "WWW")
	fmt.Println(fucnThatreturnint("String 1", "String 2"))

	value1, value2 := swap("Hello", "name")

	fmt.Println(value2, value1)

	// This function from another file
	//arrays()
	learnMoreArray()

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
