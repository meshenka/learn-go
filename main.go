package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	homepage "./homepage"
	server "./server"
)

var (
	// MlgCertFile unused yet
	MlgCertFile = os.Getenv("MLG_CERT_FILE")
	// MlgKeyFile unused yet
	MlgKeyFile = os.Getenv("MLG_KEY_FILE")
	// MlgServiceAddr service address,could be :8080
	MlgServiceAddr = os.Getenv("MLG_SERVICE_ADDR")
)

func main() {
	fmt.Println("Starting")

	mux := http.NewServeMux()

	mux.HandleFunc("/", homepage.HomeHandler)

	srv := server.New(mux, MlgServiceAddr)

	err := srv.ListenAndServeTLS(MlgCertFile, MlgKeyFile)

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
