package main

import "fmt"

func main() {
	fmt.Println("Introduction to if else statements")

	if num := 3; num <= 10 {
		fmt.Println("the value fo the number is 3")
	} else {
		fmt.Println("the value of the number is less than 10 but not 3")
	}
}
