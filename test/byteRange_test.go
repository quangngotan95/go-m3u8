package test

import (
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestByteRange_Parse(t *testing.T) {
	text := "4500@600"
	br, err := m3u8.NewByteRange(text)

	assert.Nil(t, err)
	assert.NotNil(t, br.Length)
	assert.NotNil(t, br.Start)

	assert.Equal(t, 4500, *br.Length)
	assert.Equal(t, 600, *br.Start)

	assertToString(t, text, br)
}

func TestByteRange_Parse_2(t *testing.T) {
	text := "4500"
	br, err := m3u8.NewByteRange(text)

	assert.Nil(t, err)
	assert.NotNil(t, br.Length)
	assert.Nil(t, br.Start)

	assert.Equal(t, 4500, *br.Length)

	assertToString(t, text, br)
}

func TestByteRange_New(t *testing.T) {
	br := &m3u8.ByteRange{
		Length: pointer.ToInt(4500),
		Start:  pointer.ToInt(200),
	}
	assert.Equal(t, "4500@200", br.String())
}

func TestByteRange_New_2(t *testing.T) {
	br := &m3u8.ByteRange{
		Length: pointer.ToInt(4500),
	}
	assert.Equal(t, "4500", br.String())
}
