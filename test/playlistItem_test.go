package test

import (
	"github.com/AlekSi/pointer"
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaylistItem_Parse(t *testing.T) {
	line := `#EXT-X-STREAM-INF:CODECS="avc",BANDWIDTH=540,
PROGRAM-ID=1,RESOLUTION=1920x1080,FRAME-RATE=23.976,
AVERAGE-BANDWIDTH=550,AUDIO="test",VIDEO="test2",
SUBTITLES="subs",CLOSED-CAPTIONS="caps",URI="test.url",
NAME="1080p",HDCP-LEVEL=TYPE-0`

	pi, err := m3u8.NewPlaylistItem(line)
	assert.Nil(t, err)
	assertNotNilEqual(t, "1", pi.ProgramID)
	assertNotNilEqual(t, "avc", pi.Codecs)
	assert.Equal(t, 540, pi.Bandwidth)
	assertNotNilEqual(t, 550, pi.AverageBandwidth)
	assertNotNilEqual(t, 1920, pi.Width)
	assertNotNilEqual(t, 1080, pi.Height)
	assertNotNilEqual(t, 23.976, pi.FrameRate)
	assertNotNilEqual(t, "test", pi.Audio)
	assertNotNilEqual(t, "test2", pi.Video)
	assertNotNilEqual(t, "subs", pi.Subtitles)
	assertNotNilEqual(t, "caps", pi.ClosedCaptions)
	assert.Equal(t, "test.url", pi.URI)
	assertNotNilEqual(t, "1080p", pi.Name)
	assert.False(t, pi.IFrame)
	assertNotNilEqual(t, "TYPE-0", pi.HDCPLevel)
}

func TestPlaylistItem_ToString(t *testing.T) {
	// No codecs specified
	p := &m3u8.PlaylistItem{
		Bandwidth: 540,
		URI:       "test.url",
	}
	assert.NotContains(t, p.String(), "CODECS")

	// Level not recognized
	p = &m3u8.PlaylistItem{
		Bandwidth: 540,
		URI:       "test.url",
		Level:     pointer.ToString("9001"),
	}
	assert.NotContains(t, p.String(), "CODECS")

	// Audio codec recognized but profile not recognized
	p = &m3u8.PlaylistItem{
		Bandwidth:  540,
		URI:        "test.url",
		Profile:    pointer.ToString("best"),
		Level:      pointer.ToString("9001"),
		AudioCodec: pointer.ToString("aac-lc"),
	}
	assert.NotContains(t, p.String(), "CODECS")

	// Profile and level not set, Audio codec recognized
	p = &m3u8.PlaylistItem{
		Bandwidth:  540,
		URI:        "test.url",
		AudioCodec: pointer.ToString("aac-lc"),
	}
	assert.Contains(t, p.String(), "CODECS")

	// Profile and level recognized, audio codec not recognized
	p = &m3u8.PlaylistItem{
		Bandwidth:  540,
		URI:        "test.url",
		Profile:    pointer.ToString("high"),
		Level:      pointer.ToString("4.1"),
		AudioCodec: pointer.ToString("fuzzy"),
	}
	assert.NotContains(t, p.String(), "CODECS")

	// Audio codec not set
	p = &m3u8.PlaylistItem{
		Bandwidth: 540,
		URI:       "test.url",
		Profile:   pointer.ToString("high"),
		Level:     pointer.ToString("4.1"),
	}
	assert.Contains(t, p.String(), `CODECS="avc1.640029"`)

	// Audio codec recognized
	p = &m3u8.PlaylistItem{
		Bandwidth:  540,
		URI:        "test.url",
		Profile:    pointer.ToString("high"),
		Level:      pointer.ToString("4.1"),
		AudioCodec: pointer.ToString("aac-lc"),
	}
	assert.Contains(t, p.String(), `CODECS="avc1.640029,mp4a.40.2"`)
}

func TestPlaylistItem_ToString_2(t *testing.T) {
	// All fields set
	p := &m3u8.PlaylistItem{
		Codecs:           pointer.ToString("avc"),
		Bandwidth:        540,
		URI:              "test.url",
		Audio:            pointer.ToString("test"),
		Video:            pointer.ToString("test2"),
		AverageBandwidth: pointer.ToInt(500),
		Subtitles:        pointer.ToString("subs"),
		FrameRate:        pointer.ToFloat64(30),
		ClosedCaptions:   pointer.ToString("caps"),
		Name:             pointer.ToString("SD"),
		HDCPLevel:        pointer.ToString("TYPE-0"),
		ProgramID:        pointer.ToString("1"),
	}

	expected := `#EXT-X-STREAM-INF:PROGRAM-ID=1,CODECS="avc",BANDWIDTH=540,AVERAGE-BANDWIDTH=500,FRAME-RATE=30.000,HDCP-LEVEL=TYPE-0,AUDIO="test",VIDEO="test2",SUBTITLES="subs",CLOSED-CAPTIONS="caps",NAME="SD"
test.url`
	assert.Equal(t, expected, p.String())

	// Closed captions is NONE
	p = &m3u8.PlaylistItem{
		ProgramID:      pointer.ToString("1"),
		Width:          pointer.ToInt(1920),
		Height:         pointer.ToInt(1080),
		Codecs:         pointer.ToString("avc"),
		Bandwidth:      540,
		URI:            "test.url",
		ClosedCaptions: pointer.ToString("NONE"),
	}

	expected = `#EXT-X-STREAM-INF:PROGRAM-ID=1,RESOLUTION=1920x1080,CODECS="avc",BANDWIDTH=540,CLOSED-CAPTIONS=NONE
test.url`
	assert.Equal(t, expected, p.String())

	// IFrame is true
	p = &m3u8.PlaylistItem{
		Codecs:           pointer.ToString("avc"),
		Bandwidth:        540,
		URI:              "test.url",
		IFrame:           true,
		Video:            pointer.ToString("test2"),
		AverageBandwidth: pointer.ToInt(550),
	}

	expected = `#EXT-X-I-FRAME-STREAM-INF:CODECS="avc",BANDWIDTH=540,AVERAGE-BANDWIDTH=550,VIDEO="test2",URI="test.url"`
	assert.Equal(t, expected, p.String())
}

func TestPlaylistItem_GenerateCodecs(t *testing.T) {
	assertCodecs(t, "", &m3u8.PlaylistItem{})
	assertCodecs(t, "test", &m3u8.PlaylistItem{Codecs: pointer.ToString("test")})
	assertCodecs(t, "mp4a.40.2", &m3u8.PlaylistItem{AudioCodec: pointer.ToString("aac-lc")})
	assertCodecs(t, "mp4a.40.2", &m3u8.PlaylistItem{AudioCodec: pointer.ToString("AAC-LC")})
	assertCodecs(t, "mp4a.40.5", &m3u8.PlaylistItem{AudioCodec: pointer.ToString("he-aac")})
	assertCodecs(t, "", &m3u8.PlaylistItem{AudioCodec: pointer.ToString("he-aac1")})
	assertCodecs(t, "mp4a.40.34", &m3u8.PlaylistItem{AudioCodec: pointer.ToString("mp3")})
	assertCodecs(t, "avc1.66.30", &m3u8.PlaylistItem{
		Profile: pointer.ToString("baseline"),
		Level:   pointer.ToString("3.0"),
	})
	assertCodecs(t, "avc1.66.30,mp4a.40.2", &m3u8.PlaylistItem{
		Profile:    pointer.ToString("baseline"),
		Level:      pointer.ToString("3.0"),
		AudioCodec: pointer.ToString("aac-lc"),
	})
	assertCodecs(t, "avc1.66.30,mp4a.40.34", &m3u8.PlaylistItem{
		Profile:    pointer.ToString("baseline"),
		Level:      pointer.ToString("3.0"),
		AudioCodec: pointer.ToString("mp3"),
	})
	assertCodecs(t, "avc1.42001f", &m3u8.PlaylistItem{
		Profile: pointer.ToString("baseline"),
		Level:   pointer.ToString("3.1"),
	})
	assertCodecs(t, "avc1.42001f,mp4a.40.5", &m3u8.PlaylistItem{
		Profile:    pointer.ToString("baseline"),
		Level:      pointer.ToString("3.1"),
		AudioCodec: pointer.ToString("he-aac"),
	})
	assertCodecs(t, "avc1.77.30", &m3u8.PlaylistItem{
		Profile: pointer.ToString("main"),
		Level:   pointer.ToString("3.0"),
	})
	assertCodecs(t, "avc1.77.30,mp4a.40.2", &m3u8.PlaylistItem{
		Profile:    pointer.ToString("main"),
		Level:      pointer.ToString("3.0"),
		AudioCodec: pointer.ToString("aac-lc"),
	})
	assertCodecs(t, "avc1.4d001f", &m3u8.PlaylistItem{
		Profile: pointer.ToString("main"),
		Level:   pointer.ToString("3.1"),
	})
	assertCodecs(t, "avc1.4d0028", &m3u8.PlaylistItem{
		Profile: pointer.ToString("main"),
		Level:   pointer.ToString("4.0"),
	})
	assertCodecs(t, "avc1.4d0029", &m3u8.PlaylistItem{
		Profile: pointer.ToString("main"),
		Level:   pointer.ToString("4.1"),
	})
	assertCodecs(t, "avc1.64001f", &m3u8.PlaylistItem{
		Profile: pointer.ToString("high"),
		Level:   pointer.ToString("3.1"),
	})
	assertCodecs(t, "avc1.640028", &m3u8.PlaylistItem{
		Profile: pointer.ToString("high"),
		Level:   pointer.ToString("4.0"),
	})
	assertCodecs(t, "avc1.640029", &m3u8.PlaylistItem{
		Profile: pointer.ToString("high"),
		Level:   pointer.ToString("4.1"),
	})
}

func assertCodecs(t *testing.T, codecs string, p *m3u8.PlaylistItem) {
	assert.Equal(t, codecs, p.CodecsString())
}
