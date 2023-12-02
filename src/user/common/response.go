package common

type Message struct {
	Msg string `json:"message"`
}

type Error struct {
	Msg  string `json:"message"`
	Errs any    `json:"errors"`
}
