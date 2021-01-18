package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dozi0116/go-nopaste/route"
)

func main() {
	// hosting address
	route.Registration()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("something wrong")
	}
}
