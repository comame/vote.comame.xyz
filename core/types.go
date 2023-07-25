package core

// 議題
type Topic struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Type TopicType `json:"type"`
}

// 自由形式の選択肢
type ChoiceGeneric struct {
	Id      string    `json:"id"`
	TopicId string    `json:"topic_id"`
	Type    TopicType `json:"type"`
	Order   uint      `json:"order"`
	Text    string    `json:"text"`
}

// カレンダー形式の選択肢
type ChoiceCalendar struct {
	Id       string    `json:"id"`
	TopicId  string    `json:"topic_id"`
	Type     TopicType `json:"type"`
	Order    uint      `json:"order"`
	IsAllDay bool      `json:"is_all_day"`
	// 終日の時、ここで日付を表す
	StartDateTime string `json:"start_datetime"`
	// 終日の時、無視する
	EndDateTime string `json:"end_datetime"`
}

type Vote struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	ChoiceId string `json:"choice_id"`
}

type TopicType string

const (
	TopicGeneric  TopicType = "generic"
	TopicCalendar TopicType = "calendar"
)
