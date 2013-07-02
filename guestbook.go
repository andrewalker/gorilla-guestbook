package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type Login struct {
    user string
    password string
}

type Comment struct {
    user string
    comment string
}


func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home handler");
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Login handler");
}

func GuestbookHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Guestbook handler");
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Logout handler");
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/",          HomeHandler)
    r.HandleFunc("/guestbook", GuestbookHandler)
    r.HandleFunc("/login",     LoginHandler)
    r.HandleFunc("/logout",    LogoutHandler)

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
