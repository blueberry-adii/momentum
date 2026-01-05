package models

import "time"

type Completion struct {
	ID          int
	HabitID     int
	UserID      int
	CompletedAt time.Time
	CreatedAt   time.Time
}
