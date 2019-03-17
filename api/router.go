package api

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)
func Router() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/route", Hey)
	r.HandleFunc("/", Yo)
	//r.PathPrefix()
	return r
}

func Hey(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "cat cat cat")
}

func Yo(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "cat cat cat")
}
