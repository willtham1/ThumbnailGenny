package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Use the thumbnailHandler function
	http.HandleFunc("/api/thumbnail", thumbnailHandler)

	//Serve static files from the frontend/dist directory
	fs := http.FileServer(http.Dir("../Frontend/dist"))
	http.Handle("/", fs)

	//Start the server
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

type thumbnailRequest struct {
	Url string `json:"url"`
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest

	// Try to decode the request into the thumbnailRequest struct.
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Got the following url: %s\n", decoded.Url)
}
