package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	ruban := User{"Ruban", "ruban@metakartz.com", true, 100}

	fmt.Println(ruban.Name)
	fmt.Printf("%+v\n", ruban)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
