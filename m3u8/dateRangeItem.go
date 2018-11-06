package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// DateRangeItem represents a #EXT-X-DATERANGE tag
type DateRangeItem struct {
	ID               string
	Class            *string
	StartDate        string
	EndDate          *string
	Duration         *float64
	PlannedDuration  *float64
	Scte35Cmd        *string
	Scte35Out        *string
	Scte35In         *string
	EndOnNext        bool
	ClientAttributes map[string]string
}

func NewDateRangeItem(text string) (*DateRangeItem, error) {
	attributes := ParseAttributes(text)
	duration, err := parseFloat(attributes, DurationTag)
	if err != nil {
		return nil, err
	}
	plannedDuartion, err := parseFloat(attributes, PlannedDurationTag)
	if err != nil {
		return nil, err
	}

	return &DateRangeItem{
		ID:               attributes[IDTag],
		Class:            pointerTo(attributes, ClassTag),
		StartDate:        attributes[StartDateTag],
		EndDate:          pointerTo(attributes, EndDateTag),
		Duration:         duration,
		PlannedDuration:  plannedDuartion,
		Scte35Cmd:        pointerTo(attributes, Scte35CmdTag),
		Scte35Out:        pointerTo(attributes, Scte35OutTag),
		Scte35In:         pointerTo(attributes, Scte35InTag),
		EndOnNext:        attributeExists(EndOnNextTag, attributes),
		ClientAttributes: parseClientAttributes(attributes),
	}, nil
}

func (dri *DateRangeItem) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf(quotedFormatString, IDTag, dri.ID))
	if dri.Class != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, ClassTag, *dri.Class))
	}
	slice = append(slice, fmt.Sprintf(quotedFormatString, StartDateTag, dri.StartDate))
	if dri.EndDate != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, EndDateTag, *dri.EndDate))
	}
	if dri.Duration != nil {
		slice = append(slice, fmt.Sprintf(formatString, DurationTag, *dri.Duration))
	}
	if dri.PlannedDuration != nil {
		slice = append(slice, fmt.Sprintf(formatString, PlannedDurationTag, *dri.PlannedDuration))
	}
	clientAttributes := formatClientAttributes(dri.ClientAttributes)
	slice = append(slice, clientAttributes...)

	if dri.Scte35Cmd != nil {
		slice = append(slice, fmt.Sprintf(formatString, Scte35CmdTag, *dri.Scte35Cmd))
	}
	if dri.Scte35Out != nil {
		slice = append(slice, fmt.Sprintf(formatString, Scte35OutTag, *dri.Scte35Out))
	}
	if dri.Scte35In != nil {
		slice = append(slice, fmt.Sprintf(formatString, Scte35InTag, *dri.Scte35In))
	}
	if dri.EndOnNext {
		slice = append(slice, fmt.Sprintf(`%s=YES`, EndOnNextTag))
	}

	return fmt.Sprintf("%s:%s", DateRangeItemTag, strings.Join(slice, ","))
}

func parseClientAttributes(attributes map[string]string) map[string]string {
	result := make(map[string]string)
	hasCA := false

	for key, value := range attributes {
		if strings.HasPrefix(key, "X-") {
			result[key] = value
			hasCA = true
		}
	}

	if hasCA {
		return result
	}

	return nil
}

func formatClientAttributes(ca map[string]string) []string {
	if ca == nil {
		return nil
	}

	var slice []string

	for key, value := range ca {
		formatString := `%s=%s`
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			formatString = `%s="%s"`
		}
		slice = append(slice, fmt.Sprintf(formatString, key, value))
	}

	return slice
}
