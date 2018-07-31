package models

import (
	"time"
)

type NewsPost struct {
	ID     int
	Title  string
	Body   string
	Date   time.Time
	Author string
}
