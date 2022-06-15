package jstart

import (
	"fmt"
	"strconv"
	"strings"
)

type DirectMemRule struct{}

func (d *DirectMemRule) Name() string {
	return "direct_mem"
}

func (d *DirectMemRule) ConvertOptions(jdkVersion string, originalOptions []string, ruleParam string) []string {
	ruleParam = strings.TrimSpace(ruleParam)
	if ruleParam == "" {
		return originalOptions
	}
	directMemLimitMB, err := strconv.Atoi(ruleParam)
	if err != nil {
		ERROR.Printf("can not parse %s to int", ruleParam)
		return originalOptions
	}
	if directMemLimitMB <= 0 {
		ERROR.Printf("invalid max direct memory limit %d", directMemLimitMB)
		return originalOptions
	}
	maxDirectMemOption := fmt.Sprintf("-XX:MaxDirectMemorySize=%dm", directMemLimitMB)
	return append(RemoveOptionsWithPrefix(originalOptions, "-XX:MaxDirectMemorySize="), maxDirectMemOption)
}
