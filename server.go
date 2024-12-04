package main

import (
	"net/http"

	handlers "codeujuzi.github.io/http"
	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func routes(){
	r.HandleFunc("/",handlers.HomePageHandler)
	r.HandleFunc("/signin",handlers.SigninHandler)
	r.HandleFunc("/login",handlers.LoginHandler)
	r.HandleFunc("/register",handlers.RegisterHandler)
	r.HandleFunc("/lessons",handlers.LessonsHandler)
	r.HandleFunc("/course-outline",handlers.CourseHandler)
	r.HandleFunc("/logout",handlers.LogoutHandler)
	r.HandleFunc("/profile",handlers.ProfileHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/",http.FileServer(http.Dir("./template/css"))))
	// http.Handle("/",r)
}

var server = &http.Server{
	Handler: r,
	Addr: "127.0.0.1:8000",
}