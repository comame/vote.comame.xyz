package main

import (
	"encoding/json"
	"errors"
	"io"
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

	if req.Topic.Type != core.TopicGeneric {
		responseError(w, errors.New("invalid topic type"))
		return
	}

	t, c, err := createTopicGeneric(req.Topic, req.Choices)
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

func responseError(w http.ResponseWriter, err error) {
	body := core.ResponseError{
		Error:   true,
		Message: err.Error(),
		Body:    struct{}{},
	}

	b, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":true,"message":"internal server error","body":{}}`)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
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
