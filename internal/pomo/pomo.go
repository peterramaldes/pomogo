package pomo

import "time"

type Day string

func NewDay() Day {
	return Day(time.Now().Format(time.DateOnly))
}

type TrackingFile struct {
	CurrentPomo Pomo             `json:"current_pomo"`
	Results     map[Day][]Result `json:"results"`
}

func (t *TrackingFile) StoreCurrentPomo() {
	p := t.CurrentPomo
	if (Pomo{}) == p {
		return
	}

	result := NewResult(p)
	day := NewDay()

	t.Results[day] = append(t.Results[day], result)
}

type Result struct {
	TimeElapsed time.Duration `json:"time_elapsed"`
	Start       time.Time     `json:"start"`
	End         time.Time     `json:"end"`
	Logged      bool          `json:"logged"`
}

func NewResult(pomo Pomo) Result {
	end := time.Now()
	started := pomo.Start

	return Result{
		TimeElapsed: end.Sub(started),
		Start:       started,
		End:         end,
		Logged:      false,
	}
}

type Pomo struct {
	Start    time.Time `json:"start"`
	Activity string    `json:"activity"`
}

func NewPomo(start time.Time, activity string) Pomo {
	return Pomo{
		Start:    start,
		Activity: activity,
	}
}
