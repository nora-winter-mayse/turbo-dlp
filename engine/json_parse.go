package engine

import (
	"encoding/json"
)

// Warning. Here be disgusting json parsing.
func buildRuleFromJson(input string) Rule {
        var rawResult map[string]interface{}
        json.Unmarshal([]byte(input), &rawResult)
        return buildRuleFromUnmarshalledObject(rawResult)
}

func buildRuleFromUnmarshalledObject(rawResult map[string]interface{}) Rule {
        _, ok := rawResult["pattern"]
        if ok {
                return newBasicRule(rawResult["pattern"].(string))
        }
        rules := []Rule{}
        for _, value := range rawResult["rules"].([]interface{}) {
                rules = append(rules, buildRuleFromUnmarshalledObject(value.(map[string]interface{})))
        }
        op := NOT
        switch rawResult["operation"].(string) {
        case "not":
                op = NOT
        case "and":
                op = AND
        case "or":
                op = OR
        }
        return newCompositeRule(op, rules)
}

