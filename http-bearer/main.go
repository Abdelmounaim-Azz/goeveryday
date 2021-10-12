package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func makeAuth(token string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if b := r.Header.Get("Authorization"); len(b) == 0 || b != "Bearer: "+token {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
		} else {
			next(w, r)
		}
	}
}
func main() {
	token := os.Getenv("TOKEN")
	next := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Authorized"))
	}
	http.HandleFunc("/customer/", makeAuth(token, next))
	tcpPort := 8080
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", tcpPort),
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}
	log.Printf("Listening on: %d", tcpPort)
	log.Fatal(s.ListenAndServe())
}
