package models

import "time"

type Completions struct {
	ID          int
	HabitID     int
	UserID      int
	CompletedAt time.Time
	CreatedAt   time.Time
}
