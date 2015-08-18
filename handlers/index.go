package handlers

import (
    "appengine"
    "appengine/user"
    
    "html/template"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("webapp/index.html"))
    c := appengine.NewContext(r)
    u := user.Current(c)
    err := t.Execute(w, u)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}