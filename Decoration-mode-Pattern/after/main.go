package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func Logger(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		handler(w, r)
		log.Printf("url: %s, elase: %v", r.URL, time.Since(now))

	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World."))
}

func HowAreYou(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I'am fine!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", Logger(HelloWorld))
	mux.HandleFunc("/how", Logger(HowAreYou))

	srv := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	fmt.Println("listen at: ", srv.Addr)
	srv.ListenAndServe()
}
