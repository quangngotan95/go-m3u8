package m3u8

import "fmt"

// DiscontinuityItem represents a EXT-X-DISCONTINUITY tag to indicate a
// discontinuity between the SegmentItems that proceed and follow it.
type DiscontinuityItem struct{}

func NewDiscontinuityItem() (*DiscontinuityItem, error) {
	return &DiscontinuityItem{}, nil
}

func (di *DiscontinuityItem) String() string {
	return fmt.Sprintf("%s\n", DiscontinuityItemTag)
}
