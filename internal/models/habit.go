package models

import (
	"time"

	"github.com/blueberry-adii/momentum/internal/enums/frequency"
)

type Habit struct {
	ID          int
	UserID      int
	Name        string
	Description string
	Frequency   frequency.Frequency
	CreatedAt   time.Time
}
