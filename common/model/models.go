package model

import "time"

type Question struct {
	id int
	text string
	created_at time.Time
}

type Answer struct {
	id int
	question_id int
	user_id string
	text string
	created_at time.Time
}

type User struct {
	id int
	Name string
}