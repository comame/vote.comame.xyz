package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/comame/vote/core"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}

func newTopicGeneric(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		responseError(w, err)
		return
	}

	var req core.RequestCreateTopicGeneric
	if err := json.Unmarshal(b, &req); err != nil {
		responseError(w, err)
		return
	}

	t, c, err := createTopicGeneric(r.Context(), req.Topic, req.Choices)
	if err != nil {
		responseError(w, err)
		return
	}

	res := core.ResponseCreateTopicGeneric{
		Body: core.TopicGenericWithChoises{
			Topic:   *t,
			Choices: c,
		},
	}
	resopnseBody(w, res)
}

func newTopicCalendar(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		responseError(w, err)
		return
	}

	var req core.RequestCreateTopicCalendar
	if err := json.Unmarshal(b, &req); err != nil {
		responseError(w, err)
		return
	}

	t, c, err := createTopicCalendar(r.Context(), req.Topic, req.Choices)
	if err != nil {
		responseError(w, err)
		return
	}

	res := core.ResponseCreateTopicCalendar{
		Body: core.TopicCalendarWithChoices{
			Topic:   *t,
			Choices: c,
		},
	}
	resopnseBody(w, res)
}

func responseError(w http.ResponseWriter, err error) {
	body := core.ResponseError{
		Error:   true,
		Message: "error",
		Body:    struct{}{},
	}

	userErr, ok := err.(*core.UserError)
	if ok {
		log.Println("[EXPOSED] " + err.Error())
		body.Message = userErr.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	b, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":true,"message":"error","body":{}}`)
		return
	}

	w.Write(b)
}

func resopnseBody(w http.ResponseWriter, body any) {
	b, err := json.Marshal(body)
	if err != nil {
		responseError(w, err)
		return
	}

	w.Write(b)
}
