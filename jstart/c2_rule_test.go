package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestC2Rule(t *testing.T) {
	rule := C2Rule{}
	options := rule.ConvertOptions("8", []string{}, "on")
	assert.Equal(t, []string{"-XX:+TieredCompilation", "-XX:TieredStopAtLevel=4"}, options)

	options = rule.ConvertOptions("8", []string{"-Xint"}, "off")
	assert.Equal(t, []string{"-XX:+TieredCompilation", "-XX:TieredStopAtLevel=1"}, options)

	options = rule.ConvertOptions("8", []string{"-Xint"}, "3")
	assert.Equal(t, []string{"-XX:+TieredCompilation", "-XX:TieredStopAtLevel=3"}, options)

	options = rule.ConvertOptions("8", []string{"-Xint"}, "auto")
	assert.Equal(t, 2, len(options))

	options = rule.ConvertOptions("8", []string{"-Xint"}, "autooo")
	assert.Equal(t, 1, len(options))
}
