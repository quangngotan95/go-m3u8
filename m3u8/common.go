package m3u8

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	quotedFormatString    = `%s="%v"`
	formatString          = `%s=%v`
	frameRateFormatString = `%s=%.3f`
)

var (
	parseRegex = regexp.MustCompile(`([A-z0-9-]+)\s*=\s*("[^"]*"|[^,]*)`)
)

func ParseAttributes(text string) map[string]string {
	res := make(map[string]string)
	value := strings.Replace(text, "\n", "", -1)
	matches := parseRegex.FindAllStringSubmatch(value, -1)
	for _, match := range matches {
		if len(match) >= 3 {
			key := match[1]
			value := strings.Replace(match[2], `"`, "", -1)
			res[key] = value
		}
	}

	return res
}

func parseFloat(attributes map[string]string, key string) (*float64, error) {
	stringValue, ok := attributes[key]
	if !ok {
		return nil, nil
	}

	value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func parseInt(attributes map[string]string, key string) (*int, error) {
	stringValue, ok := attributes[key]
	if !ok {
		return nil, nil
	}

	int64Value, err := strconv.ParseInt(stringValue, 0, 0)
	if err != nil {
		return nil, err
	}

	value := int(int64Value)

	return &value, nil
}

func parseYesNo(attributes map[string]string, key string) *bool {
	stringValue, ok := attributes[key]

	if !ok {
		return nil
	}

	val := false

	if stringValue == YesValue {
		val = true
	}

	return &val
}

func formatYesNo(value bool) string {
	if value {
		return YesValue
	}

	return NoValue
}

func attributeExists(key string, attributes map[string]string) bool {
	_, ok := attributes[key]
	return ok
}

func pointerTo(attributes map[string]string, key string) *string {
	value, ok := attributes[key]

	if !ok {
		return nil
	}

	return &value
}
