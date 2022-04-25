package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func WithLogger(handler http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, handler)
}

func WithCORS(router *mux.Router) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Accept"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})
	return handlers.CORS(headers, methods, origins)(router)
}
