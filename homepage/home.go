package homepage

import (
	"log"
	"net/http"
	"time"
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

// LogCallsDurationMiddleware as a MiddleWare in go do h.Logger(h.Home)
func (h *Handlers) LogCallsDurationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		h.logger.Printf("Request handled in %s\n", time.Now().Sub(startTime))
	}
}

// SetupRoutes as a MiddleWare in go do h.Logger(h.Home)
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.LogCallsDurationMiddleware(h.Home)) // enable the Logger middleware here
}
