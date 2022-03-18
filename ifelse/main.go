package main

import "fmt"

func main() {
	fmt.Println("If else starts here")

	loginCount := 2
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Fat user"
	} else {
		result = "he aint human"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number less than 3")
	} else {

	}

}
