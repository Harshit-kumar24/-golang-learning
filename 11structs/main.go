package main

import "fmt"

func main() {
	fmt.Println("Introduction to structs")

	var random User = User{"random", "random@gmail.com", true, 23}
	fmt.Println(random)
	fmt.Println(random.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
