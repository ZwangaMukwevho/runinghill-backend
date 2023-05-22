package model

type Sentence struct {
	ID       string `json:"id"`
	Sentence string `json:"sentence"`
	UserId   string `json:"userid"`
}
