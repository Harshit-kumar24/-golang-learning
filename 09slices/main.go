package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Introduction to slices")

	var fruits = []string{}
	fmt.Printf("Type of fruits is: %T \n", fruits)

	fruits = append(fruits, "apple", "mango", "banana", "grapes", "orange")
	// fruits = append(fruits[0:4])
	fmt.Println(fruits)

	highScores := make([]int, 4)
	highScores[0] = 2
	highScores[1] = 5
	highScores[2] = 7
	highScores[3] = 10

	highScores = append(highScores, 34, 45, 67, 34)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("After sorting: ", highScores)

	//remove element based on inddex
	var index int = 1
	fruits = append(fruits[:index], fruits[index+1:]...)
	fmt.Println(fruits)

}
