package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// M will be a map for alias
type M map[string]interface{}

func main() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Ray"}
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Ray"}
		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Printf("server started at localhost:9090")
	http.ListenAndServe("localhost:9090", nil)
}
