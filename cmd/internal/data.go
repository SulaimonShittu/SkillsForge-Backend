package data

import (
	"net/mail"
	"time"
)

type Comment struct {
	ID         int64
	TimePosted time.Time
	SenderName string
	Message    string
	Email      mail.Address
}
