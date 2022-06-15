package jstart

import "os"

type RemoteDebugRule struct{}

func (*RemoteDebugRule) Name() string {
	return "remote_debug"
}

func isInTestCluster() bool {
	return os.Getenv("ENV_PRIORITY") == "test"
}

func (*RemoteDebugRule) ConvertOptions(jdkVersion string, originalOptions []string, ruleParam string) []string {
	// remove jdwp related options provided in original options
	result := RemoveOptionsWithPrefix(originalOptions, "-agentlib:jdwp=")
	enableRemoteDebug := ruleParam == "on" || (ruleParam == "auto" && isInTestCluster())
	if enableRemoteDebug {
		debugOptions := []string{"-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=3330"}
		return append(debugOptions, result...)
	} else {
		return result
	}
}
