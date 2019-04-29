package api

import (
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

func (api *API) Routes() *mux.Router {
	// register routes
	r := mux.NewRouter()
	r.Handle("/api/pick", NewCollector(api.database)).Methods(http.MethodGet)

	r.Handle("/api/session", HandlerFunc(api.GetSession)).Methods(http.MethodGet)
	r.Handle("/api/session", HandlerFunc(api.CreateSession)).Methods(http.MethodPost)
	r.Handle("/api/session", HandlerFunc(api.DeleteSession)).Methods(http.MethodDelete)

	r.Handle("/api/health", HandlerFunc(api.Health)).Methods(http.MethodGet)

	// pick.js handler
	box := packr.NewBox("./../../assets")
	r.Path("/pick.js").Handler(serveTrackerFile(&box))

	return r
}

func serveTrackerFile(box *packr.Box) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Tk", "N")
		next := serveFile(box, "pick.js")
		next.ServeHTTP(w, r)
	})
}
