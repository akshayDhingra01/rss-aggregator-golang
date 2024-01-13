package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
