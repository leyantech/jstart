package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectMemRule(t *testing.T) {
	rule := DirectMemRule{}
	options := rule.ConvertOptions("8", []string{}, "")
	assert.Equal(t, 0, len(options))

	options = rule.ConvertOptions("8", []string{}, "64")
	assert.Equal(t, []string{"-XX:MaxDirectMemorySize=64m"}, options)
}
