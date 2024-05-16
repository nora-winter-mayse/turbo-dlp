package engine

import (
	"testing"
)

func TestBuildBasicRuleFromJson(t *testing.T) {
        input := "{\"pattern\": \"[0-9]\"}"
        genericRule := buildRuleFromJson(input)
        rule, ok := genericRule.(basicRule)
        if !ok {
                t.Fatal("TestBuildBasicRuleFromJson FAILED - result could not be demarshalled into basicRule")
        }
        if rule.pattern != "[0-9]" {
                t.Fatal("TesstBuildBasicRuleFromJson FAILED - pattern not correct on resulting basicRule")
        }
}

func TestBuildCompositeRuleFromJson(t *testing.T) {
        input := "{\"operation\": \"and\",\"rules\": [{\"pattern\": \"[0-9]\"}, {\"pattern\": \"[a-z]\"}]}"
        genericRule := buildRuleFromJson(input)
        rule, ok := genericRule.(compositeRule)
        if !ok {
                t.Fatal("TestBuildCompositeRuleFromJson FAILED - result could not be demarshalled into compositeRule")
        }
        if rule.op != AND {
                t.Fatal("TestBuildCompositeRuleFromJson FAILED - operation not correct on resulting compositeRule")
        }
        if rule.rules[0].(basicRule).pattern != "[0-9]" {
                t.Fatal("TestBuildCompositeRuleFromJson FAILED - pattern not correct on first basicRule")
        }
        if rule.rules[1].(basicRule).pattern != "[a-z]" {
                t.Fatal("TestBuildCompositeRuleFromJson FAILED - pattern not correct on second basicRule")
        }
}

