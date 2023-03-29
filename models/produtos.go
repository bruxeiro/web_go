package models

import (
	"web_go/db"

	_ "github.com/lib/pq"
)

//Definido o tipo/modelo produto com seus dados principais

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//Função para conexão com o banco de dados

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	//Criando um slice para o produto, adicionando os itens com id, nome, descrição, preço e quantidade.
	//Por fim, adicionando-os no slice
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

//Função para criar produto, se conecta com o banco insere os produtos no Banco

func CriarProduto(nome, descricao string, precoConvert float64, quantidadeConvert int) {
	db := db.ConectaComBancoDeDados()

	insereDadosBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosBanco.Exec(nome, descricao, precoConvert, quantidadeConvert)
	defer db.Close()
}
