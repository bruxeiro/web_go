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

		p.Id = id
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

// Função para deletar um produto, primeiramente se conecta com o banco e retira o produto, lembrar de fechar conexão
func DeleteProduto(id string) {
	//Inicia conexão com o banco de dados usando a função
	db := db.ConectaComBancoDeDados()
	//Criando o input para deletar o banco de dados
	deleteDadosBanco, err := db.Prepare("delete from produtos where id=$1")
	//Tratando um possivel erro e exibindo-o na tela
	if err != nil {
		panic(err.Error())
	}
	//Executando a deleção do item, atribuindo o id ao item a ser deletado
	deleteDadosBanco.Exec(id)
	//Fechando o banco de dados, boas praticas
	db.Close()
}

func EditaProduto(id string) Produto {

	db := db.ConectaComBancoDeDados()
	produtoBanco, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar

}

func AtualizaProduto(nome, descricao string, preco float64, quantidade int, id int) {
	db := db.ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
