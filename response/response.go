package response

import (
	"net/http"
	"encoding/json"
)

type StatusCode int

type Response struct {
	Status StatusCode  `json:"status"`
	Result interface{} `json:"result"`
}

func CreateJSON (data interface{}, status StatusCode) []byte {
	response := Response{status, data}

	res, err := json.Marshal(response)

	if err != nil {
		panic(err) 
	}

	return res
}

func RenderJSONOK(w *http.ResponseWriter, data interface{}) {
	RenderJSON(w, data, http.StatusOK)
}

func RenderJSON(w *http.ResponseWriter, data interface{}, status StatusCode) {
	res := CreateJSON(data, status)
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(res)
}
