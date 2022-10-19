package jstart

import (
	"fmt"
	"strconv"
)

// a jstart rule for setting -XX:MaxMetaspaceSize
type MaxMetaSpaceRule struct{}

func (m *MaxMetaSpaceRule) Name() string {
	return "metaspace"
}

func (m *MaxMetaSpaceRule) ConvertOptions(context Context, originalOptions []string, ruleParam string) []string {
	maxMetaSpaceMB, err := strconv.Atoi(ruleParam)
	if err != nil {
		ERROR.Printf("can not parse %s to int", ruleParam)
		return originalOptions
	}
	if maxMetaSpaceMB <= 0 {
		ERROR.Printf("invalid metaspace size %d", maxMetaSpaceMB)
		return originalOptions
	}
	return append(RemoveOptionsWithPrefix(originalOptions, "-XX:MaxMetaspaceSize="),
		fmt.Sprintf("-XX:MaxMetaspaceSize=%dm", maxMetaSpaceMB))
}
