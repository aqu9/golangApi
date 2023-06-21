package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var storedString string

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello world")
    fmt.Println("Endpoint Hit: homePage")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		String string `json:"string"`
	}{
		String: storedString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		String string `json:"string"`
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedString = requestBody.String

	response := struct {
		Message string `json:"message"`
	}{
		Message: "String stored successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/getString", GetHandler)
	http.HandleFunc("/addString", PostHandler)
	http.HandleFunc("/", homePage)

	fmt.Println("Server listening on http://localhost:10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}
