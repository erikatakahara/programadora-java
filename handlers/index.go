package handlers

import (
    "html/template"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("webapp/templates/base.html", "webapp/pages/index.html"))
    err := t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}