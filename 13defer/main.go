package main

import "fmt"

func main() {
	defer fmt.Println("Introduction to defer")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Introduction to defer line second")

}
