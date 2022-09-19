package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Superhero struct
type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Superhero) sayHello(from string, message string) string {
	return fmt.Sprintf("%s said : \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero{
			Name:    "Ray",
			Alias:   "MK Gagal",
			Friends: []string{"NGL Inori", "NGL Roxy", "NGL Acan"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":9090", nil)
}
