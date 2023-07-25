package main

import (
	"log"
	"net/http"

	"github.com/comame/router-go"
)

func main() {
	router.Get("/", getIndex)

	router.Post("/api/topic/new/generic", newTopicGeneric)

	log.Println("Start http://localhost:8080")
	http.ListenAndServe(":8080", router.Handler())
}
