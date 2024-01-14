package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // define a http handler as go library expects
	respondWithJson(w, 200, struct{}{})

}
