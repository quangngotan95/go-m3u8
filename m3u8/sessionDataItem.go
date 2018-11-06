package m3u8

import (
	"fmt"
	"strings"
)

// SessionDataItem represents a set of EXT-X-SESSION-DATA attributes
type SessionDataItem struct {
	DataID   string
	Value    *string
	URI      *string
	Language *string
}

func NewSessionDataItem(text string) (*SessionDataItem, error) {
	attributes := ParseAttributes(text)

	return &SessionDataItem{
		DataID:   attributes[DataIDTag],
		Value:    pointerTo(attributes, ValueTag),
		URI:      pointerTo(attributes, URITag),
		Language: pointerTo(attributes, LanguageTag),
	}, nil
}

func (sdi *SessionDataItem) String() string {
	slice := []string{fmt.Sprintf(quotedFormatString, DataIDTag, sdi.DataID)}

	if sdi.Value != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, ValueTag, *sdi.Value))
	}
	if sdi.URI != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, URITag, *sdi.URI))
	}
	if sdi.Language != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, LanguageTag, *sdi.Language))
	}

	return fmt.Sprintf(`%s:%s`, SessionDataItemTag, strings.Join(slice, ","))
}
