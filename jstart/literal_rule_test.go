package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLiteralRule(t *testing.T) {
	rule := LiteralJavaOptionsRule{}
	options := rule.ConvertOptions(nil, []string{}, "-server -XX:+HeapDumpOnOutOfMemoryError")
	assert.Equal(t, []string{"-server", "-XX:+HeapDumpOnOutOfMemoryError"}, options)
}
