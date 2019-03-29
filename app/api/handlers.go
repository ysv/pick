package api

import (
	"net/http"

	"github.com/ysv/pick/app"
)

func HealthHandler(w http.ResponseWriter, _ *http.Request){
	if err := app.GetDB().Ping(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.WriteHeader(http.StatusOK)
}
