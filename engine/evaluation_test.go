package engine

import (
	"testing"
)

func TestEvaluateBasicRule_whenPatternMatches_thenTrue(t *testing.T) {
	rule := newBasicRule("[0-9]")
	if !rule.evaluate("asdf3asjf", "") {
		t.Fatal("TestEvaluateBasicRule_whenPatternMatches_thenTrue FAILED - pattern should match but doesn't")
	}
}

func TestEvaluateBasicRule_whenPatternDoesntMatch_thenFalse(t *testing.T) {
	rule := newBasicRule("[0-9]")
	if rule.evaluate("asdfasdfaf", "") {
		t.Fatal("TestEvaluateBasicRule_whenPatternDoesntMatch_thenFalse FAILED - pattern shouldn't match but does")
	}
}

func TestEvaluateCompositeRule_whenAndDoesntMatch_thenFalse(t *testing.T) {
	subRuleOne := newBasicRule("[0-9]")
	subRuleTwo := newBasicRule("[a-z]")
	subRules := []Rule{subRuleOne, subRuleTwo}
	rule := newCompositeRule(AND, subRules)
	if rule.evaluate("ASEF7ASDF", "") {
		t.Fatal("TestEvaluateCompositeRule_whenAndDoesntMatch_thenFalse FAILED - pattern shouldn't match but does")
	}
}

func TestEvaluateCompositeRule_whenAndMatches_thenTrue(t *testing.T) {
	subRuleOne := newBasicRule("[0-9]")
	subRuleTwo := newBasicRule("[a-z]")
	subRules := []Rule{subRuleOne, subRuleTwo}
	rule := newCompositeRule(AND, subRules)
	if !rule.evaluate("aSEF7ASDF", "") {
		t.Fatal("TestEvaluateCompositeRule_whenAndDoesntMatch_thenFalse FAILED - pattern should match but doesn't")
	}
}

func TestEvaluateCompositeRule_whenOrDoesntMatch_thenFalse(t *testing.T) {
	subRuleOne := newBasicRule("[0-9]")
	subRuleTwo := newBasicRule("[a-z]")
	subRules := []Rule{subRuleOne, subRuleTwo}
	rule := newCompositeRule(OR, subRules)
	if rule.evaluate("ASDFASDF", "") {
		t.Fatal("TestEvaluateCompositeRule_whenOrDoesntMatch_thenFalse FAILED - pattern shouldn't match but does")
	}
}

func TestEvaluateCompositeRule_whenOrMatches_thenTrue(t *testing.T) {
	subRuleOne := newBasicRule("[0-9]")
	subRuleTwo := newBasicRule("[a-z]")
	subRules := []Rule{subRuleOne, subRuleTwo}
	rule := newCompositeRule(OR, subRules)
	if !rule.evaluate("asdfasdf", "") {
		t.Fatal("TestEvaluateCompositeRule_whenOrDoesntMatch_thenFalse FAILED - pattern should match but doesn't")
	}
}

func TestEvaluateCompositeRule_whenNotDoesntMatch_thenFalse(t *testing.T) {
	subRuleOne := newBasicRule("[0-9]")
	subRules := []Rule{subRuleOne}
	rule := newCompositeRule(NOT, subRules)
	if rule.evaluate("12341234", "") {
		t.Fatal("TestEvaluateCompositeRule_whenNotDoesntMatch_thenFalse FAILED - pattern shouldn't match but does")
	}
}

func TestEvaluateCompositeRule_whenNotMatches_thenTrue(t *testing.T) {
	subRuleOne := newBasicRule("[0-9]")
	subRules := []Rule{subRuleOne}
	rule := newCompositeRule(NOT, subRules)
	if !rule.evaluate("ASDFADF", "") {
		t.Fatal("TestEvaluateCompositeRule_whenNotDoesntMatch_thenFalse FAILED - pattern should match but doesn't")
	}
}

func TestEvaluateAAndNotB_Match(t *testing.T) {
	basicRuleOne := newBasicRule("[0-9]")
	basicRuleTwo := newBasicRule("[a-z]")
	notRule := newCompositeRule(NOT, []Rule{basicRuleOne})
	rule := newCompositeRule(AND, []Rule{notRule, basicRuleTwo})
	if !rule.evaluate("asIDdadNID", "") {
		t.Fatal("TestEvaluateAAndNotB_Match FAILED - pattern should match but doesn't")
	}
}

func TestEvaluateAAndNotB_NoMatch(t *testing.T) {
	basicRuleOne := newBasicRule("[0-9]")
	basicRuleTwo := newBasicRule("[a-z]")
	notRule := newCompositeRule(NOT, []Rule{basicRuleOne})
	rule := newCompositeRule(AND, []Rule{notRule, basicRuleTwo})
	if rule.evaluate("9asIDdadNID", "") {
		t.Fatal("TestEvaluateAAndNotB_Match FAILED - pattern shouldn't match but does")
	}
}
