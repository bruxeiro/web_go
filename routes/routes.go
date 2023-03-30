package routes

import (
	"net/http"
	"web_go/controlers"
)

// Função para carregar todas as rotas do site
func CarregarRotas() {
	http.HandleFunc("/", controlers.Index)
	http.HandleFunc("/new", controlers.New)
	http.HandleFunc("/insert", controlers.Insert)
	http.HandleFunc("/delete", controlers.Delete)
	http.HandleFunc("/edit", controlers.Edit)
}
