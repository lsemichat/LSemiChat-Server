package entity

import "time"

type User struct {
	ID          string
	UserID      string
	Name        string
	Mail        string
	Image       string
	Profile     string
	IsAdmin     int
	LoginAt     *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Password    string
	Tags        []*Tag
	Evaluations []*EvaluationScore
}
