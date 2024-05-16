package engine

import (
	"regexp"
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
	return basicRule{
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
	return compositeRule{
		op:    op,
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
