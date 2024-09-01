package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("================please rate our website from one to five=============")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	fmt.Println("Thanks for rating", input)

	// Parse the input to a float
	rating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error occurred while parsing:", err)
	} else {
		fmt.Println("You rated the website:", rating)
	}
}
