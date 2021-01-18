package c

import (
	"html/template"
	"log"
	"net/http"
)

func Page(w http.ResponseWriter, req *http.Request) {
	render, err := template.ParseFiles("template/page.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}

	if err := render.Execute(w, struct {
		Message string
	}{
		Message: "umekomi",
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}
