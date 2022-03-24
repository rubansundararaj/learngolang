package main

import (
	"fmt"
	"log"
	"net/http"

	libraryRouter "github.com/rubansundararaj/librarymanagement/router"
)

func main() {
	fmt.Println("Lets make library management apis")
	r := libraryRouter.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000....")
}
