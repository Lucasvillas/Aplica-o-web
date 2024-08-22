package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/loja_limpeza")
	if err != nil {
		panic(err.Error())
	}
	return db
}
