package test

import (
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestSegmentItem_Parse(t *testing.T) {
	time, err := m3u8.ParseTime("2010-02-19T14:54:23Z")
	assert.Nil(t, err)

	item := &m3u8.SegmentItem{
		Duration: 10.991,
		Segment:  "test.ts",
		ProgramDateTime: &m3u8.TimeItem{
			Time: time,
		},
	}

	assert.Equal(t, "#EXTINF:10.991,\n#EXT-X-PROGRAM-DATE-TIME:2010-02-19T14:54:23Z\ntest.ts", item.String())

	item = &m3u8.SegmentItem{
		Duration: 10.991,
		Segment:  "test.ts",
		Comment:  pointer.ToString("anything"),
	}

	assert.Equal(t, "#EXTINF:10.991,anything\ntest.ts", item.String())

	item = &m3u8.SegmentItem{
		Duration: 10.991,
		Segment:  "test.ts",
		Comment:  pointer.ToString("anything"),
		ByteRange: &m3u8.ByteRange{
			Length: pointer.ToInt(4500),
			Start:  pointer.ToInt(600),
		},
	}

	assert.Equal(t, "#EXTINF:10.991,anything\n#EXT-X-BYTERANGE:4500@600\ntest.ts", item.String())

	item = &m3u8.SegmentItem{
		Duration: 10.991,
		Segment:  "test.ts",
		Comment:  pointer.ToString("anything"),
		ByteRange: &m3u8.ByteRange{
			Length: pointer.ToInt(4500),
		},
	}

	assert.Equal(t, "#EXTINF:10.991,anything\n#EXT-X-BYTERANGE:4500\ntest.ts", item.String())
}
