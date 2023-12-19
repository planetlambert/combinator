package combinator

import (
	"slices"
)

func reduce(root *treeNode, b Basis, applicativeOrder bool) *treeNode {
	if root.IsLeaf {
		return root
	}

	newRoot := rewrite(root, b, applicativeOrder)

	if newRoot.IsLeaf {
		return newRoot
	}

	newRoot.Left = reduce(newRoot.Left, b, applicativeOrder)
	newRoot.Right = reduce(newRoot.Right, b, applicativeOrder)
	return newRoot
}

func rewrite(root *treeNode, b Basis, applicativeOrder bool) *treeNode {
	leftMostLeaf := getLeftMostLeaf(root)
	combinator, ok := findCombinator(leftMostLeaf.Leaf, b)
	numArgs := len(combinator.Arguments)
	numNodesToRoot := numNodesToRoot(leftMostLeaf, root)
	if ok && numArgs <= numNodesToRoot {
		// Construct new tree based off of combinator
		combinatorRoot := parse(combinator.Definition)

		// Find root of rewrite
		argumentNodes := getNRightSiblings(leftMostLeaf, numArgs)
		rewriteRoot := argumentNodes[len(argumentNodes)-1].Parent

		// Optionally do applicative first
		if applicativeOrder {
			for i := 0; i < len(argumentNodes); i++ {
				argumentNodes[i] = reduce(argumentNodes[i], b, applicativeOrder)
			}
		}

		// Apply the combinator
		combinatorRoot = apply(combinator, argumentNodes, combinatorRoot)

		// Swap parents
		if !rewriteRoot.IsRoot {
			combinatorRoot.IsRoot = false
			combinatorRoot.Parent = rewriteRoot.Parent
			if rewriteRoot != root {
				combinatorRoot.Parent.Left = combinatorRoot
			}
		}

		originalRoot := getNthParent(combinatorRoot, numNodesToRoot-numArgs)
		return rewrite(originalRoot, b, applicativeOrder)
	}
	return root
}

func apply(combinator Combinator, argumentNodes []*treeNode, combinatorRoot *treeNode) *treeNode {
	if combinatorRoot.IsLeaf {
		index := slices.Index(combinator.Arguments, combinatorRoot.Leaf)

		// For "improper" combinators. That is, combinators that are "defined" from other combinators
		if index == -1 {
			return combinatorRoot
		}

		arg := copy(argumentNodes[index])
		arg.IsRoot = combinatorRoot.IsRoot
		arg.Parent = combinatorRoot.Parent
		return arg
	}

	combinatorRoot.Left = apply(combinator, argumentNodes, combinatorRoot.Left)
	combinatorRoot.Right = apply(combinator, argumentNodes, combinatorRoot.Right)
	return combinatorRoot
}
