package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	Models "./Models"
	secret "./secret"
)

func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	CheckErr(err)
	defer db.Close()

	var topic Models.Topic
	topicRows, err := db.Query(secret.GetQuery("TOPIC"))
	CheckErr(err)
	defer topicRows.Close()

	for topicRows.Next() {
		err := topicRows.Scan(&topic.DOCIDs, &topic.TERMS)
		CheckErr(err)

		log.Println(topic.DOCIDs, topic.TERMS)
	}

	var emotion Models.Emotion
	emotionRows, err := db.Query(secret.GetQuery("EMOTION"))
	CheckErr(err)
	defer emotionRows.Close()

	for emotionRows.Next() {
		err := emotionRows.Scan(&emotion.Seq, &emotion.CmtID, &emotion.Emotion)
		CheckErr(err)

		log.Println(emotion.Seq, emotion.CmtID, emotion.Emotion)
	}

}
