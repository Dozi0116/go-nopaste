package c

import (
	"net/http"
	"io"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World\n")
}
