package pomo

import "time"

type Pomo struct {
	Start       time.Time `json:"start"`
	Description string    `json:"description"`
}
