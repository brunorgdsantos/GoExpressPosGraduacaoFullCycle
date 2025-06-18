package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"GoLang", 40}
	tmp := template.New("CursoTemplate")
	tmp, err := tmp.Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}\n")
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
