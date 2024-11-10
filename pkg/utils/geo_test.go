package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePosition(t *testing.T) {
	pos, err := ParsePosition("")
	assert.Error(t, err)

	pos, err = ParsePosition("1")
	assert.Error(t, err)

	pos, err = ParsePosition("1;")
	assert.Error(t, err)

	pos, err = ParsePosition("1;2;3")
	assert.Error(t, err)

	pos, err = ParsePosition("1")
	assert.Error(t, err)

	pos, err = ParsePosition("12.345;67.89012")
	assert.Nil(t, err)
	assert.InDelta(t, 12.345, pos.Latitude, 0.00001)
	assert.InDelta(t, 67.89012, pos.Longitude, 0.00001)

	str := PositionToString(pos)
	assert.Equal(t, "12.34500;67.89012", str)
}
