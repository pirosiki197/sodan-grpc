package model

import (
	"time"
)

type Sodan struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreaterID string    `json:"creater_id"`
	CreatedAt time.Time `json:"created_at"`
	IsClosed  bool      `json:"is_closed"`
	Tags      []*Tag    `json:"tags"`
}
