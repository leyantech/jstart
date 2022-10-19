package jstart

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func Run(registry *RuleRegistry, settingsLoader RuleParamsLoader) {
	realJavaCommand := findRealJavaExecutable(os.Args[0])
	jdkMajorVersion := getJdkMajorVersion(realJavaCommand)
	INFO.Printf("detected major jdk version %s", jdkMajorVersion)
	javaOptions, applicationArguments := splitJavaOptionsAndArguments(os.Args[1:])
	DEBUG.Printf("raw java options: %s, raw application arguments: %s", javaOptions, applicationArguments)
	ruleParams := settingsLoader.Load()
	INFO.Printf("jstart rule parameters: %v", ruleParams)

	isProd := os.Getenv("ENV_PRIORITY") != "test"
	limit, err := getMemoryLimit()
	if err != nil {
		ERROR.Printf("failed to detect memory limit: %s", err)
	}
	context := NewContextBuilder().JdkVersion(jdkMajorVersion).MemoryLimit(limit).IsProd(isProd).Build()

	for _, name := range OrderedKeys(ruleParams) {
		setting := ruleParams[name]
		rule, exist := registry.Find(name)
		if !exist {
			WARN.Printf("unknown jstart rule %s", name)
			continue
		}
		DEBUG.Printf("options before apply rule %s : %v", name, javaOptions)
		javaOptions = rule.ConvertOptions(context, javaOptions, setting)
		DEBUG.Printf("options after apply rule %s : %v", name, javaOptions)
	}
	finalCliArguments := append(javaOptions, applicationArguments...)
	INFO.Printf("final java cli arguments are %v", finalCliArguments)
	if os.Getenv("JSTART_DRY_RUN") != "" {
		fmt.Printf("%s %s\n", realJavaCommand, strings.Join(finalCliArguments, " "))
	} else {

		err := syscall.Exec(realJavaCommand, append([]string{"java"}, finalCliArguments...), syscall.Environ())
		if err != nil {
			ERROR.Fatalf("failed to execute the final java command: %s", err)
		}
	}
}

func findRealJavaExecutable(currentCommand string) string {
	//  find the real java executable, in case the current program is alias to java
	if filepath.Base(currentCommand) == "java" {
		pathList := filepath.SplitList(os.Getenv("PATH"))
		newPathList := make([]string, 0)
		hijackingRemoved := false
		for _, path := range pathList {
			maybeHijackingJava := filepath.Join(path, "java")
			stat, err := os.Lstat(maybeHijackingJava)
			// remove the hijacking of java command
			if !hijackingRemoved && err == nil && !stat.IsDir() && (stat.Mode()&os.ModeSymlink != 0) {
				DEBUG.Printf("removing %s from PATH", path)
				hijackingRemoved = true
			} else {
				newPathList = append(newPathList, path)
			}
		}
		newPathEnv := strings.Join(newPathList, string(os.PathListSeparator))
		err := os.Setenv("PATH", newPathEnv)
		if err != nil {
			WARN.Printf("failed to reset PATH environment variable: %s", err)
		}
	}
	path, err := exec.LookPath("java")
	if err != nil {
		ERROR.Fatalf("can not find command java")
	}
	return path
}

/**
get jdk major version(8, 9, 10, 11, etc)

return "unknown" for unknown jdk version
*/
func getJdkMajorVersion(javaExecutable string) string {
	command := exec.Command(javaExecutable, "-version")
	output, err := command.CombinedOutput()
	if err != nil {
		ERROR.Printf("failed to execute java -version: %s", err)
		return "unknown"
	} else {
		for _, line := range strings.Split(string(output), "\n") {
			line = strings.ToLower(line)
			leftQuoteIndex := strings.IndexByte(line, '"')
			rightQuoteIndex := strings.LastIndexByte(line, '"')
			if leftQuoteIndex == -1 || rightQuoteIndex == -1 {
				continue
			}
			fullVersion := line[leftQuoteIndex+1 : rightQuoteIndex]
			parts := strings.Split(fullVersion, ".")
			if len(parts) < 2 {
				continue
			}
			if parts[0] == "1" && parts[1] == "8" {
				return "8"
			}
			versionInt, err := strconv.Atoi(parts[0])
			if err == nil && versionInt > 8 {
				return parts[0]
			}
		}
	}
	return "unknown"
}

func splitJavaOptionsAndArguments(cliArgs []string) ([]string, []string) {
	// ref: https://docs.oracle.com/javase/8/docs/technotes/tools/windows/java.html
	// there are two modes of executing a java program
	// 1. execute a class: java [-options] class [args...]
	// 2. execute a jar file: java [-options] -jar jarfile [args...]

	var i int
	for i = 0; i < len(cliArgs); i++ {
		arg := cliArgs[i]
		if arg == "-cp" || arg == "-classpath" {
			i++
			continue
		}
		// NOTE: we'll leave the check of coexistence of -jar and main class name to jvm
		if !strings.HasPrefix(arg, "-") || arg == "-jar" || arg == "-h" || arg == "-version" || arg == "-help" {
			break
		}
	}
	if i == len(cliArgs) {
		WARN.Printf("no main class nor -jar option supplied")
	}
	// copy slice into new arrays to avoid underlying array corrupt
	return append([]string{}, cliArgs[:i]...), append([]string{}, cliArgs[i:]...)
}
