package main

import (
	"fmt"
	"log"
	"net/http"
	"userinterface/router"
)

func main() {
	fmt.Println("Welcome")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r))
	return
}
