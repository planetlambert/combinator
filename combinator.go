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

var I = Combinator{
	Name:       "I",
	Arguments:  []string{"x"},
	Definition: "x",
}

var K = Combinator{
	Name:       "K",
	Arguments:  []string{"x", "y"},
	Definition: "x",
}

var S = Combinator{
	Name:       "S",
	Arguments:  []string{"x", "y", "z"},
	Definition: "xz(yz)",
}

var SKI = Basis{S, K, I}

var Iota = Basis{
	S,
	K,
	Combinator{
		Name:       "i",
		Arguments:  []string{"x"},
		Definition: "xSK",
	},
}

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
