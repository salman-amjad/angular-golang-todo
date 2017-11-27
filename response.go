package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	writer http.ResponseWriter
}

func (r Response) SendJSON(data interface{}) {
	r.writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(r.writer).Encode(data); err != nil {
		panic(err)
	}
}
