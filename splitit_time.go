package splitit

import (
	"time"

	"github.com/btubbs/datetime"
)

type SplititTime datetime.DefaultUTC

func (d SplititTime) String() string {
	t := time.Time(d)
	return t.Format(time.RFC3339Nano)
}

func (d *SplititTime) UnmarshalJSON(data []byte) error {
	t, err := datetime.JSONParse(data, time.UTC)
	*d = SplititTime(t)
	return err
}

func (d SplititTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	return t.MarshalJSON()
}

func NewSplititTime(t time.Time) *SplititTime {
	st := SplititTime(t)
	return &st
}
