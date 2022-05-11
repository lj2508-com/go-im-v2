package main

import (
	_ "github.com/mattn/go-sqlite3"
	"go-im-v2/ctrl"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/chat", ctrl.Chat)
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	registerView()
	http.ListenAndServe(":8090", nil)
}

func registerView() {
	glob, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range glob.Templates() {
		name := v.Name()
		http.HandleFunc(name, func(writer http.ResponseWriter, request *http.Request) {
			glob.ExecuteTemplate(writer, name, nil)
		})
	}
}
