package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customer/{id:[-a-zA-Z_0-9.]+}", func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		id := v["id"]
		if id == "abdelmounaim" {
			w.Write([]byte(fmt.Sprintf("Found customer: %s", id)))
		} else {
			http.Error(w, fmt.Sprintf("Customer %s not found", id), http.StatusNotFound)
		}
	})
	tcpPort := 8080
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", tcpPort),
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		Handler: router,
	}
	log.Printf("Listening on: %d", tcpPort)
	log.Fatal(s.ListenAndServe())
}
