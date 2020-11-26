package mydb

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

func DBProcess() []Models.Topic {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	CheckErr(err)
	defer db.Close()

	// 쿼리 추출 결과를 구조체를 선언하든, 변수로 각각 선언해서 그 값을 받을 때
	// 자료형이 맞지 않아도, 필드명이 달라도 정상적으로 처리된다.
	// 결국 구조체 안에 선언된 필드의 개수만 맞으면 처리되는듯. (단, 문자열을 int로 받으면 에러발생.)
	var topicSlice []Models.Topic
	topicRows, err := db.Query(secret.GetQuery("TOPIC"))
	CheckErr(err)
	defer topicRows.Close()

	for topicRows.Next() {
		var topic Models.Topic
		err := topicRows.Scan(&topic.DOCID, &topic.TERMS)
		CheckErr(err)

		topicSlice = append(topicSlice, topic)
	}

	for _, topic := range topicSlice {
		log.Println(topic.DOCID, topic.TERMS)
	}

	// var emotionSlice []Models.Emotion
	// emotionRows, err := db.Query(secret.GetQuery("EMOTION"))
	// CheckErr(err)
	// defer emotionRows.Close()

	// for emotionRows.Next() {
	// 	var emotion Models.Emotion
	// 	err := emotionRows.Scan(&emotion.Seq, &emotion.CmtID, &emotion.Emotion)
	// 	CheckErr(err)

	// 	emotionSlice = append(emotionSlice, emotion)
	// }

	// for _, emotion := range emotionSlice {
	// 	log.Println(emotion.Seq, emotion.CmtID, emotion.Emotion)
	// }

	return topicSlice

}
