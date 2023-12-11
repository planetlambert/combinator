package combinator

import (
	"slices"
)

type (
	Combinator struct {
		Name       string
		Arguments  []string
		Definition string
	}

	Basis []Combinator
)

// SK and SKI
var (
	S = Combinator{
		Name:       "S",
		Arguments:  []string{"x", "y", "z"},
		Definition: "xz(yz)",
	}

	K = Combinator{
		Name:       "K",
		Arguments:  []string{"x", "y"},
		Definition: "x",
	}

	I = Combinator{
		Name:       "I",
		Arguments:  []string{"x"},
		Definition: "x",
	}

	SK  = Basis{S, K}
	SKI = Basis{S, K, I}
)

// BCKW
var (
	B = Combinator{
		Name:       "B",
		Arguments:  []string{"x", "y", "z"},
		Definition: "x(yz)",
	}

	C = Combinator{
		Name:       "C",
		Arguments:  []string{"x", "y", "z"},
		Definition: "xzy",
	}

	W = Combinator{
		Name:       "C",
		Arguments:  []string{"x", "y"},
		Definition: "xyy",
	}

	BCKW = Basis{B, C, K, W}
)

// Iota
var (
	Iota = Basis{
		S,
		K,
		Combinator{
			Name:       "i",
			Arguments:  []string{"x"},
			Definition: "xSK",
		},
	}
)

func (b Basis) Transform(statement string) (string, error) {
	if err := isWellDefined(statement); err != nil {
		return "", err
	}
	tree := parse(statement)
	reducedTree := reduce(tree, b, false)
	return unparse(reducedTree), nil
}

func findCombinator(name string, b Basis) (Combinator, bool) {
	index := slices.IndexFunc(b, func(c Combinator) bool {
		return c.Name == name
	})
	if index >= 0 {
		return b[index], true
	} else {
		return Combinator{}, false
	}
}
