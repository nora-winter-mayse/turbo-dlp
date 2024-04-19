package model

import (
	"testing"
)

func TestEvaluateDictionary(t *testing.T) {
	input := "If I say cat two more times, there will be three cats in this kitty-cat rhyme! - Fred Durst"
	evaluation := DictEvaluation{[]string{"cat"}, "CatRule", "ID"}
	actual := evaluation.evaluate(input, "")
	if actual != 3 {
		t.Fatalf("Failure: 'cat' appears 3 times, found %d times.", actual)
	}
}

func TestEvaluateRegex(t *testing.T) {
	input := "If I say cat two more times, there will be three cats in this kitty-cat rhyme! - Fred Durst"
	regex_evaluation := RegexEvaluation{"cat", "CatRegexRule", "ID"}
	actual := regex_evaluation.evaluate(input, "")
	if actual != 3 {
		t.Fatalf("Failure: 'cat' appears 3 times, found %d times.", actual)
	}
}
