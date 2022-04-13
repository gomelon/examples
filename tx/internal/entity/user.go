package entity

import "time"

type User struct {
	ID        int64
	Name      string
	Gender    int8
	Birthday  time.Time
	CreatedAt time.Time
}
