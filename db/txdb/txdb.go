package main

import (
	"database/sql"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"log"
)

func init() {
	// we register an sql driver named "txdb"
	txdb.Register("txdb", "postgres", "postgres://postgres:linkhub2333@127.0.0.1:5432/link_hub?sslmode=disable")
}

//数据库中原有的能查询到
//事物不会提交，db所有操作不会影响数据库，用于测试非常好👌
func main() {
	// dsn serves as an unique identifier for connection pool
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`INSERT INTO topic(name) VALUES('gopher')`); err != nil {
		log.Fatal(err)
	}

	res, _ := db.Exec(`INSERT INTO topic(name) VALUES('230')`)
	println(res)
	rows, _ := db.Query("select id,name  from topic ")
	for rows.Next() {
		var s string
		var i int
		err = rows.Scan(&i, &s)
		if err != nil {
			log.Fatalln(err)
		}
		print(i)
		println(s)
	}
}
