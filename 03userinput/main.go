package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "this is a welcome messgage"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	//comman okay syntax OR comma error
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
	fmt.Println(err)

}
