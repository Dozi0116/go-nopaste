package c

import (
	"net/http"
	"io"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World\n")
}

func Hoge(w http.ResponseWriter, req *http.Request) {
	page := `
<h1>No Paste</h1>
<p>text here!</P>
<input type="text" id="original" name="original">
	`
	io.WriteString(w, page)
}
