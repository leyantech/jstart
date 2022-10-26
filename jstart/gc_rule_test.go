package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGCRule(t *testing.T) {
	rule := GcRule{}
	options := rule.ConvertOptions(&context{memoryLimit: 1024}, []string{}, "auto")
	assert.Equal(t, []string{"-XX:+UseSerialGC"}, options)

	options = rule.ConvertOptions(&context{memoryLimit: 2048}, []string{}, "auto")
	assert.Equal(t, []string{"-XX:+UseConcMarkSweepGC"}, options)

	options = rule.ConvertOptions(&context{memoryLimit: 3096}, []string{}, "auto")
	assert.Equal(t, []string{"-XX:+UseG1GC"}, options)

	options = rule.ConvertOptions(&context{memoryLimit: 1024}, []string{}, "cms")
	assert.Equal(t, []string{"-XX:+UseConcMarkSweepGC"}, options)

	options = rule.ConvertOptions(&context{memoryLimit: 1024}, []string{"-XX:+UseConcMarkSweepGC"}, "auto")
	assert.Equal(t, []string{"-XX:+UseConcMarkSweepGC"}, options)

}
