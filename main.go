package main

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[Request] %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := ":8181"

	initialiseRoutes()
	initDB()

	wrappedMux := loggingMiddleware(http.DefaultServeMux)

	log.Printf("Server started on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, wrappedMux))
}
