package m3u8

import "fmt"

// MapItem represents a EXT-X-MAP tag which specifies how to obtain the Media
// Initialization Section
type MapItem struct {
	URI       string
	ByteRange *ByteRange
}

func NewMapItem(text string) (*MapItem, error) {
	attributes := ParseAttributes(text)

	br, err := NewByteRange(attributes[ByteRangeTag])
	if err != nil {
		return nil, err
	}

	return &MapItem{
		URI:       attributes[URITag],
		ByteRange: br,
	}, nil
}

func (mi *MapItem) String() string {
	if mi.ByteRange == nil {
		return fmt.Sprintf(`%s:%s="%s"`, MapItemTag, URITag, mi.URI)
	}

	return fmt.Sprintf(`%s:%s="%s",%s="%v"`, MapItemTag, URITag, mi.URI, ByteRangeTag, mi.ByteRange)
}
