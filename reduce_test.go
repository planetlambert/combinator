package combinator

import (
	"context"
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
			reducedNormal, _ := reduce(context.Background(), tree, basis, false, 0)
			actualResultNormal := unparse(reducedNormal)
			reducedApplicative, _ := reduce(context.Background(), tree, basis, true, 0)
			actualResultApplicative := unparse(reducedApplicative)
			if expectedResult != actualResultNormal {
				t.Errorf("parsed statement %s incorrectly with normal order, expected %s but got %s", statement, expectedResult, actualResultNormal)
			}
			if expectedResult != actualResultApplicative {
				t.Errorf("parsed statement %s incorrectly with applicative order, expected %s but got %s", statement, expectedResult, actualResultApplicative)
			}
		})
	}
}
