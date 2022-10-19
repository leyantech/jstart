package jstart

import "strings"

type LiteralJavaOptionsRule struct{}

func (r *LiteralJavaOptionsRule) Name() string {
	return "literal"
}

func (r *LiteralJavaOptionsRule) ConvertOptions(context *Context, originalOptions []string, ruleParam string) []string {
	fields := strings.Fields(ruleParam)
	return append(originalOptions, fields...)
}
