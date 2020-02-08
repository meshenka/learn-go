package homepage

import "net/http"

const message = "Hello Wold!"

// HomeHandler server the Home URI
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
