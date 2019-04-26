package test

import (
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	p        *m3u8.Playlist
	expected string
}

func TestWriter_Master(t *testing.T) {
	testCases := []testCase{
		// Master playlist
		{
			&m3u8.Playlist{
				Target: 10,
				Items: []m3u8.Item{
					&m3u8.PlaylistItem{
						ProgramID:  pointer.ToString("1"),
						URI:        "playlist_url",
						Bandwidth:  6400,
						AudioCodec: pointer.ToString("mp3"),
					},
					&m3u8.PlaylistItem{
						ProgramID:  pointer.ToString("2"),
						URI:        "playlist_url",
						Bandwidth:  50000,
						AudioCodec: pointer.ToString("aac-lc"),
						Width:      pointer.ToInt(1920),
						Height:     pointer.ToInt(1080),
						Profile:    pointer.ToString("high"),
						Level:      pointer.ToString("4.1"),
					},
					&m3u8.SessionDataItem{
						DataID:   "com.test.movie.title",
						Value:    pointer.ToString("Test"),
						URI:      pointer.ToString("http://test"),
						Language: pointer.ToString("en"),
					},
				},
			},
			`#EXTM3U
#EXT-X-STREAM-INF:PROGRAM-ID=1,CODECS="mp4a.40.34",BANDWIDTH=6400
playlist_url
#EXT-X-STREAM-INF:PROGRAM-ID=2,RESOLUTION=1920x1080,CODECS="avc1.640029,mp4a.40.2",BANDWIDTH=50000
playlist_url
#EXT-X-SESSION-DATA:DATA-ID="com.test.movie.title",VALUE="Test",URI="http://test",LANGUAGE="en"
`,
		},
		// Master playlist with single stream
		{
			&m3u8.Playlist{
				Target: 10,
				Items: []m3u8.Item{
					&m3u8.PlaylistItem{
						ProgramID:  pointer.ToString("1"),
						URI:        "playlist_url",
						Bandwidth:  6400,
						AudioCodec: pointer.ToString("mp3"),
					},
				},
			},
			`#EXTM3U
#EXT-X-STREAM-INF:PROGRAM-ID=1,CODECS="mp4a.40.34",BANDWIDTH=6400
playlist_url
`,
		},
		// Master playlist with header options
		{
			&m3u8.Playlist{
				Target:              10,
				Version:             pointer.ToInt(6),
				IndependentSegments: true,
				Items: []m3u8.Item{
					&m3u8.PlaylistItem{
						URI:        "playlist_url",
						Bandwidth:  6400,
						AudioCodec: pointer.ToString("mp3"),
					},
				},
			},
			`#EXTM3U
#EXT-X-VERSION:6
#EXT-X-INDEPENDENT-SEGMENTS
#EXT-X-STREAM-INF:CODECS="mp4a.40.34",BANDWIDTH=6400
playlist_url
`,
		},
		// New master playlist
		{
			&m3u8.Playlist{
				Master: pointer.ToBool(true),
			},
			`#EXTM3U
`,
		},
		// New media playlist
		{
			&m3u8.Playlist{
				Target: 10,
			},
			`#EXTM3U
#EXT-X-MEDIA-SEQUENCE:0
#EXT-X-TARGETDURATION:10
#EXT-X-ENDLIST
`,
		},
		// Media playlist
		{
			&m3u8.Playlist{
				Version:               pointer.ToInt(4),
				Cache:                 pointer.ToBool(false),
				Target:                6,
				Sequence:              1,
				DiscontinuitySequence: pointer.ToInt(10),
				Type:                  pointer.ToString("EVENT"),
				IFramesOnly:           true,
				Items: []m3u8.Item{
					&m3u8.SegmentItem{
						Duration: 11.344644,
						Segment:  "1080-7mbps00000.ts",
					},
				},
			},
			`#EXTM3U
#EXT-X-PLAYLIST-TYPE:EVENT
#EXT-X-VERSION:4
#EXT-X-I-FRAMES-ONLY
#EXT-X-MEDIA-SEQUENCE:1
#EXT-X-DISCONTINUITY-SEQUENCE:10
#EXT-X-ALLOW-CACHE:NO
#EXT-X-TARGETDURATION:6
#EXTINF:11.344644,
1080-7mbps00000.ts
#EXT-X-ENDLIST
`,
		},
		// Media playlist with keys
		{
			&m3u8.Playlist{
				Target:  10,
				Version: pointer.ToInt(7),
				Items: []m3u8.Item{
					&m3u8.SegmentItem{
						Duration: 11.344644,
						Segment:  "1080-7mbps00000.ts",
					},
					&m3u8.KeyItem{
						Encryptable: &m3u8.Encryptable{
							Method:            "AES-128",
							URI:               pointer.ToString("http://test.key"),
							IV:                pointer.ToString("D512BBF"),
							KeyFormat:         pointer.ToString("identity"),
							KeyFormatVersions: pointer.ToString("1/3"),
						},
					},
					&m3u8.SegmentItem{
						Duration: 11.261233,
						Segment:  "1080-7mbps00001.ts",
					},
				},
			},
			`#EXTM3U
#EXT-X-VERSION:7
#EXT-X-MEDIA-SEQUENCE:0
#EXT-X-TARGETDURATION:10
#EXTINF:11.344644,
1080-7mbps00000.ts
#EXT-X-KEY:METHOD=AES-128,URI="http://test.key",IV=D512BBF,KEYFORMAT="identity",KEYFORMATVERSIONS="1/3"
#EXTINF:11.261233,
1080-7mbps00001.ts
#EXT-X-ENDLIST
`,
		},
	}
	for _, tc := range testCases {
		tc.assert(t)
	}

	p := &m3u8.Playlist{
		Target: 10,
		Items: []m3u8.Item{
			&m3u8.PlaylistItem{
				ProgramID: pointer.ToString("1"),
				Width:     pointer.ToInt(1920),
				Height:    pointer.ToInt(1080),
				Codecs:    pointer.ToString("avc"),
				Bandwidth: 540,
				URI:       "test.url",
			},
			&m3u8.SegmentItem{
				Duration: 10.991,
				Segment:  "test.ts",
			},
		},
	}
	_, err := m3u8.Write(p)
	assert.Equal(t, m3u8.ErrPlaylistInvalidType, err)
}

func (tc testCase) assert(t *testing.T) {
	assert.Equal(t, tc.expected, tc.p.String())
}
