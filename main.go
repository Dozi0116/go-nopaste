package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Dozi0116/go-nopaste/route"
)

func main() {
	// hosting address
	router := httprouter.New()
	route.Registration(router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("something wrong")
	}
}
