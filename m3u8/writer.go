package m3u8

import (
	"fmt"
	"strings"
)

func Write(pl *Playlist) (string, error) {
	var sb strings.Builder

	if !pl.IsValid() {
		return "", ErrPlaylistInvalidType
	}
	writeHeader(&sb, pl)
	for _, item := range pl.Items {
		sb.WriteString(item.String())
		sb.WriteRune('\n')
	}
	writeFooter(&sb, pl)

	return sb.String(), nil
}

func writeHeader(sb *strings.Builder, pl *Playlist) {
	sb.WriteString(HeaderTag)
	sb.WriteRune('\n')

	if pl.IsMaster() {
		writeVersionTag(sb, pl.Version)
		writeIndependentSegmentsTag(sb, pl.IndependentSegments)
	} else {
		if pl.Type != nil {
			sb.WriteString(fmt.Sprintf("%s:%s", PlaylistTypeTag, *pl.Type))
			sb.WriteRune('\n')
		}
		writeVersionTag(sb, pl.Version)
		writeIndependentSegmentsTag(sb, pl.IndependentSegments)
		if pl.IFramesOnly {
			sb.WriteString(IFramesOnlyTag)
			sb.WriteRune('\n')
		}
		sb.WriteString(fmt.Sprintf("%s:%v", MediaSequenceTag, pl.Sequence))
		sb.WriteRune('\n')
		writeDiscontinuitySequenceTag(sb, pl.DiscontinuitySequence)
		writeCacheTag(sb, pl.Cache)
		sb.WriteString(fmt.Sprintf("%s:%v", TargetDurationTag, pl.Target))
		sb.WriteRune('\n')
	}
}

func writeFooter(sb *strings.Builder, pl *Playlist) {
	if pl.IsLive() || pl.IsMaster() {
		return
	}

	sb.WriteString(FooterTag)
	sb.WriteRune('\n')
}

func writeVersionTag(sb *strings.Builder, version *int) {
	if version == nil {
		return
	}

	sb.WriteString(fmt.Sprintf("%s:%v", VersionTag, *version))
	sb.WriteRune('\n')
}

func writeIndependentSegmentsTag(sb *strings.Builder, toWrite bool) {
	if !toWrite {
		return
	}

	sb.WriteString(IndependentSegmentsTag)
	sb.WriteRune('\n')
}

func writeDiscontinuitySequenceTag(sb *strings.Builder, sequence *int) {
	if sequence == nil {
		return
	}

	sb.WriteString(fmt.Sprintf("%s:%v", DiscontinuitySequenceTag, *sequence))
	sb.WriteRune('\n')
}

func writeCacheTag(sb *strings.Builder, cache *bool) {
	if cache == nil {
		return
	}

	sb.WriteString(fmt.Sprintf("%s:%s", CacheTag, formatYesNo(*cache)))
	sb.WriteRune('\n')
}
