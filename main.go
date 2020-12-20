package main

import (
	"fmt"
	"log"
	"net/http"

	"./c"
)

func main() {
	// hosting address
	http.HandleFunc("/", c.Hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("something wrong")
	}
}
