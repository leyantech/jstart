package jstart

type GcLogRule struct{}

func (g *GcLogRule) Name() string {
	return "gc_log"
}

func (g *GcLogRule) ConvertOptions(jdkVersion string, originalOptions []string, ruleParam string) []string {
	if jdkVersion == "8" {
		// ref: https://dzone.com/articles/enabling-and-analysing-the-garbage-collection-log
		gcLogOptions := []string{
			// always print out gc logs
			"-verbose:gc",
			// always create a heapdump on OutOfMemoryError
			"-XX:+HeapDumpOnOutOfMemoryError",
			"-XX:+PrintGCDetails",
			"-XX:+PrintGCDateStamps",
		}
		if ruleParam == "file" {
			gcLogOptions = append(gcLogOptions,
				"-Xloggc:./gc.log",
				"-XX:NumberOfGCLogFiles=10",
				"-XX:GCLogFileSize=10m",
				"-XX:+UseGCLogFileRotation")
		}
		return append(originalOptions, gcLogOptions...)
	} else {
		// ref: https://stackoverflow.com/questions/54144713/is-there-a-replacement-for-the-garbage-collection-jvm-args-in-java-11
		gcLogOption := "-Xlog:gc*" // PrintGCDetails
		if ruleParam == "file" {
			gcLogOption = gcLogOption + ":file=gc.log::filecount=10,filesize=10M"
		} else {
			gcLogOption = gcLogOption + ":stdout"
		}
		gcLogOption = gcLogOption + ":time,level,tags"
		return append(originalOptions, "-XX:+HeapDumpOnOutOfMemoryError", gcLogOption)
	}
}
