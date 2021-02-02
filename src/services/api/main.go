package main

import (
	"encoding/json"
	"fmt"
	"fs/src/services/api/payload"
	"fs/src/services/api/response"
	"github.com/gorilla/mux"
	"net/http"
)

func renderResponse(w http.ResponseWriter, response response.AcceptResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.HttpCode)
	bytes, _ := json.Marshal(response)
	fmt.Fprintf(w, string(bytes))
}

func handleAccept(w http.ResponseWriter, r *http.Request) {
	acceptPayload, err := payload.NewAcceptPayload(r)
	if err != nil {
		renderResponse(w, response.AcceptResponse{
			HttpCode: 400,
			Message:  "Reading request body has failed.",
		})
	}

	renderResponse(w, acceptPayload.Validate())
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/accept", handleAccept).Methods("POST")
	http.ListenAndServe(":10000", router)
}

func main() {
	handleRequests()
}
