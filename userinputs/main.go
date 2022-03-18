package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello there"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating....")

	input, err := reader.ReadString('\n')

	fmt.Println("This is the waht we read, ", input)

	fmt.Println("This is the waht we read, ", err)
}
