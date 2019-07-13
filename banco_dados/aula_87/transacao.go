package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@/cursogo")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?, ?)")

	// o correto é tratar cada statement
	stmt.Exec(100, "Bia")
	stmt.Exec(101, "Vanessa")

	_, err = stmt.Exec(100, "Carlos") // chave duplicada
	// o Go levanta o erro somente se for tratado. Se a validação abaixo for comentada, o commit vai ocorrer inserindo o id 100 e 101
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
	}
	tx.Commit()
}
