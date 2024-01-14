package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) { // define a http handler as go library expects
	respondWithError(w, 400, "Something Went Wrong")
}