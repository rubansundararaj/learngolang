package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learn slices")

	var fruitList = []string{"Apple", "Tomato", "peach"}
	fmt.Printf("%T", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 423
	highScores[2] = 543
	highScores[3] = 645

	fmt.Println(highScores)

	highScores = append(highScores, 656, 898, 321)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove value from slice based on index

	var courses = []string{"reactjs", "javascript", "swift"}

	fmt.Println(courses)

	index := 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
