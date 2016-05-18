package handlers

import (
    "html/template"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("webapp/templates/base.html", "webapp/pages/index.html"))
    err := t.Execute(w, "home")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func Curriculum(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("webapp/templates/base.html", "webapp/pages/curriculum.html"))
    err := t.Execute(w, "curriculum")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func TravelLab(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("webapp/templates/base.html", "webapp/pages/travel.html"))
    err := t.Execute(w, "travel")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}