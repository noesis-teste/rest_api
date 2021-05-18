package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	fmt.Println("Rest API - Simple Tutorial")

	handleRequests()

}
