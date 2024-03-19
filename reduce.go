package combinator

import (
	"slices"
)

// Reduces the tree `tree` using basis `b`
func reduce(root *Tree, b Basis, applicativeOrder bool) *Tree {
	if root.IsLeaf {
		return root
	}

	// Our algorithm is outer-first (we attempt to rewrite before recursing
	// into the left and right child)
	newTree := rewrite(root, b, applicativeOrder)
	if newTree.IsLeaf {
		return newTree
	}

	// Out algorithm is also left-first (we recurse the left subtree first)
	newTree.Left = reduce(newTree.Left, b, applicativeOrder)
	newTree.Right = reduce(newTree.Right, b, applicativeOrder)
	return newTree
}

func rewrite(root *Tree, b Basis, applicativeOrder bool) *Tree {
	// Rewrite attempts use the left-most leaf of this subtree
	leftMostLeaf := getLeftMostLeaf(root)

	// Retrieve the left-most leaf's combinator if applicable
	combinator, ok := findCombinator(leftMostLeaf.Leaf, b)

	// How many variables does the combinator's definition require before rewriting makes sense?
	numArgs := len(combinator.Arguments)

	// The amount of nodes between the left-most leaf and the root of our subtree
	numNodesToRoot := numNodesToRoot(leftMostLeaf, root)

	// If we found a combinator, and have enough arguments, attempt a rewrite
	if ok && numArgs <= numNodesToRoot {
		// Construct new tree based off of combinator
		combinatorRoot := parse(combinator.Definition)

		// Get our list of arguments (all subtrees themselves)
		argumentNodes := getNRightSiblings(leftMostLeaf, numArgs)

		// Store a reference to the root of the subtree we are rewriting
		rewriteRoot := argumentNodes[len(argumentNodes)-1].Parent

		// Applicative order is when we reduce our arguments before applying our combinator
		if applicativeOrder {
			for i := 0; i < len(argumentNodes); i++ {
				argumentNodes[i] = reduce(argumentNodes[i], b, applicativeOrder)
			}
		}

		// Apply the combinator
		combinatorRoot = apply(combinator, argumentNodes, combinatorRoot)

		// Hook our post-rewrite subtree into the rest of the tree
		if !rewriteRoot.IsRoot {
			combinatorRoot.IsRoot = false
			combinatorRoot.Parent = rewriteRoot.Parent

			// Only set the parent's child if it's not our original root
			if rewriteRoot != root {
				combinatorRoot.Parent.Left = combinatorRoot
			}
		}

		// Find the original root
		originalRoot := getNthParent(combinatorRoot, numNodesToRoot-numArgs)

		// Recursively rewrite
		return rewrite(originalRoot, b, applicativeOrder)
	}
	return root
}

// Recursively pplies the arguments in `argumentNodes` to the Combinator.
func apply(combinator Combinator, argumentNodes []*Tree, combinatorRoot *Tree) *Tree {
	if combinatorRoot.IsLeaf {
		// Find the argument the combinator definition is referring to
		index := slices.Index(combinator.Arguments, combinatorRoot.Leaf)

		// For "improper" combinators. That is, combinators that are "defined" from other combinators,
		// just return whatever was in the definition (likely another combinator)
		if index == -1 {
			return combinatorRoot
		}

		// Get a fresh copy of the argument to apply
		arg := copy(argumentNodes[index])

		// Set any root/parent info and return
		arg.IsRoot = combinatorRoot.IsRoot
		arg.Parent = combinatorRoot.Parent
		return arg
	}

	// Recurse down the combinator definition
	combinatorRoot.Left = apply(combinator, argumentNodes, combinatorRoot.Left)
	combinatorRoot.Right = apply(combinator, argumentNodes, combinatorRoot.Right)
	return combinatorRoot
}
