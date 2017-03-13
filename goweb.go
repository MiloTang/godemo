package main

import (
	"html/template"
	"log"
	"miloblog/golib/gocs"
	"miloblog/wr"
	"net/http"
	"os"
	"strings"
)

type Page struct {
	Title    string
	Lists    []Context
	Next     string
	Previous string
	Token    string
	Info     string
	Details  string
	Username string
	Password string
	Error    error
}
type Context struct {
	Introduction string
	Link         string
}

var (
	p                         = &Page{}
	cs    *gocs.CookieSession = nil
	Debug bool                = true
	err   error               = nil
	file  *os.File
)

func entry(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		wr.Index(w, r)
	} else {
		urls := strings.Split(r.URL.Path, "/")
		switch urls[1] {
		case "life":
			wr.Life(w, r)
		case "admin":
			wr.AdminLogin(w, r, cs)
		case "manuallist":
			wr.Manual(w, r)
		case "index":
			wr.Index(w, r)
		case "blog", "manual":
			wr.Details(w, r, urls[1])
		case "editor":
			wr.Editor(w, r, cs)
		case "delsession":
			wr.Delsession(w, r, cs)
		default:
			ErrorPage(w, r)
		}
	}
}
func ErrorPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("error.html")
	t.Execute(w, nil)
}
func main() {
	gocs.MaxLT = 3600
	gocs.CookieName = "Miloblog"
	cs = gocs.NewCookieSession()
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/images/", http.FileServer(http.Dir("static")))
	http.Handle("/bootstrap/", http.FileServer(http.Dir("static")))
	http.Handle("/layoutitlib/", http.FileServer(http.Dir("static")))
	http.Handle("/wysiwyg/", http.FileServer(http.Dir("static")))
	http.Handle("/fonts/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/", entry)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
