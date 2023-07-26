package main

import (
	"log"
	"net/http"

	"github.com/comame/mysql-go"
	"github.com/comame/router-go"
)

func init() {
	if err := mysql.Initialize(); err != nil {
		panic(err)
	}
}

func main() {
	router.Get("/", getIndex)

	router.Post("/api/topic/new/generic", newTopicGeneric)
	router.Post("/api/topic/new/calendar", newTopicCalendar)

	log.Println("Start http://localhost:8080")
	http.ListenAndServe(":8080", router.Handler())
}
