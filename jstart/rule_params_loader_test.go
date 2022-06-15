package jstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApolloReadSettings(t *testing.T) {
	loader := &ApolloLoader{PropertyNames: []string{"jstart"}}
	m := map[string]string{}
	assert.Nil(t, loader.readJstartSettings(m))
	m = map[string]string{"jstart": "gc=cms;predefined_properties=true"}
	settings := loader.readJstartSettings(m)
	assert.NotEqual(t, settings, nil)
	assert.Equal(t, settings["gc"], "cms")
	assert.Equal(t, settings["predefined_properties"], "true")
}
