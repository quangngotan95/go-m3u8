package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// Resolution represents a resolution for a playlist item, e.g: 1920x1080
type Resolution struct {
	Width  int
	Height int
}

func (r *Resolution) String() string {
	if r == nil {
		return ""
	}

	return fmt.Sprintf("%dx%d", r.Width, r.Height)
}

// NewResolution parses a string and returns a *Resolution
func NewResolution(text string) (*Resolution, error) {
	values := strings.Split(text, "x")
	if len(values) <= 1 {
		return nil, ErrResolutionInvalid
	}

	width, err := strconv.ParseInt(values[0], 0, 0)
	if err != nil {
		return nil, err
	}

	height, err := strconv.ParseInt(values[1], 0, 0)
	if err != nil {
		return nil, err
	}

	return &Resolution{
		Width:  int(width),
		Height: int(height),
	}, nil
}
