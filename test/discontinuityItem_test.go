package test

import (
	"testing"

	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestDiscontinuityItem_Parse(t *testing.T) {
	di, err := m3u8.NewDiscontinuityItem()
	assert.Nil(t, err)
	assert.Equal(t, m3u8.DiscontinuityItemTag+"\n", di.String())
}
