package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRollOfPaper(t *testing.T) {
	assert.True(t, IsRollOfPaper('@'))
	assert.False(t, IsRollOfPaper('#'))
	assert.False(t, IsRollOfPaper('A'))
	assert.False(t, IsRollOfPaper(' '))
}
