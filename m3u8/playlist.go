// Package m3u8 provides utilities for parsing and generating m3u8 playlists
package m3u8

// Item represents an item in a playlist
type Item interface {
	String() string
}

// Playlist represents an m3u8 playlist, it can be a master playlist or a set
// of media segments
type Playlist struct {
	Items                 []Item
	Version               *int
	Cache                 *bool
	Target                int
	Sequence              int
	DiscontinuitySequence *int
	Type                  *string
	IFramesOnly           bool
	IndependentSegments   bool
	Live                  bool
	Master                *bool
}

func (pl *Playlist) String() string {
	s, err := Write(pl)
	if err != nil {
		return ""
	}

	return s
}

// NewPlaylist returns a playlist with default target 10
func NewPlaylist() *Playlist {
	return &Playlist{
		Target: 10,
		Live:   true,
	}
}

// NewPlaylistWithItems returns a playlist with a list of items
func NewPlaylistWithItems(items []Item) *Playlist {
	return &Playlist{
		Target: 10,
		Items:  items,
	}
}

// AppendItem appends an item to the playlist
func (pl *Playlist) AppendItem(item Item) {
	pl.Items = append(pl.Items, item)
}

// IsLive checks if playlist is live (not vod)
func (pl *Playlist) IsLive() bool {
	if pl.IsMaster() {
		return false
	}

	return pl.Live
}

// IsMaster checks if a playlist is a master playlist
func (pl *Playlist) IsMaster() bool {
	if pl.Master != nil {
		return *pl.Master
	}

	plSize := pl.PlaylistSize()
	smSize := pl.SegmentSize()
	if plSize <= 0 && smSize <= 0 {
		return false
	}

	return plSize > 0
}

// PlaylistSize returns number of playlist items in a playlist
func (pl *Playlist) PlaylistSize() int {
	result := 0

	for _, item := range pl.Items {
		if _, ok := item.(*PlaylistItem); ok {
			result++
		}
	}

	return result
}

// PlaylistItems returns list of playlist items in a playlist
func (pl *Playlist) PlaylistItems() []*PlaylistItem {
	var p []*PlaylistItem
	for _, i := range pl.Items {
		if pi, ok := i.(*PlaylistItem); ok {
			p = append(p, pi)
		}
	}
	return p
}

// SegmentSize returns number of segment items in a playlist
func (pl *Playlist) SegmentSize() int {
	result := 0

	for _, item := range pl.Items {
		if _, ok := (item).(*SegmentItem); ok {
			result++
		}
	}

	return result
}

// SegmentItems returns list of segment items in a playlist
func (pl *Playlist) SegmentItems() []*SegmentItem {
	var s []*SegmentItem
	for _, i := range pl.Items {
		if si, ok := i.(*SegmentItem); ok {
			s = append(s, si)
		}
	}
	return s
}

// ItemSize returns number of items in a playlist
func (pl *Playlist) ItemSize() int {
	return len(pl.Items)
}

// IsValid checks if a playlist is valid or not
func (pl *Playlist) IsValid() bool {
	return !(pl.PlaylistSize() > 0 && pl.SegmentSize() > 0)
}

// Duration returns duration of a media playlist
func (pl *Playlist) Duration() float64 {
	duration := 0.0

	for _, item := range pl.Items {
		if segmentItem, ok := item.(*SegmentItem); ok {
			duration += segmentItem.Duration
		}
	}

	return duration
}
