package handlers

import (
    "github.com/gorilla/mux"
    "net/http"
)

func init() {
    r := mux.NewRouter()
    
    r.HandleFunc("/", Index)
    http.Handle("/", r)
}