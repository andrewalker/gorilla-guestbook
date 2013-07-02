package main

import (
    "reflect"
    "fmt"
    "net/http"
    "text/template"
    "github.com/gorilla/mux"
    "github.com/gorilla/schema"
    "github.com/gorilla/sessions"
    "database/sql"
    _ "github.com/lib/pq"
)

type Login struct {
    Username string
    Password string
}

type Comment struct {
    Username string
    Comment string
}

var db *sql.DB
var decoder       = schema.NewDecoder()
var store         = sessions.NewCookieStore([]byte("frase-ultra-secreta-de-encriptacao"))
var homeTemplate  = template.Must(template.New("home").ParseFiles("templates/home.html"))
var loginTemplate = template.Must(template.New("login").ParseFiles("templates/login.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "auth-session")
    username := session.Values["username"]
    data := make(map[string]interface{})
    if reflect.TypeOf(username).String() == "string" {
        data["logged_in"] = true
        data["username"] = username
    }
    if err := homeTemplate.ExecuteTemplate(w, "home", data); err != nil {
        fmt.Println(err);
    }
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    data := make(map[string]interface{})

    data["wrong_password"] = len(query["wrong_password"]) > 0 && query["wrong_password"][0] == "1"

    if err := loginTemplate.ExecuteTemplate(w, "login", data); err != nil {
        fmt.Println(err);
    }
}

func DoLoginHandler(w http.ResponseWriter, r *http.Request) {
    var password string
    page := "/login"
    login := new(Login)
    session, _ := store.Get(r, "auth-session")

    r.ParseForm();
    decoder.Decode(login, r.PostForm)

    err := db.QueryRow("SELECT password FROM users WHERE username=?", login.Username).Scan(&password)

    if err == sql.ErrNoRows {
        page = page + "?wrong_password=1"
    } else {
        session.Values["username"] = login.Username
    }

    session.Save(r, w)

    http.Redirect(w, r, page, http.StatusFound)
}

func GuestbookHandler(w http.ResponseWriter, r *http.Request) {
    // após gravar no guestbook, redirecione
    http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    // após logout, redirecione
    http.SetCookie(w, &http.Cookie{Name: "auth-session", MaxAge: -1, Path: "/"})
    http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
    var err error
    db, err = sql.Open("postgres", "user=andre dbname=guestbook")

    if err != nil {
        fmt.Println(err)
    }

    r := mux.NewRouter()

    r.HandleFunc("/",          HomeHandler)
    r.HandleFunc("/guestbook", GuestbookHandler)
    r.HandleFunc("/login",     LoginHandler)
    r.HandleFunc("/do_login",  DoLoginHandler)
    r.HandleFunc("/logout",    LogoutHandler)

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
