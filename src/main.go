package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/About", About)
	mux.HandleFunc("/Login", Login)
	log.Println("Memulai dengan port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BASE_URL ======> hello world,Aku sedang belajar routing di Golang"))

}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ABOUT ========> hello world,Aku sedang belajar routing di Golang"))

}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login ==========> hello world,Aku sedang belajar routing di Golang"))

}
