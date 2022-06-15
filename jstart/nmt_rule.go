package jstart

type NmtRule struct{}

func (NmtRule) Name() string {
	return "nmt"
}

func (NmtRule) ConvertOptions(jdkVersion string, originalOptions []string, ruleParam string) []string {
	result := RemoveOptionsWithPrefix(originalOptions, "-XX:NativeMemoryTracking")
	if ruleParam == "off" {
		return append(result, "-XX:NativeMemoryTracking=off")
	} else if ruleParam == "summary" || ruleParam == "on" {
		return append(result, "-XX:NativeMemoryTracking=summary")
	} else if ruleParam == "detail" {
		return append(result, "-XX:NativeMemoryTracking=detail")
	} else {
		WARN.Printf("unexpected param for rule nmt: %s", ruleParam)
		return result
	}
}
