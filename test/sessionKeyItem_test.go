package test

import (
	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionKeyItem_Parse(t *testing.T) {
	line := `#EXT-X-SESSION-KEY:METHOD=AES-128,URI="http://test.key",IV=D512BBF,KEYFORMAT="identity",KEYFORMATVERSIONS="1/3"`

	ski, err := m3u8.NewSessionKeyItem(line)
	assert.Nil(t, err)
	assert.NotNil(t, ski.Encryptable)

	assert.Equal(t, "AES-128", ski.Encryptable.Method)
	assertNotNilEqual(t, "http://test.key", ski.Encryptable.URI)
	assertNotNilEqual(t, "D512BBF", ski.Encryptable.IV)
	assertNotNilEqual(t, "identity", ski.Encryptable.KeyFormat)
	assertNotNilEqual(t, "1/3", ski.Encryptable.KeyFormatVersions)

	assertToString(t, line, ski)
}
