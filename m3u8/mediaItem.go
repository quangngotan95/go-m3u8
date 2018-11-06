package m3u8

import (
	"fmt"
	"strings"
)

// MediaItem represents a set of EXT-X-MEDIA attributes
type MediaItem struct {
	Type            string
	GroupID         string
	Name            string
	Language        *string
	AssocLanguage   *string
	AutoSelect      *bool
	Default         *bool
	Forced          *bool
	URI             *string
	InStreamID      *string
	Characteristics *string
	Channels        *string
}

func NewMediaItem(text string) (*MediaItem, error) {
	attributes := ParseAttributes(text)

	return &MediaItem{
		Type:            attributes[TypeTag],
		GroupID:         attributes[GroupIDTag],
		Name:            attributes[NameTag],
		Language:        pointerTo(attributes, LanguageTag),
		AssocLanguage:   pointerTo(attributes, AssocLanguageTag),
		AutoSelect:      parseYesNo(attributes, AutoSelectTag),
		Default:         parseYesNo(attributes, DefaultTag),
		Forced:          parseYesNo(attributes, ForcedTag),
		URI:             pointerTo(attributes, URITag),
		InStreamID:      pointerTo(attributes, InStreamIDTag),
		Characteristics: pointerTo(attributes, CharacteristicsTag),
		Channels:        pointerTo(attributes, ChannelsTag),
	}, nil
}

func (mi *MediaItem) String() string {
	slice := []string{
		fmt.Sprintf(formatString, TypeTag, mi.Type),
		fmt.Sprintf(quotedFormatString, GroupIDTag, mi.GroupID),
	}

	if mi.Language != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, LanguageTag, *mi.Language))
	}
	if mi.AssocLanguage != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, AssocLanguageTag, *mi.AssocLanguage))
	}
	slice = append(slice, fmt.Sprintf(quotedFormatString, NameTag, mi.Name))
	if mi.AutoSelect != nil {
		slice = append(slice, fmt.Sprintf(formatString, AutoSelectTag, formatYesNo(*mi.AutoSelect)))
	}
	if mi.Default != nil {
		slice = append(slice, fmt.Sprintf(formatString, DefaultTag, formatYesNo(*mi.Default)))
	}
	if mi.URI != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, URITag, *mi.URI))
	}
	if mi.Forced != nil {
		slice = append(slice, fmt.Sprintf(formatString, ForcedTag, formatYesNo(*mi.Forced)))
	}
	if mi.InStreamID != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, InStreamIDTag, *mi.InStreamID))
	}
	if mi.Characteristics != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, CharacteristicsTag, *mi.Characteristics))
	}
	if mi.Channels != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, ChannelsTag, *mi.Channels))
	}

	return fmt.Sprintf("%s:%s", MediaItemTag, strings.Join(slice, ","))
}
