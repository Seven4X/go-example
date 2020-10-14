package query

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//使用queryContext也能保存数据
func InsertByQueryContext() {
	connStr := "postgres://postgres:linkhub2333@127.0.0.1/link_hub?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	row, err := db.QueryContext(ctx, `INSERT INTO "topic" ("name","tags","create_by","score","agree","disagree") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`, "ttest1", "", "", 0, 0, 0)
	if err != nil {
		println(err.Error())
	}
	println(row)
}
