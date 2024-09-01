package main

import (
	"fmt"
)

func main() {
	fmt.Println("Introduction to maps")

	languages := make(map[string]string)
	languages["en"] = "english"
	languages["de"] = "deutsch"
	languages["fr"] = "français"

	fmt.Println("the value of the map is", languages)
	delete(languages, "en")
	fmt.Println("the value of the map is", languages)

	for key, value := range languages {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}
