package combinator

import "testing"

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
				t.Errorf("transformed SKI statement %s incorrectly, expected %s but got %s", statement, expectedResult, actualResult)
			}
		})
	}
}
