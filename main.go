package main

import (
	"net/http"
	"log"

	"github.com/Dozi0116/go-nopaste/C"
)

func main() {
	// hosting address
	http.HandleFunc("/", C.Hello)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Exit("ListenAndServe: ", err.String())
	}
}
