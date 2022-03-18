package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file manipulations")

	content := "The content of the file goes here"

	file, err := os.Create("./createdfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)

	CheckNilErr(err)
	fmt.Println("length: ", length)

	defer file.Close()
	readFile("./createdfile.txt")
}

func readFile(filename string) {

	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data in the file \n", string(dataByte))

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
