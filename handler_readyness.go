package main 

import "net/http"

func handlerReadyness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}