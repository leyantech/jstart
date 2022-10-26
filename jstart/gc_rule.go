package jstart

import "regexp"

type GcRule struct{}

func (GcRule) Name() string {
	return "gc"
}

func (GcRule) ConvertOptions(context Context, originalOptions []string, ruleParam string) []string {
	if isIncludeGcParam(originalOptions) {
		WARN.Printf("java original option already include gc param...")
		return originalOptions
	}
	var gcOptions []string
	memLimit := context.GetMemoryLimit()
	if memLimit <= 0 {
		return originalOptions
	}
	if ruleParam == "auto" {
		gcOptions = detectGcForMemLimit(memLimit)
	} else if ruleParam == "serial" {
		gcOptions = serialGcOptions(memLimit)
	} else if ruleParam == "parallel" {
		gcOptions = parallelGcOptions(memLimit)
	} else if ruleParam == "cms" {
		gcOptions = cmsGcOptions(memLimit)
	} else if ruleParam == "g1" {
		gcOptions = g1GcOptions(memLimit)
	} else if ruleParam == "shenandoah" {
		gcOptions = shenandoahGcOptions(memLimit)
	} else {
		WARN.Printf("unknown gc rule param %s", ruleParam)
		return originalOptions
	}
	// leave conflict detecting to raw java command
	return append(gcOptions, originalOptions...)
}

func isIncludeGcParam(javaOptions []string) bool {
	for _, option := range javaOptions {
		matched, _ := regexp.Match("^-XX:+Use[A-Za-z]+GC$", []byte(option))
		if matched {
			return true
		}
	}
	return false
}

func detectGcForMemLimit(memLimit int64) []string {
	if memLimit <= 1024 {
		return serialGcOptions(memLimit)
	} else if memLimit <= 2048 {
		return cmsGcOptions(memLimit)
	} else {
		return g1GcOptions(memLimit)
	}
}

func serialGcOptions(memLimit int64) []string {
	return []string{"-XX:+UseSerialGC"}
}

func parallelGcOptions(memLimit int64) []string {
	return []string{"-XX:+UseParallelGC"}
}

func cmsGcOptions(memLimit int64) []string {
	return []string{"-XX:+UseConcMarkSweepGC"}
}

func g1GcOptions(memLimit int64) []string {
	return []string{"-XX:+UseG1GC"}
}

func shenandoahGcOptions(memLimit int64) []string {
	return []string{"-XX:+UseShenandoahGC"}
}
