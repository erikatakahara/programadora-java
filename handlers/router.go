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
    r.HandleFunc("/labs/notification", NotificationLab)
    r.HandleFunc("/labs/notification/push", PushNotificationLab)
    http.Handle("/", r)
}