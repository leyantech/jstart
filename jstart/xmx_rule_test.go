package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArithmeticExpressionEval(t *testing.T) {
	variables := make(map[string]int64)
	variables["quota"] = 1024
	maxXmx := 1024 - minimumNonHeap
	xmx, err := evaluateArithmeticExpression("quota*0.66", minimumXmx, maxXmx, variables)
	assert.Nil(t, err)
	assert.Equal(t, int64(675), xmx)

	xmx, err = evaluateArithmeticExpression("quota*2/3", minimumXmx, maxXmx, variables)
	assert.Nil(t, err)
	assert.Equal(t, int64(682), xmx)

	xmx, err = evaluateArithmeticExpression("0.66*quota", minimumXmx, maxXmx, variables)
	assert.Nil(t, err)
	assert.Equal(t, int64(675), xmx)

	xmx, err = evaluateArithmeticExpression("quota-512", minimumXmx, maxXmx, variables)
	assert.Nil(t, err)
	assert.Equal(t, int64(512), xmx)

	_, err = evaluateArithmeticExpression("quota*0.01", minimumXmx, maxXmx, variables)
	assert.NotNil(t, err)

	_, err = evaluateArithmeticExpression("quota", minimumXmx, maxXmx, variables)
	assert.NotNil(t, err)

	_, err = evaluateArithmeticExpression("quota>100", minimumXmx, maxXmx, variables)
	assert.NotNil(t, err)
}

func TestXmxRule(t *testing.T) {
	rule := XmxRule{}
	options := rule.ConvertOptions(&Context{JdkVersion: "8", MemoryLimit: 1024}, []string{"-version"}, "quota*2/3")
	assert.NotEmpty(t, FindOptionWithPrefix(options, "-Xmx"))
	assert.NotEmpty(t, FindOptionWithPrefix(options, "-Xms"))

	options = rule.ConvertOptions(&Context{JdkVersion: "8", MemoryLimit: 1024}, []string{"-version"}, "quota*2/3,xmx-2")
	assert.NotEmpty(t, FindOptionWithPrefix(options, "-Xmx"))
	assert.NotEmpty(t, FindOptionWithPrefix(options, "-Xms"))

	options = rule.ConvertOptions(&Context{JdkVersion: "8", MemoryLimit: 1024}, []string{"-version"}, "quota*2/3,auto")
	assert.NotEmpty(t, FindOptionWithPrefix(options, "-Xmx"))
	assert.Empty(t, FindOptionWithPrefix(options, "-Xms"))
}
