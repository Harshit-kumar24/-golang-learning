package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to study of time in golang")

	presentTime := time.Now()
	fmt.Println("The current time is: ", presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.April, 2, 12, 34, 34, 23, time.UTC).Format("02-01-2006 Monday")
	fmt.Println(createDate)

}
