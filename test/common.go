package test

import (
	"strings"
	"testing"

	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func assertNotNilEqual(t *testing.T, expected interface{}, ptr interface{}) {
	assert.NotNil(t, ptr)
	switch ptr.(type) {
	case *string:
		s, ok := ptr.(*string)
		assert.True(t, ok)
		assert.Equal(t, expected, *s)
	case *float64:
		f, ok := ptr.(*float64)
		assert.True(t, ok)
		assert.Equal(t, expected, *f)
	case *int:
		i, ok := ptr.(*int)
		assert.True(t, ok)
		assert.Equal(t, expected, *i)
	case *bool:
		b, ok := ptr.(*bool)
		assert.True(t, ok)
		assert.Equal(t, expected, *b)
	default:
		t.Fatal("not supported assert type")
	}
}

func assertEqualWithoutNewLine(t *testing.T, expected string, actual string) {
	removedNewLine := strings.Replace(expected, "\n", "", -1)
	assert.Equal(t, removedNewLine, actual)
}

func assertToString(t *testing.T, expected string, item m3u8.Item) {
	assertEqualWithoutNewLine(t, expected, item.String())
}
