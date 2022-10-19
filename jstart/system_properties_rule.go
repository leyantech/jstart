package jstart

import "fmt"

type PredefinedSystemPropertiesRule struct {
	name       string
	properties map[string]string
}

func CreatePredefinedSystemPropertiesRule(name string, properties map[string]string) *PredefinedSystemPropertiesRule {
	return &PredefinedSystemPropertiesRule{
		name:       name,
		properties: properties,
	}
}

func (r *PredefinedSystemPropertiesRule) Name() string {
	return r.name
}

func (r *PredefinedSystemPropertiesRule) ConvertOptions(context *Context, originalOptions []string, ruleParam string) []string {
	if ruleParam == "off" || ruleParam == "no" || ruleParam == "false" {
		return originalOptions
	}
	result := originalOptions
	for _, key := range OrderedKeys(r.properties) {
		value := r.properties[key]
		existed := FindOptionWithPrefix(originalOptions, fmt.Sprintf("-D=%s", key))
		optionToInsert := fmt.Sprintf("-D%s=%s", key, value)
		if existed != "" {
			WARN.Printf("system property exists: %s, will not insert: %s", existed, optionToInsert)
		} else {
			result = append(result, optionToInsert)
		}
	}
	return result
}
