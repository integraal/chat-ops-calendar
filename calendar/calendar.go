package calendar

import "time"

func (e *Event) GetStartTime() time.Time {
	return time.Unix(e.Start, 0)
}

func (e *Event) GetEndTime() time.Time {
	return time.Unix(e.End, 0)
}

func (e *Event) GetDuration() time.Duration {
	return e.GetEndTime().Sub(e.GetStartTime())
}
