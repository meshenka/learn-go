package main

import (
	"fmt"
	"log"
	"net/http"
)

const message = "Hello Wold!"

func main() {
	fmt.Println(message)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
