package main

import (
	"fmt"
)


func main() {
	fmt.Println("Enter something: ")	

	var smth string

	fmt.Scan(&smth)
	fmt.Println("Enter another thing: ")

	var smthElse string
	fmt.Scanf(&smthElse)

	fmt.Println(smth + " " + smthElse)
}