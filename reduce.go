package combinator

import (
	"slices"
)

func reduce(root *TreeNode, b Basis, applicativeOrder bool) *TreeNode {
	leftMostLeaf := getLeftMostLeaf(root)

	combinator, ok := findCombinator(leftMostLeaf.Leaf, b)
	numArgs := len(combinator.Arguments)
	if ok && numArgs <= numNodesToRoot(leftMostLeaf, root) {
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

		// Rewrite the tree
		combinatorRoot = rewrite(combinator, argumentNodes, combinatorRoot)

		// Swap parents
		if !rewriteRoot.IsRoot {
			combinatorRoot.IsRoot = false
			combinatorRoot.Parent = rewriteRoot.Parent
			combinatorRoot.Parent.Left = combinatorRoot
		}

		return reduce(getRoot(combinatorRoot), b, applicativeOrder)
	} else {
		return root
	}
}

func rewrite(combinator Combinator, argumentNodes []*TreeNode, combinatorRoot *TreeNode) *TreeNode {
	if combinatorRoot.IsLeaf {
		arg := copy(argumentNodes[slices.Index(combinator.Arguments, combinatorRoot.Leaf)])
		arg.IsRoot = combinatorRoot.IsRoot
		arg.Parent = combinatorRoot.Parent
		return arg
	}

	combinatorRoot.Left = rewrite(combinator, argumentNodes, combinatorRoot.Left)
	combinatorRoot.Right = rewrite(combinator, argumentNodes, combinatorRoot.Right)
	return combinatorRoot
}
