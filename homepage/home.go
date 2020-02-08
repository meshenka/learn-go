package homepage

import (
	"log"
	"net/http"
)

const message = "Hello Wold!"

// Handlers DI
type Handlers struct {
	logger *log.Logger
}

// Home server the Home URI. This is a method on Handlers type
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("In Home")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// NewHandlers create Handlers
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
