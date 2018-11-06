package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// ByteRange represents sub range of a resource
type ByteRange struct {
	Length *int
	Start  *int
}

func NewByteRange(text string) (*ByteRange, error) {
	if text == "" {
		return nil, nil
	}

	values := strings.Split(text, "@")

	lengthValue, err := strconv.Atoi(values[0])
	if err != nil {
		return nil, err
	}

	br := ByteRange{Length: &lengthValue}

	if len(values) >= 2 {
		startValue, err := strconv.Atoi(values[1])
		if err != nil {
			return &br, err
		}
		br.Start = &startValue
	}

	return &br, nil
}

func (br *ByteRange) String() string {
	if br.Start == nil {
		return fmt.Sprintf("%d", *br.Length)
	}

	return fmt.Sprintf("%d@%d", *br.Length, *br.Start)
}
