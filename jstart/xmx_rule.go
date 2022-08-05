package jstart

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
	"strings"
)

var (
	minimumXmx     int64 = 64
	minimumNonHeap int64 = 64
)

// XmxRule setting Xmx and Xms via an arithmetic expression.
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

	var xmxExpression, xmsExpression string
	parts := strings.Split(ruleParam, ",")
	if len(parts) == 0 {
		ERROR.Printf("no xmx/xms rule parameter found.")
		return originalOptions
	} else if len(parts) == 1 {
		xmxExpression = parts[0]
		xmsExpression = "xmx"
	} else if len(parts) ==2 {
		xmxExpression = parts[0]
		xmsExpression = parts[1]
	} else {
		ERROR.Printf("too many comma separated parts in xmx rule setting")
		return originalOptions
	}

	xmxEvaluateVariables := make(map[string]int64)
	xmxEvaluateVariables["quota"] = memoryLimit
	xmx, err := evaluateArithmeticExpression(xmxExpression, minimumXmx, memoryLimit-minimumNonHeap, xmxEvaluateVariables)
	if err != nil {
		return originalOptions
	}

	if xmsExpression == "auto" {
		return append(RemoveOptionsWithPrefix(originalOptions, "-Xmx"), fmt.Sprintf("-Xmx%dm", xmx))
	}

	xmsEvaluateVariables := make(map[string]int64)
	xmsEvaluateVariables["xmx"] = xmx

	// -Xms should be greater than 1MB, according to: https://docs.oracle.com/en/java/javase/13/docs/specs/man/java.html
	xms, err := evaluateArithmeticExpression(xmsExpression, 1, xmx, xmsEvaluateVariables)
	if err != nil {
		ERROR.Printf("failed to evaluate xms: %s, will not insert -Xms option", err)
		return append(RemoveOptionsWithPrefix(originalOptions, "-Xmx"), fmt.Sprintf("-Xmx%dm", xmx))
	} else {
		return append(RemoveOptionsWithPrefix(originalOptions, "-Xmx", "-Xms"),
			fmt.Sprintf("-Xmx%dm", xmx),
			fmt.Sprintf("-Xms%dm", xms),
		)
	}
}

func evaluateArithmeticExpression(expressionString string, minimum, maximum int64, variables map[string]int64) (int64, error) {
	expression, err := govaluate.NewEvaluableExpression(expressionString)
	if err != nil {
		ERROR.Printf("not valid expression for result: %s", err)
		return 0, err
	}
	params := make(map[string]interface{})
	for k, v := range variables {
		params[k] = v
	}
	evaluated, err := expression.Evaluate(params)
	if err != nil {
		ERROR.Printf("failed to evaluate expression %s : %s, wont add any result options", expressionString, err)
		return 0, err
	}

	var result int64
	switch v := evaluated.(type) {
	case float32:
		result = int64(v)
	case float64:
		result = int64(v)
	case int:
		result = int64(v)
	case uint64:
		result = int64(v)
	case uint32:
		result = int64(v)
	case int64:
		result = v
	default:
		ERROR.Printf("invalid result evaluation evaluated type: %s(%v)", reflect.TypeOf(evaluated), evaluated)
		return 0, fmt.Errorf("invalid evaluated type")
	}
	if result < minimum {
		ERROR.Printf("too small result: %d", result)
		return 0, fmt.Errorf("result too small")
	} else if result > maximum {
		ERROR.Printf("too large result: %d, which will lead few space for non heap usage: %d", result, maximum)
		return 0, fmt.Errorf("result too large")
	}
	INFO.Printf("%s evaluated to %dMB with %v", expressionString, result, variables)
	return result, nil
}
