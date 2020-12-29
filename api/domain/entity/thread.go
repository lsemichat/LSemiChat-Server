package entity

import "time"

type Thread struct {
	ID          string
	Name        string
	Description string
	LimitUsers  int
	Author      *User
	IsPublic    int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Tags        []*Tag
}
