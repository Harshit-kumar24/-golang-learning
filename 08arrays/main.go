package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays")

	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "banana"
	fruits[2] = "Mango"
	fruits[3] = "Grapes"
	fruits[4] = "Orange"

	fmt.Println("The fruits are:", fruits)
	fmt.Println("the length of fruits is: ", len(fruits))

	var veggies = [3]string{"mushrooms", "tomatoes", "potatoes"}
	fmt.Println("The veggies are:", veggies)
}
