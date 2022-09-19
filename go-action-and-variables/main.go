package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Info structure for person
type Info struct {
	Affiliation string
	Address     string
}

// Person with Info struct embedded
type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

// GetAffInfo is affiliation from Info Struct
func (t Info) GetAffInfo() string {
	return "(BURU CARI GOBLOK!)"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Mon",
			Gender:  "Male",
			Hobbies: []string{"Game", "BBall"},
			Info:    Info{"Belom Kerja", "TAR DULU!!"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Printf("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
