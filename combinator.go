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

// From section 3 of the paper
var (
	// Identity
	I = Combinator{
		Name:       "I",
		Arguments:  []string{"x"},
		Definition: "x",
	}

	// Constancy
	K = Combinator{
		Name:       "K",
		Arguments:  []string{"x", "y"},
		Definition: "x",
	}

	// Interchange
	T = Combinator{
		Name:       "T",
		Arguments:  []string{"x", "y", "z"},
		Definition: "xzy",
	}

	// Composition
	Z = Combinator{
		Name:       "Z",
		Arguments:  []string{"x", "y", "z"},
		Definition: "x(yz)",
	}

	// Fusion
	S = Combinator{
		Name:       "S",
		Arguments:  []string{"x", "y", "z"},
		Definition: "xz(yz)",
	}
)

// SK and SKI (https://en.wikipedia.org/wiki/SKI_combinator_calculus)
var (
	SK  = Basis{S, K}
	SKI = Basis{S, K, I}
)

// BCKW (https://en.wikipedia.org/wiki/B,_C,_K,_W_system)
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
			Name:      "i",
			Arguments: []string{"x"},
			// Note the use of other combinators in the definition
			// makes Iota "improper"
			Definition: "xSK",
		},
	}
)

// Transforms the statement using the Basis `b`
func (b Basis) Transform(statement string) (string, error) {
	if err := isWellDefined(statement); err != nil {
		return "", err
	}
	tree := parse(statement)
	reducedTree := reduce(tree, b, false)
	return unparse(reducedTree), nil
}

// Transforms the statement using the Combinator `c`
func (c Combinator) Transform(statement string) (string, error) {
	if err := isWellDefined(statement); err != nil {
		return "", err
	}
	tree := parse(statement)
	reducedTree := reduce(tree, Basis{c}, false)
	return unparse(reducedTree), nil
}

// Attempts to find a Combinator named `name` in the Basis `b`
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
