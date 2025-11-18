package model

import "time"

type Question struct {
	Id int `json:"id"` //gorm:"primaryKey;autoIncrement"`
	Text string `json:"text"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Answer struct {
	Id int `json:"id"` //gorm:"primaryKey;autoIncrement`
	Question_id int `json:"question_id"`
	User_id string `json:"user_id"`
	Text string `json:"text"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type User struct {
	Id int `json:"id"` //gorm:"primaryKey;autoIncrement`
	Name string `json:"name"`
}