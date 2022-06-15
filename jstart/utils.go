package jstart

import (
	"github.com/cloudfoundry/gosigar"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var (
	// UnlimitedMemorySize defines the bytes size when memory limit is not set (2 ^ 63 - 4096)
	UnlimitedMemorySize = "9223372036854771712"
)

func RemoveOptionsWithPrefix(options []string, prefixList ...string) []string {
	result := make([]string, 0)
	for _, option := range options {
		hasPrefix := false
		for _, prefix := range prefixList {
			if strings.HasPrefix(option, prefix) {
				hasPrefix = true
				break
			}
		}
		if !hasPrefix {
			result = append(result, option)
		}
	}
	return result
}

func FindOptionWithPrefix(options []string, prefix string) string {
	for _, option := range options {
		if strings.HasPrefix(option, prefix) {
			return option
		}
	}
	return ""
}

// get memory limit(in mega bytes) for the current runtime.
func getMemoryLimit() (int64, error) {
	// NOTE: gosigar not working properly inside a docker container.
	file, err := ioutil.ReadFile("/sys/fs/cgroup/memory/memory.limit_in_bytes")
	if err == nil {
		limit := strings.TrimSpace(string(file))
		if limit != UnlimitedMemorySize {
			limitBytes, err := strconv.ParseInt(limit, 10, 64)
			if err != nil {
				return 0, err
			} else {
				return limitBytes >> 20, nil
			}
		}
	}
	mem := &sigar.Mem{}
	err = mem.Get()
	if err != nil {
		ERROR.Printf("failed to detect memory limit %s", err)
		return 0, err
	} else {
		// convert to mega bytes
		return int64(mem.Total >> 20), err
	}
}

// NOTE: this hack avoids random iterate order on golang maps
func OrderedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
