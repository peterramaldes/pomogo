package pomo

import "time"

type Pomo struct {
	Start       time.Time `json:"start"`
	Description string    `json:"description"`
}

func NewPomo(start time.Time, description string) Pomo {
	return Pomo{Start: start, Description: description}
}
