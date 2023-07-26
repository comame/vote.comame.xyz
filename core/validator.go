package core

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidFormat     = NewUserError(errors.New("invalid format"))
	ErrInvalidTopicType  = NewUserError(errors.New("invalid topic type"))
	ErrInvalidChoiceType = NewUserError(errors.New("invalid choice type"))
	ErrInvalidTimeRange  = NewUserError(errors.New("invalid time range"))
)

func AssertTopicGeneric(v Topic) error {
	if v.Name == "" {
		return ErrInvalidFormat
	}

	if v.Type != TopicGeneric {
		return ErrInvalidTopicType
	}

	return nil
}

func AssertTopicCalendar(v Topic) error {
	if v.Name == "" {
		return ErrInvalidFormat
	}

	if v.Type != TopicCalendar {
		return ErrInvalidTopicType
	}

	return nil
}

func AssertChoiceGeneric(v ChoiceGeneric) error {
	if v.Type != TopicGeneric {
		return ErrInvalidChoiceType
	}

	return nil
}

func AssertChoiceCalendar(v ChoiceCalendar) error {
	if v.Type != TopicCalendar {
		return ErrInvalidChoiceType
	}

	const layout = "2006-01-02 15:04:05"

	start, err := time.Parse(layout, v.StartDateTime)
	if err != nil {
		return err
	}

	if !v.IsAllDay {
		end, err := time.Parse(layout, v.EndDateTime)
		if err != nil {
			return err
		}

		if start.After(end) {
			return ErrInvalidTimeRange
		}
	}

	return nil
}

func AssertVote(v Vote) error {
	if strings.TrimSpace(v.UserName) == "" {
		return ErrInvalidFormat
	}

	return nil
}
