package model

import (
	"regexp"
	"encoding/json"
)

type operation int

const (
	NOT operation = iota
	AND
	OR
)

type Rule interface {
	evaluate(event string, context string) bool
}

type basicRule struct {
	pattern string
}

func newBasicRule(pattern string) basicRule {
	return basicRule {
		pattern: pattern,
	}
}

func (rule basicRule) evaluate(event string, context string) bool {
	regex := regexp.MustCompile(rule.pattern)
	return len(regex.FindAllStringIndex(event, -1)) > 0
}

type compositeRule struct {
	op    operation
	rules []Rule
}

func newCompositeRule(op operation, rules []Rule) compositeRule {
	return compositeRule {
		op: op,
		rules: rules,
	}	
}

func (rule compositeRule) evaluate(event string, context string) bool {
	switch rule.op {
	case NOT:
		for _, element := range rule.rules {
			if element.evaluate(event, context) {
				return false
			}
		}
		return true
	case AND:
		for _, element := range rule.rules {
			if !element.evaluate(event, context) {
				return false
			}
		}
		return true
	case OR:
		for _, element := range rule.rules {
			if element.evaluate(event, context) {
				return true
			}
		}
		return false
	}
	return false
}

//Warning. Here be disgusting json parsing.
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
