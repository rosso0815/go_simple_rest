package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start")
	http.HandleFunc("/catalog/products", productsHandler)
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)
	log.Println("done")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("up && running"))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Schuhe, Hose, Hemd"))
}
