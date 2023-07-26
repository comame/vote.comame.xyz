package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/comame/mysql-go"
	"github.com/comame/vote/core"
	"github.com/google/uuid"
)

var (
	ErrNotFound = core.NewUserError(errors.New("not found"))
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

func getTopic(ctx context.Context, topicId string) (*core.Topic, error) {
	db := mysql.Conn()

	var topic core.Topic
	row := db.QueryRowContext(ctx, "SELECT id,name,type FROM topic WHERE id=?", topicId)

	if err := row.Scan(&topic.Id, &topic.Name, &topic.Type); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &topic, nil
}

func getChoiceGeneric(ctx context.Context, topicId string) ([]core.ChoiceGeneric, error) {
	db := mysql.Conn()

	var choices []core.ChoiceGeneric
	rows, err := db.QueryContext(ctx, "SELECT id,topic_id,`order`,`text` FROM choice_generic WHERE topic_id=?", topicId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	for rows.Next() {
		var choice core.ChoiceGeneric
		if err := rows.Scan(&choice.Id, &choice.TopicId, &choice.Order, &choice.Text); err != nil {
			return nil, err
		}
		choices = append(choices, choice)
	}

	return choices, nil
}

func getChoiceCalendar(ctx context.Context, topicId string) ([]core.ChoiceCalendar, error) {
	db := mysql.Conn()

	var choices []core.ChoiceCalendar
	rows, err := db.QueryContext(ctx, "SELECT id,topic_id,`order`,is_all_day,start_datetime,end_datetime FROM choice_calendar WHERE topic_id=?", topicId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	for rows.Next() {
		var choice core.ChoiceCalendar
		if err := rows.Scan(&choice.Id, &choice.TopicId, &choice.Order, &choice.IsAllDay, &choice.StartDateTime, &choice.EndDateTime); err != nil {
			return nil, err
		}
		choices = append(choices, choice)
	}

	return choices, nil
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
	if len(choices) == 0 {
		return nil, nil, core.ErrInvalidFormat
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
	if len(choices) == 0 {
		return nil, nil, core.ErrInvalidFormat
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
			"INSERT INTO choice_calendar "+
				"(id, topic_id, `order`, is_all_day, start_datetime, end_datetime) "+
				"VALUES "+
				"(?, ?, ?, ?, ?, ?) ",
			choiceId, topicId, choices[i].Order, choices[i].IsAllDay, choices[i].StartDateTime, choices[i].EndDateTime); err != nil {
			return nil, nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return &topic, choices, nil
}

func getVote(ctx context.Context, topicId string) ([]core.Vote, error) {
	db := mysql.Conn()

	var votes []core.Vote

	rows, err := db.QueryContext(ctx,
		"SELECT v.id, v.user_name, v.choice_id FROM vote AS v "+
			"LEFT OUTER JOIN choice_generic AS g ON g.id = v.choice_id "+
			"LEFT OUTER JOIN choice_calendar AS c ON c.id = v.choice_id "+
			"WHERE g.topic_id = ? OR c.topic_id = ? ",
		topicId, topicId,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var vote core.Vote
		if err := rows.Scan(&vote.Id, &vote.UserName, &vote.ChoiceId); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}

	return votes, nil
}

func createVote(ctx context.Context, vote core.Vote) (*core.Vote, error) {
	voteId, err := generateUUIDv4()
	if err != nil {
		return nil, err
	}

	vote.Id = voteId

	db := mysql.Conn()

	if _, err := db.ExecContext(
		ctx,
		"INSERT INTO vote (id,user_name,choice_id) VALUES (?, ?, ?)",
		voteId, vote.UserName, vote.ChoiceId,
	); err != nil {
		return nil, err
	}

	return &vote, nil
}
