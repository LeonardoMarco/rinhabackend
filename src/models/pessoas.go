package models

import (
	"rinhabackendleo/src/config"

	"github.com/lib/pq"
)

type Pessoas struct {
	Id         string         `json:"id"`
	Apelido    string         `json:"apelido"`
	Nome       string         `json:"nome"`
	Nascimento string         `json:"nascimento"`
	Stack      pq.StringArray `json:"stack"`
}

func CreatePessoa(req Pessoas) (Pessoas, error) {
	db := config.ConnectDatabase()
	var pessoa Pessoas

	query := `
		INSERT INTO pessoas
		(
			apelido, nome, nascimento, stack
		)
		VALUES
			(?,?,?,?)
		RETURNING *
	`

	err := db.Raw(query, req.Apelido, req.Nome, req.Nascimento, req.Stack).Scan(&pessoa).Error

	return pessoa, err
}

func GetPessoas(term string) ([]Pessoas, error) {
	var pessoas []Pessoas

	db := config.ConnectDatabase()

	query := `
		SELECT * 
		FROM 
			pessoas 
		WHERE
			upper(apelido) LIKE upper('%' || $1 || '%')
			OR upper(nome) LIKE upper('%' || $1 || '%')
			OR upper(stack::text) LIKE upper('%' || $1 || '%')
		LIMIT 50
	`

	err := db.Raw(query, term).Scan(&pessoas).Error

	return pessoas, err
}

func GetPessoasById(id string) (Pessoas, error) {
	var pessoa Pessoas

	db := config.ConnectDatabase()

	err := db.Raw("SELECT * FROM pessoas WHERE id = ?", id).Scan(&pessoa).Error

	return pessoa, err
}

func GetPessoasCount() (int, error) {
	var count int

	db := config.ConnectDatabase()

	err := db.Raw("SELECT COUNT(*) FROM pessoas").Row().Scan(&count)

	return count, err
}
