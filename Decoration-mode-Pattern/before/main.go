package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World."))

	log.Printf("url: %s, elase: %v", r.URL, time.Since(now))
}

func HowAreYou(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I'am fine!"))

	log.Printf("url: %s, elase: %v", r.URL, time.Since(now))

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", HelloWorld)
	mux.HandleFunc("/how", HowAreYou)

	srv := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	fmt.Println("listen at: ", srv.Addr)
	srv.ListenAndServe()
}
