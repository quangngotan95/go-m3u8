package test

import (
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaybackStart_Parse(t *testing.T) {
	line := `#EXT-X-START:TIME-OFFSET=20.2,PRECISE=YES`

	ps, err := m3u8.NewPlaybackStart(line)
	assert.Nil(t, err)
	assert.Equal(t, 20.2, ps.TimeOffset)
	assertNotNilEqual(t, true, ps.Precise)

	assertToString(t, line, ps)
}

func TestPlaybackStart_Parse_2(t *testing.T) {
	line := `#EXT-X-START:TIME-OFFSET=-12.9`

	ps, err := m3u8.NewPlaybackStart(line)
	assert.Nil(t, err)
	assert.Equal(t, -12.9, ps.TimeOffset)
	assert.Nil(t, ps.Precise)

	assertToString(t, line, ps)
}
