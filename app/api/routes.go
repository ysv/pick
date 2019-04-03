package api

import (
	"fmt"
	"github.com/ysv/pick/app"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// GET, OPTIONS /api/health
	router.HandleFunc("/health", HealthHandler).Methods(http.MethodGet, http.MethodOptions)

	// GET /api/pick
	router.HandleFunc("/pick", PickHandler).Methods(http.MethodGet)

	router.Use(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formattedR := fmt.Sprintf("%s: %s", r.Method, r.RequestURI)
		app.GetLogger().Infoln(formattedR)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
