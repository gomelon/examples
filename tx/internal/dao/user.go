package dao

import "time"

type User struct {
	Id        int64
	Name      string
	Gender    int8
	Birthday  time.Time
	CreatedAt time.Time
}
