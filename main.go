package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	server "./server"
)

const message = "Hello Wold!"

var (
	// MlgCertFile unused yet
	MlgCertFile = os.Getenv("MLG_CERT_FILE")
	// MlgKeyFile unused yet
	MlgKeyFile = os.Getenv("MLG_KEY_FILE")
	// MlgServiceAddr service address,could be :8080
	MlgServiceAddr = os.Getenv("MLG_SERVICE_ADDR")
)

func main() {
	fmt.Println(message)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	srv := server.New(mux, MlgServiceAddr)

	err := srv.ListenAndServeTLS(MlgCertFile, MlgKeyFile)

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
