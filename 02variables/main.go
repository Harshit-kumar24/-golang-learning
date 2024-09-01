package main

import "fmt"

const LoginToken string = "Login token" //capital letter means public

func main() {
	var username string = "harshit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T line ends here \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T line ends here \n", isLoggedIn)

	var smallVal int = 4 //0 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T line ends here \n", smallVal)

	var smallFloat float32 = 34.4
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T line ends here \n", smallFloat)

	//aliases and default values

	//var style
	var website = "learncodeonline.com"
	fmt.Println(website)

	//no var style
	numberOfUsers := 2093483
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)

}
