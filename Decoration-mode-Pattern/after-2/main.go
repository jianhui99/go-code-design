package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("url: %s, elase: %v", r.URL, time.Since(now))

	}

	return http.HandlerFunc(fn)
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

	mux.HandleFunc("/hello", HelloWorld)
	mux.HandleFunc("/how", HowAreYou)

	srv := http.Server{
		Addr:    ":3000",
		Handler: Logger(mux),
	}

	fmt.Println("listen at: ", srv.Addr)
	srv.ListenAndServe()
}
