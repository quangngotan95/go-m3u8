package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// SegmentItem represents EXTINF attributes with the URI that follows,
// optionally allowing an EXT-X-BYTERANGE tag to be set.
type SegmentItem struct {
	Duration        float64
	Segment         string
	Comment         *string
	ProgramDateTime *TimeItem
	ByteRange       *ByteRange
}

func NewSegmentItem(text string) (*SegmentItem, error) {
	var si SegmentItem
	line := strings.Replace(text, SegmentItemTag+":", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	values := strings.Split(line, ",")
	d, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return nil, err
	}

	si.Duration = d
	if len(values) > 1 && values[1] != "" {
		si.Comment = &values[1]
	}

	return &si, nil
}

func (si *SegmentItem) String() string {
	date := ""
	if si.ProgramDateTime != nil {
		date = fmt.Sprintf("%v\n", si.ProgramDateTime)
	}
	byteRange := ""
	if si.ByteRange != nil {
		byteRange = fmt.Sprintf("\n%s:%v", ByteRangeItemTag, si.ByteRange.String())
	}

	comment := ""
	if si.Comment != nil {
		comment = *si.Comment
	}

	return fmt.Sprintf("%s:%v,%s%s\n%s%s", SegmentItemTag, si.Duration, comment, byteRange, date, si.Segment)
}
