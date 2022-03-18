package main

import "fmt"

func main() {
	fmt.Println("Hello pointers")

	var ptr *int
	fmt.Println(ptr)

	myNumber := 23

	var add = &myNumber

	fmt.Println(&add)
	fmt.Println(*add)
	fmt.Println(add)
}
