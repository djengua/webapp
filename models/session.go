package models

import "time"

type Session struct {
	id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}
