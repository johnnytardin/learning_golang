package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João Gilberto")

	res, _ := stmt.Exec("João Roberto")

	id, _ := res.LastInsertId()
	fmt.Printf("Last id: %d\n", id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Number of lines: %d\n", linhas)

}
