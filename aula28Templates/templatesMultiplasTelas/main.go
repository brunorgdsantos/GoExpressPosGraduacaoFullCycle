package main

import (
	"html/template"
	"os"
	_ "os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 30},
	})
	if err != nil {
		panic(err)
	}
	//})
	//http.ListenAndServe(":8003", nil)
}
