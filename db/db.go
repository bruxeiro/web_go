package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Chamando a função para conectar com o banco de dados, conectando com o postgres (existem outros metodos para conexão, ofuscando assim a senha e dando mais segurança para a aplicação)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=9513 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
