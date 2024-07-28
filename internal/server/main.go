package server

import (
	"fmt"
	"net/http"

	"github.com/inveracity/go-embed-spa/ui"
)

func cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		handler.ServeHTTP(w, r)
	})
}

func Server(port uint) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"hello": "world"}`))
	})

	mux.Handle("/", ui.SvelteKitHandler("/"))
	mux.Handle("/about", ui.SvelteKitHandler("/about"))

	var handler http.Handler = mux

	handler = cors(handler)

	fmt.Printf("api served at http://localhost:%d/api/v1\ngui served at http://localhost:%d/\n", port, port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
