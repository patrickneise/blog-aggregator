package main

import "net/http"

func (cfg *apiConfig) handlerHealhtz(w http.ResponseWriter, req *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, response{Status: "ok"})
}
