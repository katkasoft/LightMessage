package main

import (
	"lightmessage/api"
	"log"
	"net/http"
)

func initialiseRoutes() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/login.html")
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/register.html")
	})

	http.HandleFunc("/api/register", api.Register)
	http.HandleFunc("/api/login", api.Login)

	log.Printf("Routes initialised")
}
