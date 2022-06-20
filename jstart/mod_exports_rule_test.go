package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModExportsRule(t *testing.T) {
	rule := ModExportsRule{}
	options := rule.ConvertOptions("8", []string{}, "")
	assert.Equal(t, []string{}, options)

	options = rule.ConvertOptions("11", []string{}, "")
	assert.Equal(t, modExportsOptions, options)
}
