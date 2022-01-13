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

func main() {
	fmt.Println("Hello World")
	// Variable Declaration
	name := "Scott"
	// Variable Update
	name = "Scott Suk"
	length, _ := lenAndUpper(name)
	fmt.Println(length)
	fmt.Println(name)
}
