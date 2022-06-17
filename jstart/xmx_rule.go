package jstart

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
)

var (
	minimumXmx     int64 = 64
	minimumNonHeap int64 = 64
)

// a jstart rule for setting Xmx by multiple total memory limit with specific ratio
type XmxRule struct {
}

func (r *XmxRule) Name() string {
	return "xmx"
}

func (r *XmxRule) ConvertOptions(jdkVersion string, originalOptions []string, ruleParam string) []string {
	memoryLimit, err := getMemoryLimit()
	if err != nil {
		ERROR.Printf("failed to detect memory limit, wont add any xmx options: %s", err)
		return originalOptions
	}
	xmx, err := evaluateXmxExpression(memoryLimit, ruleParam)
	if err != nil {
		return originalOptions
	} else {
		return append(RemoveOptionsWithPrefix(originalOptions, "-Xmx", "-Xms"),
			fmt.Sprintf("-Xmx%dm", xmx),
			fmt.Sprintf("-Xms%dm", xmx),
		)
	}
}

func evaluateXmxExpression(memoryLimit int64, expressionString string) (int64, error) {
	expression, err := govaluate.NewEvaluableExpression(expressionString)
	if err != nil {
		ERROR.Printf("not valid expression for xmx: %s", err)
		return 0, err
	}
	params := map[string]interface{}{
		"quota": memoryLimit,
	}
	result, err := expression.Evaluate(params)
	if err != nil {
		ERROR.Printf("failed to evaluate expression %s : %s, wont add any xmx options", expressionString, err)
		return 0, err
	}

	var xmx int64
	switch v := result.(type) {
	case float32:
		xmx = int64(v)
	case float64:
		xmx = int64(v)
	case int:
		xmx = int64(v)
	case uint64:
		xmx = int64(v)
	case uint32:
		xmx = int64(v)
	case int64:
		xmx = v
	default:
		ERROR.Printf("invalid xmx evaluation result type: %s(%v)", reflect.TypeOf(result), result)
		return 0, fmt.Errorf("invalid xmx evaluation result type")
	}
	if xmx < minimumXmx {
		ERROR.Printf("too small xmx: %d", xmx)
		return 0, fmt.Errorf("xmx too small")
	} else if xmx >= (memoryLimit - minimumNonHeap) {
		ERROR.Printf("too large xmx: %d, which will lead few space for non heap usage: %d", xmx, memoryLimit-xmx)
		return 0, fmt.Errorf("xmx too large")
	}
	INFO.Printf("xmx rule %s evaluated to %dMB", expressionString, xmx)
	return xmx, nil
}
