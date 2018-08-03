package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/SomniaStellarum/StellarUtilities/slog"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("Path requested: ", r.URL.String())
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Panic("Template Error: ", err)
	}
}

type logRequester struct {
	h http.Handler
}

func (l logRequester) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("Static Path requested: ", r.URL.String())
	l.h.ServeHTTP(w, r)
}

func logRequest(f http.Handler) http.Handler {
	return logRequester{h: f}
}

func main() {
	slog.SetDebug()
	http.HandleFunc("/", index)
	http.Handle("/css/", logRequest(http.StripPrefix("/css" ,http.FileServer(http.Dir("./css")))))
	http.Handle("/js/", logRequest(http.StripPrefix("/js" ,http.FileServer(http.Dir("./js")))))
	http.ListenAndServe(":8080", nil)
}
