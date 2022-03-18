package main

import "fmt"

func main() {
	fmt.Println("Learn maps here")

	language := make(map[string]string)

	language["JS"] = "javascript"
	language["RB"] = "ruby"
	language["PY"] = "python"

	fmt.Println(language["RB"])

	delete(language, "RB")

	fmt.Println(language)

	// Lets LOOP through
	for key, value := range language {
		fmt.Println(key)
		fmt.Println(value)
	}
}
