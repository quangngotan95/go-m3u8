package test

import (
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAttributes(t *testing.T) {
	line := "TEST-ID=\"Help\",URI=\"http://test\",ID=33\n"
	mapAttr := m3u8.ParseAttributes(line)

	assert.NotNil(t, mapAttr)
	assert.Equal(t, "Help", mapAttr["TEST-ID"])
	assert.Equal(t, "http://test", mapAttr["URI"])
	assert.Equal(t, "33", mapAttr["ID"])
}
