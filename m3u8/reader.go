package m3u8

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type state struct {
	open        bool
	currentItem Item
	master      bool
}

func ReadString(text string) (*Playlist, error) {
	return Read(strings.NewReader(text))
}

func ReadFile(path string) (*Playlist, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Read(bytes.NewReader(f))
}

func Read(reader io.Reader) (*Playlist, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}

	pl := NewPlaylist()
	st := &state{}
	eof := false
	header := true

	for !eof {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			eof = true
		} else if err != nil {
			return nil, err
		}

		value := strings.TrimSpace(line)
		if header && value != HeaderTag {
			return nil, ErrPlaylistInvalid
		}

		if err := parseLine(value, pl, st); err != nil {
			return nil, err
		}

		header = false
	}

	return pl, nil
}

func parseLine(line string, pl *Playlist, st *state) error {
	var err error
	switch {
	// basic tags
	case matchTag(line, VersionTag):
		pl.Version, err = parseIntPtr(line, VersionTag)
	// media segment tags
	case matchTag(line, SegmentItemTag):
		st.currentItem, err = NewSegmentItem(line)
		st.master = false
		st.open = true
	case matchTag(line, DiscontinuityItemTag):
		st.master = false
		st.open = false
		item, err := NewDiscontinuityItem()
		if err != nil {
			return parseError(line, err)
		}
		pl.Items = append(pl.Items, item)
	case matchTag(line, ByteRangeItemTag):
		value := strings.Replace(line, ByteRangeItemTag+":", "", -1)
		value = strings.Replace(value, "\n", "", -1)
		br, err := NewByteRange(value)
		if err != nil {
			return parseError(line, err)
		}
		mit, ok := st.currentItem.(*MapItem)
		if ok {
			mit.ByteRange = br
			st.currentItem = mit
		} else {
			sit, ok := st.currentItem.(*SegmentItem)
			if ok {
				sit.ByteRange = br
				st.currentItem = sit
			}
		}

	case matchTag(line, KeyItemTag):
		item, err := NewKeyItem(line)
		if err != nil {
			return parseError(line, err)
		}
		pl.Items = append(pl.Items, item)
	case matchTag(line, MapItemTag):
		item, err := NewMapItem(line)
		if err != nil {
			return parseError(line, err)
		}
		pl.Items = append(pl.Items, item)
	case matchTag(line, TimeItemTag):
		pdt, err := NewTimeItem(line)
		if err != nil {
			return parseError(line, err)
		}
		if st.open {
			if item, ok := st.currentItem.(*SegmentItem); !ok {
				return parseError(line, ErrSegmentItemInvalid)
			} else {
				item.ProgramDateTime = pdt
			}
		} else {
			pl.Items = append(pl.Items, pdt)
		}
	case matchTag(line, DateRangeItemTag):
		dri, err := NewDateRangeItem(line)
		if err != nil {
			return parseError(line, err)
		}
		pl.Items = append(pl.Items, dri)

	// media playlist tags
	case matchTag(line, MediaSequenceTag):
		pl.Sequence, err = parseIntValue(line, MediaSequenceTag)
	case matchTag(line, DiscontinuitySequenceTag):
		pl.DiscontinuitySequence, err = parseIntPtr(line, DiscontinuitySequenceTag)
	case matchTag(line, CacheTag):
		ptr := parseYesNoPtr(line, CacheTag)
		pl.Cache = ptr
	case matchTag(line, TargetDurationTag):
		pl.Target, err = parseIntValue(line, TargetDurationTag)
	case matchTag(line, IFramesOnlyTag):
		pl.IFramesOnly = true
	case matchTag(line, PlaylistTypeTag):
		pl.Type = parseStringPtr(line, PlaylistTypeTag)

	// master playlist tags
	case matchTag(line, MediaItemTag):
		st.open = false
		mi, err := NewMediaItem(line)
		if err != nil {
			return parseError(line, err)
		}
		pl.Items = append(pl.Items, mi)
	case matchTag(line, SessionDataItemTag):
		sdi, err := NewSessionDataItem(line)
		if err != nil {
			return parseError(line, err)
		}
		pl.Items = append(pl.Items, sdi)
	case matchTag(line, SessionKeyItemTag):
		ski, err := NewSessionKeyItem(line)
		if err != nil {
			return parseError(line, err)
		}
		pl.Items = append(pl.Items, ski)
	case matchTag(line, PlaylistItemTag):
		st.master = true
		st.open = true
		pi, err := NewPlaylistItem(line)
		if err != nil {
			return parseError(line, err)
		}
		st.currentItem = pi
	case matchTag(line, PlaylistIframeTag):
		st.master = true
		st.open = false
		pi, err := NewPlaylistItem(line)
		if err != nil {
			return parseError(line, err)
		}
		pi.IFrame = true
		pl.Items = append(pl.Items, pi)
		st.currentItem = pi

	// universal tags
	case matchTag(line, PlaybackStartTag):
		ps, err := NewPlaybackStart(line)
		if err != nil {
			return parseError(line, err)
		}
		pl.Items = append(pl.Items, ps)
	case matchTag(line, IndependentSegmentsTag):
		pl.IndependentSegments = true
	default:
		if st.currentItem != nil && st.open {
			return parseNextLine(line, pl, st)
		}
	}

	return parseError(line, err)
}

func parseNextLine(line string, pl *Playlist, st *state) error {
	value := strings.Replace(line, "\n", "", -1)
	value = strings.Replace(value, "\r", "", -1)
	if st.master {
		// PlaylistItem
		it, ok := st.currentItem.(*PlaylistItem)
		if !ok {
			return parseError(line, ErrPlaylistItemInvalid)
		}
		it.URI = value
		pl.Items = append(pl.Items, it)
	} else {
		// SegmentItem
		it, ok := st.currentItem.(*SegmentItem)
		if !ok {
			return parseError(line, ErrSegmentItemInvalid)
		}
		it.Segment = value
		pl.Items = append(pl.Items, it)
	}

	st.open = false

	return nil
}

func matchTag(line, tag string) bool {
	return strings.HasPrefix(line, tag) && !strings.HasPrefix(line, tag+"-")
}

func parseIntValue(line string, tag string) (int, error) {
	var v int
	_, err := fmt.Sscanf(line, tag+":%d", &v)
	return v, err
}

func parseIntPtr(line string, tag string) (*int, error) {
	var ptr int
	_, err := fmt.Sscanf(line, tag+":%d", &ptr)
	return &ptr, err
}

func parseStringPtr(line string, tag string) *string {
	value := strings.Replace(line, tag+":", "", -1)
	if value == "" {
		return nil
	}
	return &value
}

func parseYesNoPtr(line string, tag string) *bool {
	value := strings.Replace(line, tag+":", "", -1)
	var b bool
	if value == YesValue {
		b = true
	} else {
		b = false
	}

	return &b
}

func parseError(line string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("error: %v when parsing playlist error for line: %s", err, line)
}
