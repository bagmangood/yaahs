package yaahs

import (
	"encoding/json"
	"net/http"
)

const (
	contentType     = "Content-Type"
	contentTypeJSON = "application/json"
)

type JSONResponse struct {
	StatusCode int
	Body       interface{}
}

func SendJSONResponse(w http.ResponseWriter, resp JSONResponse) {
	w.Header().Set(contentType, contentTypeJSON)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
}
