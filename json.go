package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError( w http.ResponseWriter, code int, msg string)  { // http handlers use this http.ResposeWriter
	if code > 499 {
		log.Println("Responding with Error 5xx:", msg)
	}
	type errorResponse struct {
		Error string `json: "error"` // json error is used to write so to say that in key-value pair, value of key should be error
	}
	respondWithJson(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJson( w http.ResponseWriter, code int, payload interface{})  { // http handlers use this http.ResposeWriter
	data, err := json.Marshal(payload)
	if err!= nil {
		log.Printf("Failed to Marshal Json Response : %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
