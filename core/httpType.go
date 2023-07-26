package core

type TopicGenericWithChoises struct {
	Topic   Topic           `json:"topic"`
	Choices []ChoiceGeneric `json:"choices"`
}

type TopicCalendarWithChoices struct {
	Topic   Topic            `json:"topic"`
	Choices []ChoiceCalendar `json:"choices"`
}

type ResponseGetTopicGeneric = TopicGenericWithChoises

type ResponseGetTopicCalendar = TopicCalendarWithChoices

type RequestCreateTopicGeneric = TopicGenericWithChoises

type ResponseCreateTopicGeneric = response[TopicGenericWithChoises]

type RequestCreateTopicCalendar = TopicCalendarWithChoices

type ResponseCreateTopicCalendar = response[TopicCalendarWithChoices]

type RequestVote = Vote

type ResponseVote = response[struct{}]

type ResponseError = response[struct{}]

type response[T any] struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    T      `json:"body"`
}
