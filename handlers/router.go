package handlers

import (
    "github.com/gorilla/mux"
    "net/http"
)

func init() {
    r := mux.NewRouter()
    
    r.HandleFunc("/", Index)
    r.HandleFunc("/curriculum", Curriculum)
    r.HandleFunc("/labs/travel", TravelLab)
    http.Handle("/", r)
}