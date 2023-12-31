package combinator

import (
	"fmt"
	"testing"
)

func TestSchonfinkel(t *testing.T) {
	tests := map[string]string{
		"SKKx":          "Ix",
		"SKyx":          "Ix",
		"S(KS)Kxyz":     "Zxyz",
		"S(ZZS)(KK)xyz": "Txyz",
	}
	for statement1, statement2 := range tests {
		t.Run(fmt.Sprintf("%s = %s", statement1, statement2), func(t *testing.T) {
			result1, _ := Schonfinkel.Transform(statement1)
			result2, _ := Schonfinkel.Transform(statement2)
			if result1 != result2 {
				t.Errorf("combinators %s and %s not equal, got %s and %s", statement1, statement2, result1, result2)
			}
		})
	}
}

func TestSKI(t *testing.T) {
	tests := map[string]string{
		"Ix":          "x",
		"SKSK":        "K",
		"SKKx":        "x",
		"S(K(SI))Kab": "ba",
	}
	for statement, expectedResult := range tests {
		t.Run(statement, func(t *testing.T) {
			actualResult, _ := SKI.Transform(statement)
			if expectedResult != actualResult {
				t.Errorf("transformed SKI statement %s incorrectly, expected %s but got %s", statement, expectedResult, actualResult)
			}
		})
	}
}

func TestIota(t *testing.T) {
	tests := map[string]string{
		"iix":           "x",
		"(i(i(ii)))":    "K",
		"(i(i(i(ii))))": "S",
	}
	for statement, expectedResult := range tests {
		t.Run(statement, func(t *testing.T) {
			actualResult, _ := Iota.Transform(statement)
			if expectedResult != actualResult {
				t.Errorf("transformed Iota statement %s incorrectly, expected %s but got %s", statement, expectedResult, actualResult)
			}
		})
	}
}

func TestChurch(t *testing.T) {
	tests := map[string]string{
		"0fx":                      "x",
		"S(0)fx":                   "fx",
		"S(S(0))fx":                "f(fx)",
		"P(0)(0)fx":                "x",
		"P(S(S(S(0))))(S(S(0)))fx": "f(f(f(f(fx))))",
		"M(S(S(S(0))))(S(S(0)))fx": "f(f(f(f(f(fx)))))",
		"E(S(S(S(0))))(S(S(0)))fx": "f(f(f(f(f(f(f(f(fx))))))))",
	}
	for statement, expectedResult := range tests {
		t.Run(statement, func(t *testing.T) {
			actualResult, _ := Church.Transform(statement)
			if expectedResult != actualResult {
				t.Errorf("transformed Church statement %s incorrectly, expected %s but got %s", statement, expectedResult, actualResult)
			}
		})
	}
}
