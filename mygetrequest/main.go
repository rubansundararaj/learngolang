package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Lets make a get request")
	//PerformGetRequest()

	//PerformPostJsonRequest()

	PerformFormPostRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	resp, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("Content length is: ", resp.ContentLength)

	var responseStrings strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)

	bytesCount, _ := responseStrings.Write(content)

	fmt.Println(bytesCount)
	fmt.Println(responseStrings.String())

	fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`{
		"courseName" : "lets go with go lang",
		"price" : 500,
		"platform" : "letscode.xyz"
	}`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformFormPostRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("firstname", "ruban")
	data.Add("lastname", "t")
	data.Add("email", "ruban@autocodegen.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
