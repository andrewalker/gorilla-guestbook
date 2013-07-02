package main

import (
    "fmt"
    "net/http"
    "text/template"
    "github.com/gorilla/mux"
    "github.com/gorilla/schema"
)

type Login struct {
    Username string
    Password string
}

type Comment struct {
    Username string
    Comment string
}

var decoder       = schema.NewDecoder()
var homeTemplate  = template.Must(template.New("home").ParseFiles("templates/home.html"))
var loginTemplate = template.Must(template.New("login").ParseFiles("templates/login.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if err := homeTemplate.ExecuteTemplate(w, "home", nil); err != nil {
        fmt.Println(err);
    }
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if err := loginTemplate.ExecuteTemplate(w, "login", nil); err != nil {
        fmt.Println(err);
    }
}

func DoLoginHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm();
    login := new(Login)
    decoder.Decode(login, r.PostForm)

    switch login.Username {
        case "andre": {
            if login.Password == "123" {
                fmt.Println("Usuário é André.");
            } else {
                fmt.Println("Senha errada.");
            }
        }
        case "pedro": {
            if login.Password == "123" {
                fmt.Println("Usuário é Pedro.");
            } else {
                fmt.Println("Senha errada.");
            }
        }
        case "lucas": {
            if login.Password == "123" {
                fmt.Println("Usuário é Lucas.");
            } else {
                fmt.Println("Senha errada.");
            }
        }
        default: fmt.Println("O usuário não existe.");
    }

    http.Redirect(w, r, "/login", http.StatusFound)
}

func GuestbookHandler(w http.ResponseWriter, r *http.Request) {
    // após gravar no guestbook, redirecione
    http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    // após logout, redirecione
    http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/",          HomeHandler)
    r.HandleFunc("/guestbook", GuestbookHandler)
    r.HandleFunc("/login",     LoginHandler)
    r.HandleFunc("/do_login",  DoLoginHandler)
    r.HandleFunc("/logout",    LogoutHandler)

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
