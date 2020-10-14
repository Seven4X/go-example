package convert



type Message struct {
	Id   int    `db:"id"`
	Text string `db:"text"`
}

type Message2 struct {
	Id   int    `json:"id" xml:"id,attr"`
	Text string `json:"text"`
}

