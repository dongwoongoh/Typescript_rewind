package model

import (
	"time"
)

type LoanRecord struct {
	BaseModel
	UserID     uint
	BookID     uint
	LoanedAt   time.Time
	ReturnedAt time.Time
}
