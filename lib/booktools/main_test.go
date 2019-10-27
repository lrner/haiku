package booktools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLetters(t *testing.T) {
	letterSlice := GetLetters(1, 5)
	assert.Equal(t, []string{"a", "a", "a", "a", "a"}, letterSlice)
	assert.Equal(t, len(letterSlice), 5)

	letterSlice = GetLetters(27, 5)
	assert.Equal(t, []string{"a", "a", "a", "b", "a"}, letterSlice)
	assert.Equal(t, len(letterSlice), 5)

	letterSlice = GetLetters(705, 4)
	assert.Equal(t, []string{"a", "b", "b", "c"}, letterSlice)
	assert.Equal(t, len(letterSlice), 4)

}

func TestFiveRow(t *testing.T) {
	slice, letterSlice := FiveRow(1)
	assert.Equal(t, slice, []int{1, 1, 1, 1, 1})
	assert.Equal(t, letterSlice, []string{"a", "a", "a", "a", "a"})

	slice, letterSlice = FiveRow(27)
	assert.Equal(t, slice, []int{1, 1, 1, 1, 1})
	assert.Equal(t, letterSlice, []string{"a", "a", "a", "b", "a"})

}

func TestSevenRow(t *testing.T) {
	slice, letterSlice := SevenRow(1)

	assert.Equal(t, slice, []int{1, 1, 1, 1, 1, 1, 1})
	assert.Equal(t, letterSlice, []string{"a", "a", "a", "a", "a", "a", "a"})

	slice, letterSlice = SevenRow(35)

	assert.Equal(t, slice, []int{1, 1, 1, 1, 1, 1, 1})
	assert.Equal(t, letterSlice, []string{"a", "a", "a", "a", "a", "b", "i"})

}
