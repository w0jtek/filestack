package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type AcceptPayload struct {
	ImageURL string `json:"imageUrl"`
}

func handleAccept(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var payload AcceptPayload
	json.Unmarshal(reqBody, &payload)

	fmt.Fprintf(w, payload.ImageURL)
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/accept", handleAccept).Methods("POST")
	http.ListenAndServe(":10000", router)
}

func main() {
	handleRequests()
}
