package types

type Message struct {
	UserId  int    `json:"userId"`
	ChatId  int    `json:"chatId"`
	Message string `json:"message"`
}
