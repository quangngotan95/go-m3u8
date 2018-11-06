package m3u8

import "errors"

var (
	ErrPlaylistInvalid     = errors.New("invalid playlist, must start with #EXTM3U")
	ErrPlaylistInvalidType = errors.New("invalid playlist, mixed master and media")
	ErrResolutionInvalid   = errors.New("invalid resolution")
	ErrBandwidthMissing    = errors.New("missing bandwidth")
	ErrBandwidthInvalid    = errors.New("invalid bandwidth")
	ErrSegmentItemInvalid  = errors.New("invalid segment item")
	ErrPlaylistItemInvalid = errors.New("invalid playlist item")
)
