package entity

import "time"

type Message struct {
	ID        string
	Message   string
	Grade     int
	Author    *User
	Thread    *Thread
	CreatedAt *time.Time
}
