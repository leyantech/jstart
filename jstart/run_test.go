package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitJavaCliOptionsAndArgs(t *testing.T) {
	options, arguments := splitJavaOptionsAndArguments([]string{"-server", "-XX:+HeapDumpOnOutOfMemoryError", "-cp", "server.jar", "Main"})
	assert.Equal(t, 4, len(options))
	assert.Equal(t, []string{"Main"}, arguments)
}
