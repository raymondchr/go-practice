package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// M as map
type M map[string]interface{}

func main() {
	//HTML
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

	//Static Handler//
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	//Static Handler//

	//Server//
	var address = "localhost:9090"
	fmt.Printf("server started at %s\n", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	//Server//
}
