package main

import (
	"errors"

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

func createTopicGeneric(topic core.Topic, choices []core.ChoiceGeneric) (*core.Topic, []core.ChoiceGeneric, error) {
	return nil, nil, errors.New("unimplemented")
}
