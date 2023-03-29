package controlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web_go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

//Listar todos os produtos/Index

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

//Função para chamar a página Ne

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", "nil")
}

//Função para inserir novos itens na lista

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//Convertendo preço para float 64 pois o insert retorna uma string, usada a biblioteca strconv
		precoConvert, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		//Convertendo quantidade para int pois o insert retorna uma string, usada a biblioteca strconv
		quantidadeConvert, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}
		models.CriarProduto(nome, descricao, precoConvert, quantidadeConvert)
	}
	http.Redirect(w, r, "/", 301)

}
