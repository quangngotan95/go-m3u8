package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// PlaybackStart represents a #EXT-X-START tag and attributes
type PlaybackStart struct {
	TimeOffset float64
	Precise    *bool
}

func NewPlaybackStart(text string) (*PlaybackStart, error) {
	attributes := ParseAttributes(text)

	timeOffset, err := strconv.ParseFloat(attributes[TimeOffsetTag], 64)
	if err != nil {
		return nil, err
	}

	return &PlaybackStart{
		TimeOffset: timeOffset,
		Precise:    parseYesNo(attributes, PreciseTag),
	}, nil
}

func (ps *PlaybackStart) String() string {
	slice := []string{fmt.Sprintf(formatString, TimeOffsetTag, ps.TimeOffset)}
	if ps.Precise != nil {
		slice = append(slice, fmt.Sprintf(formatString, PreciseTag, formatYesNo(*ps.Precise)))
	}

	return fmt.Sprintf(`%s:%s`, PlaybackStartTag, strings.Join(slice, ","))
}
