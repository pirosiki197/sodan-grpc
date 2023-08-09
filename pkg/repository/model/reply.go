package model

import (
	"time"
)

type Reply struct {
	ID        uint      `json:"id"`
	Text      string    `json:"text"`
	SodanID   uint      `json:"sodan_id"`
	CreaterID string    `json:"creater_id"`
	CreatedAt time.Time `json:"created_at"`
}
