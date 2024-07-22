package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world")
	test1(19, "WWW")
	fmt.Println(fucnThatreturnint("String 1", "String 2"))

}

func test1(value int, value2 string) {
	fmt.Println("This is a random int from 0 to 100 -->", rand.Intn(100))
	fmt.Println(value, value2)
}

func fucnThatreturnint(v, v1 string) int {
	fmt.Println("Passing data -->", v, v1)
	return 12
}
