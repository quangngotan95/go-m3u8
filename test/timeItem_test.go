package test

import (
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeItem_New(t *testing.T) {
	timeVar, err := m3u8.ParseTime("2010-02-19T14:54:23.031Z")
	assert.Nil(t, err)
	ti := &m3u8.TimeItem{
		Time: timeVar,
	}

	assert.Equal(t, "#EXT-X-PROGRAM-DATE-TIME:2010-02-19T14:54:23.031Z", ti.String())
}

func TestTimeItem_Parse(t *testing.T) {
	ti, err := m3u8.NewTimeItem("#EXT-X-PROGRAM-DATE-TIME:2010-02-19T14:54:23.031Z")
	assert.Nil(t, err)

	expected, err := time.Parse(time.RFC3339Nano, "2010-02-19T14:54:23.031Z")
	assert.Nil(t, err)

	assert.Equal(t, expected, ti.Time)
}
