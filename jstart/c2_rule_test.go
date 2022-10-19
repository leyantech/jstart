package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestC2Rule(t *testing.T) {
	rule := C2Rule{}
	options := rule.ConvertOptions(nil, []string{}, "on")
	assert.Equal(t, []string{"-XX:+TieredCompilation", "-XX:TieredStopAtLevel=4"}, options)

	options = rule.ConvertOptions(nil, []string{"-Xint"}, "off")
	assert.Equal(t, []string{"-XX:+TieredCompilation", "-XX:TieredStopAtLevel=1"}, options)

	options = rule.ConvertOptions(nil, []string{"-Xint"}, "3")
	assert.Equal(t, []string{"-XX:+TieredCompilation", "-XX:TieredStopAtLevel=3"}, options)

	options = rule.ConvertOptions(&Context{MemoryLimit: 1024}, []string{"-Xint"}, "auto")
	assert.Equal(t, []string{"-XX:+TieredCompilation", "-XX:TieredStopAtLevel=4"}, options)

	options = rule.ConvertOptions(&Context{MemoryLimit: 300}, []string{"-Xint"}, "auto")
	assert.Equal(t, []string{"-XX:+TieredCompilation", "-XX:TieredStopAtLevel=1"}, options)

	options = rule.ConvertOptions(nil, []string{"-Xint"}, "autooo")
	assert.Equal(t, 1, len(options))
}
