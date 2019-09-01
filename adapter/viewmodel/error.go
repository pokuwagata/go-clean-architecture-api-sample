package viewmodel

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}