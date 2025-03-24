package data

import (
	"net/mail"
	"time"
)

type Comment struct {
	ID         int64        `json:"id"`
	TimePosted time.Time    `json:"-"`
	SenderName string       `json:"senderName"`
	Message    string       `json:"message"`
	Email      mail.Address `json:"email,omitempty"`
}
