package m3u8

import (
	"fmt"
	"strings"
	"time"
)

const (
	dateTimeFormat = time.RFC3339Nano
)

// TimeItem represents EXT-X-PROGRAM-DATE-TIME
type TimeItem struct {
	Time time.Time
}

func NewTimeItem(text string) (*TimeItem, error) {
	timeString := strings.Replace(text, TimeItemTag+":", "", -1)

	t, err := ParseTime(timeString)

	if err != nil {
		return nil, err
	}

	return &TimeItem{
		Time: t,
	}, nil

}

func (ti *TimeItem) String() string {
	return fmt.Sprintf("%s:%s", TimeItemTag, ti.Time.Format(dateTimeFormat))
}

func FormatTime(time time.Time) string {
	return time.Format(dateTimeFormat)
}

func ParseTime(value string) (time.Time, error) {
	layouts := []string{
		"2006-01-02T15:04:05.999999999Z0700",
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02T15:04:05.999999999Z07",
	}
	var (
		err error
		t   time.Time
	)
	for _, layout := range layouts {
		if t, err = time.Parse(layout, value); err == nil {
			return t, nil
		}
	}
	return t, err
}
