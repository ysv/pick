package api

import (
	"net/http"

	"github.com/ysv/pick/app"
)

func HealthHandler(w http.ResponseWriter, _ *http.Request){
	if err := app.GetDB().Health(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.WriteHeader(http.StatusOK)
}

func PickHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}
