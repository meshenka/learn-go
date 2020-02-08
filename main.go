package main

import (
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
	logger := log.New(os.Stdout, "MLG", log.LstdFlags|log.Lshortfile)
	logger.Println("Starting")

	mux := http.NewServeMux()

	//This is Dependencies injection in Go
	h := homepage.NewHandlers(logger)

	h.SetupRoutes(mux)

	srv := server.New(mux, MlgServiceAddr)

	err := srv.ListenAndServeTLS(MlgCertFile, MlgKeyFile)

	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
