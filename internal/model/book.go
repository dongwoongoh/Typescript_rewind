package model

import (
	"time"
)

type Book struct {
	BaseModel
	Title       string
	Author      string
	PublishedAt time.Time
	ISBN        string
}
