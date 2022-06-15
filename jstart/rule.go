package jstart

type Rule interface {
	Name() string
	ConvertOptions(jdkVersion string, originalOptions []string, ruleParam string) []string
}
