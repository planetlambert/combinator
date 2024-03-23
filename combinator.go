package combinator

import (
	"context"
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

	// All of Schönfinkel's defined combinators
	Schonfinkel = Basis{I, K, T, Z, S}
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

// Iota (https://en.wikipedia.org/wiki/Iota_and_Jot)
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

// Church Encoding (https://en.wikipedia.org/wiki/Church_encoding)
var (
	Zero = Combinator{
		Name:       "0",
		Arguments:  []string{"f", "x"},
		Definition: "x",
	}

	Succ = Combinator{
		Name:       "S",
		Arguments:  []string{"n", "f", "x"},
		Definition: "f(nfx)",
	}

	Plus = Combinator{
		Name:       "P",
		Arguments:  []string{"m", "n", "f", "x"},
		Definition: "mf(nfx)",
	}

	Mult = Combinator{
		Name:       "M",
		Arguments:  []string{"m", "n", "f", "x"},
		Definition: "m(nf)x",
	}

	Exp = Combinator{
		Name:       "E",
		Arguments:  []string{"m", "n", "f", "x"},
		Definition: "nmfx",
	}

	Church = Basis{Zero, Succ, Plus, Mult, Exp}
)

// Transforms the statement using the Basis `b`
func (b Basis) Transform(ctx context.Context, statement string) (string, error) {
	if err := isWellDefined(statement); err != nil {
		return "", err
	}
	tree := parse(statement)
	reducedTree, canceled := reduce(ctx, tree, b, false)
	if canceled {
		return "", ctx.Err()
	}
	return unparse(reducedTree), nil
}

// Adds an additional Combinator to the Basis
func (b Basis) With(combinator Combinator) Basis {
	return append(b, combinator)
}

// Transforms the statement using the Combinator `c`
func (c Combinator) Transform(ctx context.Context, statement string) (string, error) {
	if err := isWellDefined(statement); err != nil {
		return "", err
	}
	tree := parse(statement)
	reducedTree, canceled := reduce(ctx, tree, Basis{c}, false)
	if canceled {
		return "", ctx.Err()
	}
	return unparse(reducedTree), nil
}

// Returns whether the statement is well defined or not
func WellDefined(statement string) (bool, error) {
	err := isWellDefined(statement)
	return err == nil, err
}

// Parses the statement into a Tree
func Parse(statement string) *Tree {
	return parse(statement)
}

// Reduces the Tree using the Basis `b`
func (b Basis) Reduce(ctx context.Context, tree *Tree) (*Tree, error) {
	reduced, canceled := reduce(ctx, tree, b, false)
	if canceled {
		return nil, ctx.Err()
	}
	return reduced, nil
}

// Joins the two trees together under a new root
func Join(a *Tree, b *Tree) *Tree {
	return join(a, b)
}

// Produces a copy of the tree
func Copy(t *Tree) *Tree {
	return copy(t)
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
