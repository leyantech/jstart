package jstart

import (
	"strings"
)

// C2Rule enable/disable jit c2 compilation
type C2Rule struct {
}

func (c *C2Rule) Name() string {
	return "c2"
}

func (c *C2Rule) ConvertOptions(jdkVersion string, originalOptions []string, ruleParam string) []string {
	ruleParam = strings.TrimSpace(ruleParam)
	level := "1"
	if ruleParam == "0" || ruleParam == "1" || ruleParam == "2" || ruleParam == "3" || ruleParam == "4" {
		level = ruleParam
	} else if ruleParam == "on" {
		level = "4"
	} else if ruleParam == "off" {
		level = "1"
	} else if ruleParam == "auto" {
		limit, err := getMemoryLimit()
		if err != nil {
			return originalOptions
		}
		if limit > 300 * 1024 * 1024 { // 300MB
			level = "4"
		} else {
			level = "1"
		}
	} else {
		return originalOptions
	}
	afterCleanUpOptions := RemoveOptionsWithPrefix(originalOptions, "-Xint", "-XX:+TieredCompilation", "-XX:-TieredCompilation", "-XX:TieredStopAtLevel")
	return append(afterCleanUpOptions, "-XX:+TieredCompilation", "-XX:TieredStopAtLevel=" + level)
}
