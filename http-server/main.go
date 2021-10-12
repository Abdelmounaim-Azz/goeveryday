package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("======================================================")
		log.Printf("Query-string: %s", r.URL.RawQuery)
		log.Printf("Path: %s", r.URL.Path)
		log.Printf("Method: %s", r.Method)
		log.Printf("Path: %s", r.Host)
		for k, v := range r.Header {
			log.Printf("Header %s=%s", k, v)
		}
		if r.Body != nil {
			body, _ := ioutil.ReadAll(r.Body)
			log.Printf("Body: %s", string(body))
		}
		log.Printf("======================================================")
		w.WriteHeader(http.StatusAccepted)
	})
	port := 8080
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Listening on port: %d\n", port)
	log.Fatal(s.ListenAndServe())
}
