package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start")
	http.HandleFunc("/catalog/products", productsHandler)
	http.ListenAndServe(":8080", nil)
	log.Println("done")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Schuhe, Hose, Hemd"))
}
