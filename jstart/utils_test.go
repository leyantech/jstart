package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveOptionsWithPrefix(t *testing.T) {
	removed := RemoveOptionsWithPrefix([]string{"-Xmx256m", "-Xms128m"}, "-Xmx", "-Xms")
	assert.Equal(t, 0, len(removed))
}
