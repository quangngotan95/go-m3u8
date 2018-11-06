package m3u8

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

func NewPlaylist() *Playlist {
	return &Playlist{
		Target: 10,
	}
}

func NewPlaylistWithItems(items []Item) *Playlist {
	return &Playlist{
		Target: 10,
		Items:  items,
	}
}

func (pl *Playlist) String() string {
	s, err := Write(pl)
	if err != nil {
		return ""
	}

	return s
}

func (pl *Playlist) AppendItem(item Item) {
	pl.Items = append(pl.Items, item)
}

func (pl *Playlist) IsLive() bool {
	if pl.IsMaster() {
		return false
	}

	return pl.Live
}

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

func (pl *Playlist) PlaylistSize() int {
	result := 0

	for _, item := range pl.Items {
		if _, ok := item.(*PlaylistItem); ok {
			result++
		}
	}

	return result
}

func (pl *Playlist) SegmentSize() int {
	result := 0

	for _, item := range pl.Items {
		if _, ok := (item).(*SegmentItem); ok {
			result++
		}
	}

	return result
}

func (pl *Playlist) ItemSize() int {
	return len(pl.Items)
}

func (pl *Playlist) IsValid() bool {
	return !(pl.PlaylistSize() > 0 && pl.SegmentSize() > 0)
}

func (pl *Playlist) Duration() float64 {
	duration := 0.0

	for _, item := range pl.Items {
		if segmentItem, ok := item.(*SegmentItem); ok {
			duration += segmentItem.Duration
		}
	}

	return duration
}
