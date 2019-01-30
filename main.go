package main

import (
	"net/http"
	"tasty/ap"
	"tasty/ap/cp"
	"tasty/bp"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ap.AP()
		bp.BP()
		cp.CP()
	})
	http.ListenAndServe(":8080", router)
}
