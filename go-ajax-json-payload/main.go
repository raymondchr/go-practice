package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Name   string `json:"name"`
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		byteArray, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		message := fmt.Sprintf(
			"hello, my name is %s. I'm %d year old %s",
			payload.Name,
			payload.Age,
			payload.Gender,
		)

		w.Write([]byte(message))
		w.Write([]byte(byteArray))

	}

	http.Error(w, "Only accept POST request", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/save", handleSave)

	http.Handle("/static/",
		http.StripPrefix("/static",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("Server started at localhost:9090")
	http.ListenAndServe(":9090", nil)
}
