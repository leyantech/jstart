package jstart

type ModExportsRule struct{}

func (r *ModExportsRule) Name() string {
	return "mod_exports"
}

// export frequently used opens jdk's internal modules, to ease the adoption of newer jdk
// refs:
// - https://openjdk.java.net/jeps/261
// - https://openjdk.java.net/jeps/403
var modExportsOptions = []string{
	"--illegal-access=warn",
	"--add-opens", "java.base/java.nio=ALL-UNNAMED",
	"--add-opens", "java.base/java.lang=ALL-UNNAMED",
	"--add-opens", "java.base/java.lang.reflect=ALL-UNNAMED",
}

func (r *ModExportsRule) ConvertOptions(context *Context, originalOptions []string, ruleParam string) []string {
	// we only support jdk8+, so wont handle jdkVersion 7/6/5 here.
	if context.GetJdkVersion() != "8" {
		originalOptions = append(originalOptions, modExportsOptions...)
	}

	return originalOptions
}
