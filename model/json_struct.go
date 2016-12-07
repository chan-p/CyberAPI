package model

//イベント情報を返すjson構造
type json_all struct {
	Status bool `json:status`
	Data   []json_event
}

//イベントの詳細な情報用のjson構造
type json_event struct {
	Id          int    `json:id`
	Summary     string `json:"summary"`
	Dtstart     string `json:"dtstart"`
	Dtend       string `json:"dtend"`
	Description string `json:"description"`
}

//タスク情報を返すjson構造
type json struct {
	Status bool `json:status`
	Data   []json_task
}

//タスクの詳細な情報用のjson構造
type json_task struct {
	Id       int    `json:id`
	Title    string `json:"summary"`
	Sub_task string `json:"sub_task"`
	year     int    `json:"year"`
	month    int    `json:"month"`
	day      int    `json:"day"`
}

type json_comp struct {
	Status bool `json:status`
	Data   []json_day
}

type json_day struct {
	Day []json_event
}
