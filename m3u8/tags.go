package m3u8

const (
	// Item tags
	SessionKeyItemTag    = `#EXT-X-SESSION-KEY`
	KeyItemTag           = `#EXT-X-KEY`
	DiscontinuityItemTag = `#EXT-X-DISCONTINUITY`
	TimeItemTag          = `#EXT-X-PROGRAM-DATE-TIME`
	DateRangeItemTag     = `#EXT-X-DATERANGE`
	MapItemTag           = `#EXT-X-MAP`
	SessionDataItemTag   = `#EXT-X-SESSION-DATA`
	SegmentItemTag       = `#EXTINF`
	ByteRangeItemTag     = `#EXT-X-BYTERANGE`
	PlaybackStartTag     = `#EXT-X-START`
	MediaItemTag         = `#EXT-X-MEDIA`
	PlaylistItemTag      = `#EXT-X-STREAM-INF`
	PlaylistIframeTag    = `#EXT-X-I-FRAME-STREAM-INF`

	// Playlist tags
	HeaderTag                = `#EXTM3U`
	FooterTag                = `#EXT-X-ENDLIST`
	TargetDurationTag        = `#EXT-X-TARGETDURATION`
	CacheTag                 = `#EXT-X-ALLOW-CACHE`
	DiscontinuitySequenceTag = `#EXT-X-DISCONTINUITY-SEQUENCE`
	IndependentSegmentsTag   = `#EXT-X-INDEPENDENT-SEGMENTS`
	PlaylistTypeTag          = `#EXT-X-PLAYLIST-TYPE`
	IFramesOnlyTag           = `#EXT-X-I-FRAMES-ONLY`
	MediaSequenceTag         = `#EXT-X-MEDIA-SEQUENCE`
	VersionTag               = `#EXT-X-VERSION`

	// ByteRange tags
	ByteRangeTag = "BYTERANGE"

	// Encryptable tags
	MethodTag            = "METHOD"
	URITag               = "URI"
	IVTag                = "IV"
	KeyFormatTag         = "KEYFORMAT"
	KeyFormatVersionsTag = "KEYFORMATVERSIONS"

	// DateRangeItem tags
	IDTag              = "ID"
	ClassTag           = "CLASS"
	StartDateTag       = "START-DATE"
	EndDateTag         = "END-DATE"
	DurationTag        = "DURATION"
	PlannedDurationTag = "PLANNED-DURATION"
	Scte35CmdTag       = "SCTE35-CMD"
	Scte35OutTag       = "SCTE35-OUT"
	Scte35InTag        = "SCTE35-IN"
	EndOnNextTag       = "END-ON-NEXT"

	// PlaybackStart tags
	TimeOffsetTag = "TIME-OFFSET"
	PreciseTag    = "PRECISE"

	// SessionDataItem tags
	DataIDTag   = "DATA-ID"
	ValueTag    = "VALUE"
	LanguageTag = "LANGUAGE"

	// MediaItem tags
	TypeTag            = "TYPE"
	GroupIDTag         = "GROUP-ID"
	AssocLanguageTag   = "ASSOC-LANGUAGE"
	NameTag            = "NAME"
	AutoSelectTag      = "AUTOSELECT"
	DefaultTag         = "DEFAULT"
	ForcedTag          = "FORCED"
	InStreamIDTag      = "INSTREAM-ID"
	CharacteristicsTag = "CHARACTERISTICS"
	ChannelsTag        = "CHANNELS"

	/// PlaylistItem tags
	ResolutionTag       = "RESOLUTION"
	ProgramIDTag        = "PROGRAM-ID"
	CodecsTag           = "CODECS"
	BandwidthTag        = "BANDWIDTH"
	AverageBandwidthTag = "AVERAGE-BANDWIDTH"
	FrameRateTag        = "FRAME-RATE"
	VideoTag            = "VIDEO"
	AudioTag            = "AUDIO"
	SubtitlesTag        = "SUBTITLES"
	ClosedCaptionsTag   = "CLOSED-CAPTIONS"
	HDCPLevelTag        = "HDCP-LEVEL"

	// Values
	NoneValue = "NONE"
	YesValue  = "YES"
	NoValue   = "NO"
)
