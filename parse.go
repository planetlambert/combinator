package combinator

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Parses a statement into a tree
func parse(statement string) *Tree {
	if len(statement) == 0 {
		return nil
	}

	// Returns a single-node tree if there is only one character
	if utf8.RuneCountInString(statement) == 1 {
		return &Tree{
			IsRoot: true,
			IsLeaf: true,
			Leaf:   statement,
		}
	}

	var root *Tree
	for _, subStatement := range getSubStatements(statement) {
		// Recursively parse each substatement
		subTree := parse(subStatement)
		subTree.IsRoot = false

		// Send the subtree down the left-most spine
		if root == nil {
			root = subTree
		} else {
			root = &Tree{
				Left:  root,
				Right: subTree,
			}
			root.Left.Parent = root
			root.Right.Parent = root
		}
	}

	// Return our tree, ensuring the root boolean is set
	root.IsRoot = true
	return root
}

// Returns a slice of sub statements. For `f(xy)z`, the
// sub statements are `f`, `xy`, and `z` are sub statements
func getSubStatements(statement string) []string {
	subStatements := []string{}

	var count int
	var tempStatement strings.Builder
	for _, ch := range statement {
		switch ch {
		case '(':
			if count > 0 {
				tempStatement.WriteRune(ch)
			}
			count++
		case ')':
			count--
			if count > 0 {
				tempStatement.WriteRune(ch)
			}
			if count == 0 {
				subStatements = append(subStatements, tempStatement.String())
				tempStatement.Reset()
			}
		default:
			if count > 0 {
				tempStatement.WriteRune(ch)
			} else {
				subStatements = append(subStatements, string(ch))
			}
		}
	}

	return subStatements
}

// Unparses a tree into a string statement
func unparse(root *Tree) string {
	var statement strings.Builder

	// Start with the left-most leaf
	current := getLeftMostLeaf(root)
	statement.WriteString(current.Leaf)

	for {
		// When we get to the root, we are done
		if current == root {
			break
		}

		// Move up the tree
		current = current.Parent

		// Recursively unparse the right child
		if current.Right.IsLeaf {
			statement.WriteString(unparse(current.Right))
		} else {
			statement.WriteString(fmt.Sprintf("(%s)", unparse(current.Right)))
		}
	}

	return statement.String()
}

// Checks if the statement is well-defined
func isWellDefined(statement string) error {
	// Empty statements don't make much sense
	if len(statement) == 0 {
		return errors.New("statement cannot be empty")
	}

	var count int
	var prevCharOpenParen bool
	for _, ch := range statement {
		switch ch {
		case '(':
			count++
			prevCharOpenParen = true
		case ')':
			// No `()`
			if prevCharOpenParen {
				return errors.New("parens cannot be empty ()")
			}
			// No `x)`
			if count == 0 {
				return errors.New("closed paren must have open paren")
			}
			count--
		default:
			prevCharOpenParen = false
		}
	}
	// Parentheses obviously must match
	if count != 0 {
		return errors.New("parens do not match")
	}
	return nil
}
