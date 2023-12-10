package combinator

import (
	"testing"
)

func TestReduce(t *testing.T) {
	basis := Basis{
		Combinator{
			Name:       "I",
			Arguments:  []string{"x"},
			Definition: "x",
		},
		Combinator{
			Name:       "R",
			Arguments:  []string{"x", "y"},
			Definition: "yx",
		},
	}

	tests := map[string]string{
		"Ix":        "x",
		"Rxy":       "yx",
		"Ryx":       "xy",
		"RIxy":      "xIy",
		"II":        "I",
		"x(Ix)":     "xx",
		"x(I(Ix))x": "xxx",
	}
	for statement, expectedResult := range tests {
		t.Run(statement, func(t *testing.T) {
			tree := parse(statement)
			actualResultNormal := unparse(reduce(tree, basis, false))
			actualResultApplicative := unparse(reduce(tree, basis, true))
			if expectedResult != actualResultNormal {
				t.Errorf("parsed statement %s incorrectly with normal order, expected %s but got %s", statement, expectedResult, actualResultNormal)
			}
			if expectedResult != actualResultApplicative {
				t.Errorf("parsed statement %s incorrectly with applicative order, expected %s but got %s", statement, expectedResult, actualResultApplicative)
			}
		})
	}
}
