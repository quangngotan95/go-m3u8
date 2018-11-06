package test

import (
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapItem_Parse(t *testing.T) {
	line := `#EXT-X-MAP:URI="frelo/prog_index.m3u8",BYTERANGE="3500@300"`

	mi, err := m3u8.NewMapItem(line)
	assert.Nil(t, err)
	assert.Equal(t, "frelo/prog_index.m3u8", mi.URI)
	assert.NotNil(t, mi.ByteRange)
	assertNotNilEqual(t, 3500, mi.ByteRange.Length)
	assertNotNilEqual(t, 300, mi.ByteRange.Start)

	assertToString(t, line, mi)
}

func TestMapItem_Parse_2(t *testing.T) {
	line := `#EXT-X-MAP:URI="frelo/prog_index.m3u8"`

	mi, err := m3u8.NewMapItem(line)
	assert.Nil(t, err)
	assert.Equal(t, "frelo/prog_index.m3u8", mi.URI)
	assert.Nil(t, mi.ByteRange)

	assertToString(t, line, mi)
}
