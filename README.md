[![Build Status](https://travis-ci.org/quangngotan95/go-m3u8.svg?branch=master)](https://travis-ci.org/quangngotan95/go-m3u8)
[![codecov](https://codecov.io/gh/quangngotan95/go-m3u8/branch/master/graph/badge.svg)](https://codecov.io/gh/quangngotan95/go-m3u8)
[![Go Report Card](https://goreportcard.com/badge/github.com/quangngotan95/go-m3u8)](https://goreportcard.com/report/github.com/quangngotan95/go-m3u8)
[![GoDoc](https://godoc.org/github.com/quangngotan95/go-m3u8/m3u8?status.svg)](https://godoc.org/github.com/quangngotan95/go-m3u8/m3u8)
# go-m3u8
Golang package for m3u8 (ported m3u8 gem https://github.com/sethdeckard/m3u8)

`go-m3u8` provides easy generation and parsing of m3u8 playlists defined in the HTTP Live Streaming (HLS) Internet Draft published by Apple.
* The library completely implements version 20 of the HLS Internet Draft.
* Provides parsing of an m3u8 playlist into an object model from any File, io.Reader or string.
* Provides ability to write playlist to a string via String()
* Distinction between a master and media playlist is handled automatically (single Playlist class).
* Optionally, the library can automatically generate the audio/video codecs string used in the CODEC attribute based on specified H.264, AAC, or MP3 options (such as Profile/Level).

## Installation
`go get github.com/quangngotan95/go-m3u8`

## Usage (creating playlists)
Create a master playlist and child playlists for adaptive bitrate streaming:
```go
import (
    "github.com/quangngotan95/go-m3u8/m3u8"
    "github.com/AlekSi/pointer"
)

playlist := m3u8.NewPlaylist()
```
Create a new playlist item:
```go
item := &m3u8.PlaylistItem{
    Width:      pointer.ToInt(1920),
    Height:     pointer.ToInt(1080),
    Profile:    pointer.ToString("high"),
    Level:      pointer.ToString("4.1"),
    AudioCodec: pointer.ToString("aac-lc"),
    Bandwidth:  540,
    URI:        "test.url",
}
playlist.AppendItem(item)
```
Add alternate audio, camera angles, closed captions and subtitles by creating MediaItem instances and adding them to the Playlist:
```go
item := &m3u8.MediaItem{
    Type:          "AUDIO",
    GroupID:       "audio-lo",
    Name:          "Francais",
    Language:      pointer.ToString("fre"),
    AssocLanguage: pointer.ToString("spoken"),
    AutoSelect:    pointer.ToBool(true),
    Default:       pointer.ToBool(false),
    Forced:        pointer.ToBool(true),
    URI:           pointer.ToString("frelo/prog_index.m3u8"),
}
playlist.AppendItem(item)
```
Create a standard playlist and add MPEG-TS segments via SegmentItem. You can also specify options for this type of playlist, however these options are ignored if playlist becomes a master playlist (anything but segments added):
```go
playlist := &m3u8.Playlist{
    Target:   12,
    Sequence: 1,
    Version:  pointer.ToInt(1),
    Cache:    pointer.ToBool(false),
    Items: []m3u8.Item{
        &m3u8.SegmentItem{
            Duration: 11,
            Segment:  "test.ts",
        },
    },
}
```
You can also access the playlist as a string:
```go
var str string
str = playlist.String()
...
fmt.Print(playlist)
```
Alternatively you can set codecs rather than having it generated automatically:
```go
item := &m3u8.PlaylistItem{
    Width:     pointer.ToInt(1920),
    Height:    pointer.ToInt(1080),
    Codecs:    pointer.ToString("avc1.66.30,mp4a.40.2"),
    Bandwidth: 540,
    URI:       "test.url",
}
```

## Usage (parsing playlists)
Parse from file
```go
playlist, err := m3u8.ReadFile("path/to/file")
```
Read from string
```go
playlist, err := m3u8.ReadString(string)
```
Read from generic `io.Reader`
```go
playlist, err := m3u8.Read(reader)
```

Access items in playlist:
```go
gore> playlist.Items[0]
(*m3u8.SessionKeyItem)#EXT-X-SESSION-KEY:METHOD=AES-128,URI="https://priv.example.com/key.php?r=52"
gore> playlist.Items[1]
(*m3u8.PlaybackStart)#EXT-X-START:TIME-OFFSET=20.2
```

## Misc
Codecs:
* Values for audio_codec (codec name): aac-lc, he-aac, mp3
* Values for profile (H.264 Profile): baseline, main, high.
* Values for level (H.264 Level): 3.0, 3.1, 4.0, 4.1.

Not all Levels and Profiles can be combined and validation is not currently implemented, consult H.264 documentation for further details.

## Contributing
1. Fork it https://github.com/quangngotan95/go-m3u8/fork
2. Create your feature branch `git checkout -b my-new-feature`
3. Run tests `go test ./test/...`, make sure they all pass and new features are covered
4. Commit your changes `git commit -am "Add new features"`
5. Push to the branch `git push origin my-new-feature`
6. Create a new Pull Request

## License
MIT License - See [LICENSE](https://github.com/quangngotan95/go-m3u8/blob/master/LICENSE) for details
