package main

import (
	"github.com/leyantech/jstart/jstart"
)

func main() {
	jstart.Run(jstart.NewRuleRegistry(), &jstart.EnvLoader{})
}
