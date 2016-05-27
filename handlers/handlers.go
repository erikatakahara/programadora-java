package handlers

import (
    "html/template"
    "net/http"
    "io/ioutil"
    "bytes"

    "appengine"
    "appengine/urlfetch"

    "infra"
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

func NotificationLab(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("webapp/templates/base.html", "webapp/pages/notification.html"))
    err := t.Execute(w, "notification")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


func PushNotificationLab(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    url := "https://android.googleapis.com/gcm/send"
    json, _ := ioutil.ReadAll(r.Body)
    c.Infof("response Status: %v", string(json))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "key=" + infra.Conf().GcmPushKey)

    client := &http.Client{
        Transport: &urlfetch.Transport{Context: c},
    }
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}