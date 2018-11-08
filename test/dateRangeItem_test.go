package test

import (
	"testing"

	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestDateRangeItem_Parse(t *testing.T) {
	line := `#EXT-X-DATERANGE:ID="splice-6FFFFFF0",CLASS="test_class",
START-DATE="2014-03-05T11:15:00Z",
END-DATE="2014-03-05T11:16:00Z",DURATION=60.1,
PLANNED-DURATION=59.993,
SCTE35-CMD=0xFC002F0000000000FF2,
SCTE35-OUT=0xFC002F0000000000FF0,
SCTE35-IN=0xFC002F0000000000FF1,
END-ON-NEXT=YES
`
	dri, err := m3u8.NewDateRangeItem(line)

	assert.Nil(t, err)
	assert.Equal(t, "splice-6FFFFFF0", dri.ID)
	assert.Equal(t, "2014-03-05T11:15:00Z", dri.StartDate)

	assertNotNilEqual(t, "test_class", dri.Class)
	assertNotNilEqual(t, "2014-03-05T11:16:00Z", dri.EndDate)
	assertNotNilEqual(t, 60.1, dri.Duration)
	assertNotNilEqual(t, 59.993, dri.PlannedDuration)
	assertNotNilEqual(t, "0xFC002F0000000000FF2", dri.Scte35Cmd)
	assertNotNilEqual(t, "0xFC002F0000000000FF0", dri.Scte35Out)
	assertNotNilEqual(t, "0xFC002F0000000000FF1", dri.Scte35In)
	assert.True(t, dri.EndOnNext)
	assert.Nil(t, dri.ClientAttributes)

	assertToString(t, line, dri)
}

func TestDateRangeItem_Parse_2(t *testing.T) {
	line := `#EXT-X-DATERANGE:ID="splice-6FFFFFF0",
START-DATE="2014-03-05T11:15:00Z"
`
	dri, err := m3u8.NewDateRangeItem(line)

	assert.Nil(t, err)
	assert.Equal(t, "splice-6FFFFFF0", dri.ID)
	assert.Equal(t, "2014-03-05T11:15:00Z", dri.StartDate)

	assert.Nil(t, dri.Class)
	assert.Nil(t, dri.EndDate)
	assert.Nil(t, dri.Duration)
	assert.Nil(t, dri.PlannedDuration)
	assert.Nil(t, dri.Scte35In)
	assert.Nil(t, dri.Scte35Out)
	assert.Nil(t, dri.Scte35Cmd)
	assert.Nil(t, dri.ClientAttributes)
	assert.False(t, dri.EndOnNext)

	assertToString(t, line, dri)
}

func TestDateRangeItem_Parse_3(t *testing.T) {
	line := `#EXT-X-DATERANGE:ID="splice-6FFFFFF0",
START-DATE="2014-03-05T11:15:00Z",
X-CUSTOM-VALUE="test_value"
`
	dri, err := m3u8.NewDateRangeItem(line)

	assert.Nil(t, err)
	assert.NotNil(t, dri.ClientAttributes)
	assert.Equal(t, "test_value", dri.ClientAttributes["X-CUSTOM-VALUE"])

	assertToString(t, line, dri)
}
