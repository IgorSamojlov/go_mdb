package types

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	UserId  int    `json:"userId"`
	ChatId  int    `json:"chatId"`
	Message string `json:"message"`
}

type JournalItem struct {
	Message Message   `json:"message"`
	GetAt   time.Time `json:"get_at"`
	UUID    uuid.UUID `json:"uuid"`
}
