package jstart

type Rule interface {
	Name() string
	ConvertOptions(context *Context, originalOptions []string, ruleParam string) []string
}
