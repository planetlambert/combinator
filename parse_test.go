package combinator

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]string{
		"x":         "x",
		"(x)":       "x",
		"((x))":     "x",
		"xx":        "xx",
		"(x)x":      "xx",
		"((x)x)":    "xx",
		"((x)(x))":  "xx",
		"x(x)":      "xx",
		"xxx":       "xxx",
		"(xxx)":     "xxx",
		"((x)xx)":   "xxx",
		"(((x)x)x)": "xxx",
		"((x)(xx))": "x(xx)",
		"x(xx)":     "x(xx)",
		"x((x)x)":   "x(xx)",
		"(x(xx))":   "x(xx)",
		"xx(xx)":    "xx(xx)",
		"x(xx)x":    "x(xx)x",
	}
	for statement, expectedResult := range tests {
		t.Run(statement, func(t *testing.T) {
			tree := parse(statement)
			actualResult := unparse(tree)
			if expectedResult != actualResult {
				t.Errorf("parsed statement %s incorrectly, expected %s but got %s", statement, expectedResult, actualResult)
			}
		})
	}
}

func TestWellDefined(t *testing.T) {
	tests := map[string]bool{
		"x":        true,
		"(x)":      true,
		"(xy)":     true,
		"((x))":    true,
		"(x(x))":   true,
		"":         false,
		"()":       false,
		"x()":      false,
		"(()":      false,
		")(":       false,
		"(x(":      false,
		"(x))((x)": false,
	}
	for statement, expectedResult := range tests {
		t.Run(statement, func(t *testing.T) {
			err := isWellDefined(statement)
			actualResult := err == nil
			if expectedResult != actualResult {
				t.Errorf("tested well-definedness for statement %s incorrectly, expected %t but got %t", statement, expectedResult, actualResult)
			}
		})
	}
}
