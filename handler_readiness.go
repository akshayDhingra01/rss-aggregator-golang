package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // define a http handler as go library expects
	fmt.Println("Helllo")
	log.Fatal("JELLEO")
	fmt.Printf("Server started working on Port\n")
	// Log.Println("Helllo")
	respondWithJson(w, 200, struct{}{})

}
