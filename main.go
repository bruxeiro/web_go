package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaBanco() *sql.DB {
	conexao := "user=posgres dbname=alura_loja password=9513 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())

	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*"))

func main() {
	db := conectaBanco()
	defer db.Close()

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
