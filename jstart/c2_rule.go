package jstart

import (
	"strings"
)

// C2Rule enable/disable jit c2 compilation
type C2Rule struct {
	Quota int64
}

func (c *C2Rule) Name() string {
	return "c2"
}

func (c *C2Rule) ConvertOptions(jdkVersion string, originalOptions []string, ruleParam string) []string {
	ruleParam = strings.TrimSpace(ruleParam)
	var level string
	if ruleParam == "0" || ruleParam == "1" || ruleParam == "2" || ruleParam == "3" || ruleParam == "4" {
		level = ruleParam
	} else if ruleParam == "on" {
		level = "4"
	} else if ruleParam == "off" {
		level = "1"
	} else if ruleParam == "auto" {
		var limit int64
		if c.Quota > 0 {
			limit = c.Quota
		} else {
			var err error
			limit, err = getMemoryLimit()
			if err != nil {
				return originalOptions
			}
		}
		if limit > 300 { // 300MB
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
