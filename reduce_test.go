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
		"Ix":   "x",
		"Rxy":  "yx",
		"Ryx":  "xy",
		"RIxy": "xIy",
		"II":   "I",
	}
	for statement, expectedResult := range tests {
		t.Run(statement, func(t *testing.T) {
			tree := parse(statement)
			actualResult := unparse(reduce(tree, basis, false))
			if expectedResult != actualResult {
				t.Errorf("parsed statement %s incorrectly, expected %s but got %s", statement, expectedResult, actualResult)
			}
		})
	}
}
