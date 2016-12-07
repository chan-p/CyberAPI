package model

//イベント情報を返すjson構造
type response_json struct {
	Data   []data
}

//イベントの詳細な情報用のjson構造
type data struct {
	Id          int    `json:id`
	Summary     string `json:"summary"`
	Dtstart     string `json:"dtstart"`
	Dtend       string `json:"dtend"`
	Description string `json:"description"`
}

