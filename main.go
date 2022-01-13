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

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canISmoke(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println("Hello World")
	// Variable Declaration
	name := "Scott"
	// Variable Update
	name = "Scott Suk"

	// Functions
	length, _ := lenAndUpper(name)
	fmt.Println(length)
	fmt.Println(name)

	// For loop
	sum_result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(sum_result)

	// If Else
	if canIDrink(16) {
		fmt.Println("You can drink!")
	} else {
		fmt.Println("You cannot drink!")
	}

	// Switch
	if canISmoke(10) {
		fmt.Println("You can smoke!")
	} else {
		fmt.Println("You cannot smoke!")
	}
}
