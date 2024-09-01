package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	var ptr *int

	fmt.Println("the value of the pointer is ", ptr)
	fmt.Printf("The data type is: %T \n", ptr)

	var value int = 43
	var ptr1 *int = &value
	fmt.Println("the value of ptr1 is", ptr1)
	fmt.Println("The value of value is: ", *ptr1)
}
