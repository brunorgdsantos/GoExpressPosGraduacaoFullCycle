package main

import "net/http"

func main() {
	mux := http.NewServeMux() //Com o Mux vc consegue adicionar varias rotas simultaneas

	http.ListenAndServe(":8080", mux) //Nesse server vai estar a rota da linha 10

	/*mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //Como func(w http.ResponseWriter, r *http.Request) é uma func, voce poderia substiruit pela linha 14 e 16
		w.Write([]byte("hello world"))
	})*/

	mux.HandleFunc("/", HomeHandler) //É o mesmo que fazer a linha 10
	mux.Handle("/blog", &blog{title: "MyBlog"})

	//Criando a mesa rota "/" porém em outro servidor (8001)
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //Mesma rota da linha 14 porem no server 8001
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8001", mux2) //O Mux permite criar varias rotas, repetir rotas, e criar varios servers diferentes.
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type blog struct {
	title string
}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
