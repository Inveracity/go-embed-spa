package main

import (
	"log"
	"net/http"

	embed "github.com/inveracity/go-embed-spa"
)

func main() {
	http.HandleFunc("/", handleSPA)
	log.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSPA(w http.ResponseWriter, r *http.Request) {
	http.FileServer(embed.UI()).ServeHTTP(w, r)
}
