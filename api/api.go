package api

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

//func Handler(w http.ResponseWriter, r *http.Request){
//	m := RegisterRoutes()
//	//m.Handle()
//}

func RegisterRoutes(router *mux.Router) {
	//r := mux.NewRouter()
	//r.Use()
	//r.PathPrefix(pathPrefix)
	router.HandleFunc("/pick", Pick)
	router.HandleFunc("/health", Health)
}

func Pick(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "pick pick")
}

func Health(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "health health")
}

//Handler
