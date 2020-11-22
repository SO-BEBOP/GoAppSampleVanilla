package main

import (
	"GoAppSampleVanilla/internal/route"
	"GoAppSampleVanilla/pkg/utils"
	"net/http"
	"time"
)

func main() {
	utils.P("GoAppSampleVanilla", utils.Version(), "started at", utils.Config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(utils.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", route.Index)
	// error
	mux.HandleFunc("/err", route.Err)

	// defined in route_auth.go
	mux.HandleFunc("/login", route.Login)
	mux.HandleFunc("/logout", route.Logout)
	mux.HandleFunc("/signup", route.Signup)
	mux.HandleFunc("/signup_account", route.SignupAccount)
	mux.HandleFunc("/authenticate", route.Authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", route.NewThread)
	mux.HandleFunc("/thread/create", route.CreateThread)
	mux.HandleFunc("/thread/post", route.PostThread)
	mux.HandleFunc("/thread/read", route.ReadThread)

	// starting up the server
	server := &http.Server{
		Addr:           utils.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(utils.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(utils.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
