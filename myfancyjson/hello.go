package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("Here we play with json")
	//EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Plan     string   `json:"plan"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abc123452", []string{"web-dev", "js"}},
		{"This that Bootcamp", 159, "LearnCodeOnline.in", "abc123452", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"name": "MERN Bootcamp",
			"price": 199,
			"plan": "LearnCodeOnline.in",
			"tags": ["web-dev","js"]
		}
		`)
	var lcoCourses course

	checkIfValid := json.Valid(jsonDataFromWeb)

	if checkIfValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON not valid")
	}

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("%v and %v", k, v)
	}
}
