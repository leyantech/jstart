package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoteDebugRule(t *testing.T) {
	rule := RemoteDebugRule{}
	options := rule.ConvertOptions("8", []string{}, "off")
	assert.Equal(t, []string{}, options)
	options = rule.ConvertOptions("8", []string{"-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=3330"}, "off")
	assert.Equal(t, []string{}, options)
	options = rule.ConvertOptions("8", []string{}, "on")
	assert.Equal(t, []string{"-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=3330"}, options)
	options = rule.ConvertOptions("9", []string{}, "on")
	assert.Equal(t, []string{"-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:3330"}, options)
}
