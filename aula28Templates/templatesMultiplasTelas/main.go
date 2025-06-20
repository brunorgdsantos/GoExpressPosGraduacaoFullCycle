package main

import (
	"html/template"
	"os"
	_ "os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	//t = template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"go", 40},
		{"java", 20},
		{"python", 30},
	})
	if err != nil {
		panic(err)
	}
}
