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
	router.Get("/api/topic/get/:id", handleGetTopic)

	router.Get("/api/vote/get/:id", handleGetVote)
	router.Post("/api/vote/new", handleCreateVote)

	router.Post("/api/topic/new/generic", handleNewTopicGeneric)
	router.Post("/api/topic/new/calendar", handleNewTopicCalendar)

	router.Get("/*", handleGetStatic)

	log.Println("Start http://localhost:8080")
	http.ListenAndServe(":8080", router.Handler())
}
