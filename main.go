package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/ethereum/go-ethereum/common"
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
		log.Panic("Template error on index: ", err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("Path requested: ", r.URL.String())
	err := tpl.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		log.Panic("Template error on register: ", err)
	}
}

func doRegister(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("doRegister called")

	ethAddress := common.HexToAddress(r.PostFormValue("ethAddress"))
	email := r.PostFormValue("email")

	err := PutEmail(ethAddress, email)
	if err != nil {
		log.Panic("Error calling PutEmail: ", err)
	}

	slog.DebugPrint("I guess it's in the DB!")
}

func getUserEmail(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("getUserEmail called")

	ethAddress := common.HexToAddress(r.PostFormValue("ethAddress"))

	email, err := GetEmail(ethAddress)
	if err != nil {
		log.Panic("Error calling GetEmail: ", err)
	}

	w.Write([]byte(email))
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
	http.HandleFunc("/register", register)
	http.HandleFunc("/doRegister", doRegister)
	http.HandleFunc("/getUserEmail", getUserEmail)
	http.Handle("/css/", logRequest(http.StripPrefix("/css" ,http.FileServer(http.Dir("./css")))))
	http.Handle("/js/", logRequest(http.StripPrefix("/js" ,http.FileServer(http.Dir("./js")))))
	http.ListenAndServe(":8080", nil)
}
