package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

func main() {
	t := template.Must(template.New("content.html").ParseFiles("content.html"))
	err := t.Execute(os.Stdout, Curso{
		"Go",
		40,
	})
	if err != nil {
		panic(t)
	}
}
