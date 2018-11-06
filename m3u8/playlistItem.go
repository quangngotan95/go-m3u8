package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// PlaylistItem represents a set of EXT-X-STREAM-INF or
// EXT-X-I-FRAME-STREAM-INF attributes
type PlaylistItem struct {
	Bandwidth int
	URI       string
	IFrame    bool

	Name             *string
	Width            *int
	Height           *int
	AverageBandwidth *int
	ProgramID        *string
	Codecs           *string
	AudioCodec       *string
	Profile          *string
	Level            *string
	Video            *string
	Audio            *string
	Subtitles        *string
	ClosedCaptions   *string
	FrameRate        *float64
	HDCPLevel        *string
	Resolution       *Resolution
}

func NewPlaylistItem(text string) (*PlaylistItem, error) {
	attributes := ParseAttributes(text)

	resolution, err := parseResolution(attributes, ResolutionTag)
	if err != nil {
		return nil, err
	}
	var width, height *int
	if resolution != nil {
		width = &resolution.Width
		height = &resolution.Height
	}

	averageBandwidth, err := parseInt(attributes, AverageBandwidthTag)
	if err != nil {
		return nil, err
	}

	frameRate, err := parseFloat(attributes, FrameRateTag)
	if err != nil {
		return nil, err
	}
	if frameRate != nil && *frameRate <= 0 {
		frameRate = nil
	}

	bandwidth, err := parseBandwidth(attributes, BandwidthTag)
	if err != nil {
		return nil, err
	}

	return &PlaylistItem{
		ProgramID:        pointerTo(attributes, ProgramIDTag),
		Codecs:           pointerTo(attributes, CodecsTag),
		Width:            width,
		Height:           height,
		Bandwidth:        bandwidth,
		AverageBandwidth: averageBandwidth,
		FrameRate:        frameRate,
		Video:            pointerTo(attributes, VideoTag),
		Audio:            pointerTo(attributes, AudioTag),
		URI:              attributes[URITag],
		Subtitles:        pointerTo(attributes, SubtitlesTag),
		ClosedCaptions:   pointerTo(attributes, ClosedCaptionsTag),
		Name:             pointerTo(attributes, NameTag),
		HDCPLevel:        pointerTo(attributes, HDCPLevelTag),
		Resolution:       resolution,
	}, nil
}

func (pi *PlaylistItem) String() string {
	var slice []string
	// Check resolution
	if pi.Resolution == nil && pi.Width != nil && pi.Height != nil {
		r := &Resolution{
			Width:  *pi.Width,
			Height: *pi.Height,
		}
		pi.Resolution = r
	}
	if pi.ProgramID != nil {
		slice = append(slice, fmt.Sprintf(formatString, ProgramIDTag, *pi.ProgramID))
	}
	if pi.Resolution != nil {
		slice = append(slice, fmt.Sprintf(formatString, ResolutionTag, pi.Resolution.String()))
	}
	codecs := formatCodecs(pi)
	if codecs != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, CodecsTag, *codecs))
	}
	slice = append(slice, fmt.Sprintf(formatString, BandwidthTag, pi.Bandwidth))
	if pi.AverageBandwidth != nil {
		slice = append(slice, fmt.Sprintf(formatString, AverageBandwidthTag, *pi.AverageBandwidth))
	}
	if pi.FrameRate != nil {
		slice = append(slice, fmt.Sprintf(frameRateFormatString, FrameRateTag, *pi.FrameRate))
	}
	if pi.HDCPLevel != nil {
		slice = append(slice, fmt.Sprintf(formatString, HDCPLevelTag, *pi.HDCPLevel))
	}
	if pi.Audio != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, AudioTag, *pi.Audio))
	}
	if pi.Video != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, VideoTag, *pi.Video))
	}
	if pi.Subtitles != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, SubtitlesTag, *pi.Subtitles))
	}
	if pi.ClosedCaptions != nil {
		cc := *pi.ClosedCaptions
		fs := quotedFormatString
		if cc == NoneValue {
			fs = formatString
		}
		slice = append(slice, fmt.Sprintf(fs, ClosedCaptionsTag, cc))
	}
	if pi.Name != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, NameTag, *pi.Name))
	}

	attributesString := strings.Join(slice, ",")

	if pi.IFrame {
		return fmt.Sprintf(`%s:%s,%s="%s"`, PlaylistIframeTag, attributesString, URITag, pi.URI)
	}

	return fmt.Sprintf("%s:%s\n%s", PlaylistItemTag, attributesString, pi.URI)
}

func (pi *PlaylistItem) CodecsString() string {
	codecsPtr := formatCodecs(pi)
	if codecsPtr == nil {
		return ""
	}

	return *codecsPtr
}

func formatCodecs(pi *PlaylistItem) *string {
	if pi.Codecs != nil {
		return pi.Codecs
	}

	videoCodecPtr := videoCodec(pi.Profile, pi.Level)
	// profile or level were specified but not recognized any codecs
	if !(pi.Profile == nil && pi.Level == nil) && videoCodecPtr == nil {
		return nil
	}

	audioCodecPtr := audioCodec(pi.AudioCodec)
	// audio codec was specified but not recognized
	if !(pi.AudioCodec == nil) && audioCodecPtr == nil {
		return nil
	}

	var slice []string
	if videoCodecPtr != nil {
		slice = append(slice, *videoCodecPtr)
	}
	if audioCodecPtr != nil {
		slice = append(slice, *audioCodecPtr)
	}

	if len(slice) <= 0 {
		return nil
	}

	value := strings.Join(slice, ",")
	return &value
}

func parseBandwidth(attributes map[string]string, key string) (int, error) {
	bw, ok := attributes[key]
	if !ok {
		return 0, ErrBandwidthMissing
	}

	bandwidth, err := strconv.ParseInt(bw, 0, 0)
	if err != nil {
		return 0, ErrBandwidthInvalid
	}

	return int(bandwidth), nil
}

func parseResolution(attributes map[string]string, key string) (*Resolution, error) {
	resolution, ok := attributes[key]
	if !ok {
		return nil, nil
	}

	return NewResolution(resolution)
}
