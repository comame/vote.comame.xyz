package main

import (
	"context"

	"github.com/comame/mysql-go"
	"github.com/comame/vote/core"
	"github.com/google/uuid"
)

// MySQL の場合 Primary Key に設定することでパフォーマンスの劣化があるが、どうせ大して使われないだろうという予測と、
// ID を予測不可能にしたいという要望からまあヨシとする。
func generateUUIDv4() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func createTopicGeneric(
	ctx context.Context,
	topic core.Topic,
	choices []core.ChoiceGeneric,
) (*core.Topic, []core.ChoiceGeneric, error) {
	if err := core.AssertTopicGeneric(topic); err != nil {
		return nil, nil, err
	}
	for _, choice := range choices {
		if err := core.AssertChoiceGeneric(choice); err != nil {
			return nil, nil, err
		}
	}

	topicId, err := generateUUIDv4()
	if err != nil {
		return nil, nil, err
	}

	topic.Id = topicId

	db := mysql.Conn()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `
		INSERT INTO topic
		(id, name, type)
		VALUES
		(?, ?, ?)
	`, topicId, topic.Name, topic.Type); err != nil {
		return nil, nil, err
	}

	for i, choice := range choices {
		choiceId, err := generateUUIDv4()
		if err != nil {
			return nil, nil, err
		}

		choices[i].Id = choiceId

		if _, err := tx.ExecContext(ctx,
			"INSERT INTO choice_generic"+
				"(id, topic_id, `order`, `text`)"+
				"VALUES"+
				"(?, ?, ?, ?)",
			choiceId, topicId, choice.Order, choice.Text); err != nil {
			return nil, nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return &topic, choices, nil
}

func createTopicCalendar(
	ctx context.Context,
	topic core.Topic,
	choices []core.ChoiceCalendar,
) (*core.Topic, []core.ChoiceCalendar, error) {
	if err := core.AssertTopicCalendar(topic); err != nil {
		return nil, nil, err
	}
	for _, choice := range choices {
		if err := core.AssertChoiceCalendar(choice); err != nil {
			return nil, nil, err
		}
	}

	topicId, err := generateUUIDv4()
	if err != nil {
		return nil, nil, err
	}

	topic.Id = topicId

	db := mysql.Conn()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `
		INSERT INTO topic
		(id, name, type)
		VALUES
		(?, ?, ?)
	`, topicId, topic.Name, topic.Type); err != nil {
		return nil, nil, err
	}

	for i := range choices {
		choiceId, err := generateUUIDv4()
		if err != nil {
			return nil, nil, err
		}

		choices[i].Id = choiceId

		if choices[i].IsAllDay {
			choices[i].EndDateTime = choices[i].StartDateTime
		}

		if _, err := tx.ExecContext(ctx,
			"INSERT INTO choice_calendar"+
				"(id, topic_id, `order`, is_all_day, start_datetime, end_datetime)"+
				"VALUES"+
				"(?, ?, ?, ?, ?, ?)",
			choiceId, topicId, choices[i].Order, choices[i].IsAllDay, choices[i].StartDateTime, choices[i].EndDateTime); err != nil {
			return nil, nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return &topic, choices, nil
}
