package model

import (
	"regexp"
	"strings"
)

type Evaluation interface {
	getId() string
	getName() string
	evaluate(event string, context string) int
}

type DictEvaluation struct {
	terms []string
	name  string
	id    string
}

func (e DictEvaluation) getName() string {
	return e.name
}

func (e DictEvaluation) evaluate(event string, context string) int {
	count := 0
	for _, element := range e.terms {
		count += strings.Count(event, element)
	}
	return count
}

type RegexEvaluation struct {
	pattern string
	name    string
	id      string
}

func (e RegexEvaluation) getName() string {
	return e.name
}

func (e RegexEvaluation) getId() string {
	return e.id
}

func (e RegexEvaluation) evaluate(event string, context string) int {
	regex := regexp.MustCompile(e.pattern)
	return len(regex.FindAllStringIndex(event, -1))
}
