package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	// defer command is evaluated when function call is finished.
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for idx, number := range numbers {
		fmt.Println("index: ", idx)
		total += number
	}
	return total
}

func main() {
	fmt.Println("Hello World")
	// Variable Declaration
	name := "Scott"
	// Variable Update
	name = "Scott Suk"
	length, _ := lenAndUpper(name)
	fmt.Println(length)
	fmt.Println(name)
	sum_result = superAdd(1, 2, 3, 4, 5)
	fmt.Println(sum_result)
}
