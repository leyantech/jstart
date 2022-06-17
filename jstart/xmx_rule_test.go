package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXmxExpression(t *testing.T) {
	xmx, err := evaluateXmxExpression(1024, "quota*0.66")
	assert.Nil(t, err)
	assert.Equal(t, int64(675), xmx)

	xmx, err = evaluateXmxExpression(1024, "quota*2/3")
	assert.Nil(t, err)
	assert.Equal(t, int64(682), xmx)

	xmx, err = evaluateXmxExpression(1024, "0.66*quota")
	assert.Nil(t, err)
	assert.Equal(t, int64(675), xmx)

	xmx, err = evaluateXmxExpression(1024, "quota-512")
	assert.Nil(t, err)
	assert.Equal(t, int64(512), xmx)

	_, err = evaluateXmxExpression(1024, "quota*0.01")
	assert.NotNil(t, err)

	_, err = evaluateXmxExpression(1024, "quota")
	assert.NotNil(t, err)

	_, err = evaluateXmxExpression(1024, "quota>100")
	assert.NotNil(t, err)
}

func TestXmxRule(t *testing.T) {
	rule := XmxRule{}
	options := rule.ConvertOptions("8", []string{"-version"}, "quota*2/3")
	assert.LessOrEqual(t, len(options), 3)
}
