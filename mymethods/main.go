package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	ruban := User{"Ruban", "ruban@metakartz.com", true, 100}

	fmt.Println(ruban.Name)
	ruban.GetStatus()
	fmt.Println(ruban.Name)
	ruban.GetSus()
	fmt.Println(ruban.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {

	u.Name = "Bang bang"

	fmt.Println(u.Name)

}

func (u *User) GetSus() {
	u.Name = "bung bung"

	fmt.Println(u.Name)
}
