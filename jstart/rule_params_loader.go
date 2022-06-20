package jstart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// RuleParams is a map of rule name to rule settings.
type RuleParams = map[string]string

type RuleParamsLoader interface {
	Load() RuleParams
}

func splitStringToMap(s string) map[string]string {
	entries := strings.Split(s, ";")

	m := make(map[string]string)
	for _, e := range entries {
		equalMarkIndex := strings.IndexRune(e, '=')
		if equalMarkIndex == -1 {
			WARN.Printf("%s is not valid jstart rule param syntax", e)
		} else {
			key := strings.TrimSpace(e[:equalMarkIndex])
			value := strings.TrimSpace(e[equalMarkIndex+1:])
			m[key] = value
		}
	}
	return m
}

/**
apollo(https://github.com/apolloconfig/apollo) is a configuration management system. This loader load rule parameters
from the application namespace of a given apollo app id.

rule parameters are encoded with format '<key>=<value>;<key>=<value>...'(e.g. 'xmx=quota-512;gc=auto;nmt=off').
 */
type ApolloLoader struct {
	ApolloUrl     string    // apollo's api endpoint
	ApolloAppId   string    // apollo app id
	PropertyNames []string   // property names to search for rule parameters, the first one with non-empty value will be effective.
}

type apolloResponse struct {
	Configurations map[string]string `json:"configurations"`
}

func (a *ApolloLoader) loadConfigMap() (map[string]string, error) {
	if a.ApolloAppId == "" || a.ApolloUrl == "" {
		return nil, fmt.Errorf("not in a container with access to apollo")
	}
	apiEndpoint := fmt.Sprintf("%s/configs/%s/default/application", a.ApolloUrl, a.ApolloAppId)
	httpResponse, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(httpResponse.Body)
	response := &apolloResponse{Configurations: make(map[string]string)}
	err = decoder.Decode(response)
	if err != nil {
		return nil, err
	}
	return response.Configurations, nil
}

func (a ApolloLoader) readJstartSettings(configMap map[string]string) RuleParams {
	for _, property := range a.PropertyNames {
		if value, ok := configMap[property]; ok {
			return splitStringToMap(value)
		}
	}
	WARN.Printf("no jstart settings found in apollo application namespace")
	return nil
}

/**
load rule parameters from apollo configuration management system.

homepage of apollo: https://github.com/apolloconfig/apollo
 */
func (a *ApolloLoader) Load() RuleParams {
	configMap, err := a.loadConfigMap()
	if err != nil {
		WARN.Printf("failed to load config from apollo: %s", err)
		return nil
	}
	return a.readJstartSettings(configMap)
}

type EnvLoader struct{}

func (e *EnvLoader) Load() RuleParams {
	if value, exist := os.LookupEnv("JSTART"); exist {
		return splitStringToMap(value)
	} else {
		return nil
	}
}

type WithStaticDefaultsLoader struct {
	delegate RuleParamsLoader
	defaults RuleParams
}

func AppendDefaults(loader RuleParamsLoader, defaults RuleParams) RuleParamsLoader {
	if defaults == nil {
		panic("default should not be nil")
	}
	return &WithStaticDefaultsLoader{
		delegate: loader,
		defaults: defaults,
	}
}

func (l *WithStaticDefaultsLoader) Load() RuleParams {
	params := l.delegate.Load()
	if params == nil {
		params = make(map[string]string)
	}
	for name, value := range l.defaults {
		if _, exist := params[name]; !exist {
			params[name] = value
		}
	}
	return params
}

type ChainedLoader struct {
	delegates []RuleParamsLoader
}

func NewChainedLoader(loader RuleParamsLoader, others ...RuleParamsLoader) *ChainedLoader {
	return &ChainedLoader{
		delegates: append([]RuleParamsLoader{loader}, others...),
	}
}

func (l *ChainedLoader) Load() RuleParams {
	for _, delegate := range l.delegates {
		ruleParams := delegate.Load()
		if ruleParams != nil {
			return ruleParams
		}
	}
	return nil
}
