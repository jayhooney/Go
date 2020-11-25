package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	secret "./secret"
)

type Topic struct {
	DOCID int
	TERMS string
}

func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	CheckErr(err)
	defer db.Close()

	rows, err := db.Query(secret.GetQuery())
	CheckErr(err)
	defer rows.Close()

	var topic Topic
	for rows.Next() {
		err := rows.Scan(&topic.DOCID, &topic.TERMS)
		CheckErr(err)

		log.Println(topic.DOCID, topic.TERMS)
	}

}
