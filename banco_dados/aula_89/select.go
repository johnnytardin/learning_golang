package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:secret@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 5)
	defer rows.Close()

	slUsers := []usuario{}
	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		slUsers = append(slUsers, u)
		fmt.Println(u)
	}
	fmt.Println("O slice de struct de usuario ficou:")
	fmt.Println(slUsers)

}
