package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("Post"))
		case "GET":
			w.Write([]byte("Get"))
		default:
			http.Error(w, "", http.StatusBadRequest)

		}
	})

	fmt.Printf("Server started at :9000")
	http.ListenAndServe(":9000", nil)
}
