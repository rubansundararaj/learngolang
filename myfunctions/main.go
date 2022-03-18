package main

import "fmt"

func main() {
	fmt.Println("Here goes function")

	proResult, _ := proAdder(1, 2, 3, 4, 5)

	fmt.Println(proResult)
}

func proAdder(values ...int) (int, string) {
	proResult := 0

	for _, v := range values {
		proResult += v
	}

	return proResult, "Results"
}
