package c

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Dozi0116/go-nopaste/response"
)

func GetMessageList(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	response.RenderJSONOK(&w, "OK!")
}

func PostMessage(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	response.RenderJSON(&w, "coming soon", http.StatusNotFound)
}

func GetMessageById(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	response.RenderJSON(&w, "coming soon", http.StatusNotFound)
}

