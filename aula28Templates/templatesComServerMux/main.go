package main

import (
	"html/template"
	"net/http"
	_ "os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles("content.html"))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 20},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8003", nil)
}
