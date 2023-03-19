package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 213, Quantidade: 1},
		{"Tenis", "Preto", 93, 2},
		{"Fone", "Show de Bola", 120, 3},
	}

	temp.ExecuteTemplate(w, "index", produtos)

}
