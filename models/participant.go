package models

import "time"

type Participant struct {
	Id        int
	Uuid      string
	Raffle    string
	Active    bool
	CreatedAt time.Time
}
