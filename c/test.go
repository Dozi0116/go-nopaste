package c

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func Page(w http.ResponseWriter, req *http.Request) {
	render := template.ParseFiles("../template/page.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}

	if err := t.Execute(w, struct {
		Message string
	}{
		Message: "umekomi",
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}
