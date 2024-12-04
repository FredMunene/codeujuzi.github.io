package main

import (
	"net/http"

	handlers "codeujuzi.github.io/http"
	"codeujuzi.github.io/middleware"
	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func routes(){
	r.HandleFunc("/",handlers.HomePageHandler)
	r.HandleFunc("/signin",handlers.SigninHandler)
	r.HandleFunc("/signup",handlers.SignupHandler)

	r.HandleFunc("/login",handlers.LoginHandler)
	r.HandleFunc("/register", handlers.RegisterHandler)
	r.HandleFunc("/logout",handlers.LogoutHandler)

	r.HandleFunc("/lessons",middleware.TokenValidation(handlers.LessonsHandler))
	r.HandleFunc("/course-outline",handlers.CourseHandler)
	r.HandleFunc("/profile",handlers.ProfileHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/",http.FileServer(http.Dir("./template/css"))))
	// http.Handle("/",r)
}

var server = &http.Server{
	Handler: r,
	Addr: "127.0.0.1:8000",
}