package m3u8

import (
	"fmt"
	"strings"
)

type Encryptable struct {
	Method            string
	URI               *string
	IV                *string
	KeyFormat         *string
	KeyFormatVersions *string
}

func NewEncryptable(attributes map[string]string) *Encryptable {
	return &Encryptable{
		Method:            attributes[MethodTag],
		URI:               pointerTo(attributes, URITag),
		IV:                pointerTo(attributes, IVTag),
		KeyFormat:         pointerTo(attributes, KeyFormatTag),
		KeyFormatVersions: pointerTo(attributes, KeyFormatVersionsTag),
	}
}

func (e *Encryptable) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf(formatString, MethodTag, e.Method))
	if e.URI != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, URITag, *e.URI))
	}
	if e.IV != nil {
		slice = append(slice, fmt.Sprintf(formatString, IVTag, *e.IV))
	}
	if e.KeyFormat != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, KeyFormatTag, *e.KeyFormat))
	}
	if e.KeyFormatVersions != nil {
		slice = append(slice, fmt.Sprintf(quotedFormatString, KeyFormatVersionsTag, *e.KeyFormatVersions))
	}

	return strings.Join(slice, ",")
}
