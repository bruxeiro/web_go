package main

import (
	"net/http"
	"web_go/routes"

	_ "github.com/lib/pq"
)

// Chamando função para carregar rotas e utilizar a porta
func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
