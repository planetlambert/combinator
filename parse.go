package combinator

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func parse(statement string) *TreeNode {
	if utf8.RuneCountInString(statement) == 1 {
		return &TreeNode{
			IsRoot: true,
			IsLeaf: true,
			Leaf:   statement,
		}
	}

	var root *TreeNode
	for _, subStatement := range getSubStatements(statement) {
		subTree := parse(subStatement)
		subTree.IsRoot = false

		if root == nil {
			root = subTree
		} else {
			root = &TreeNode{
				Left:  root,
				Right: subTree,
			}
			root.Left.Parent = root
			root.Right.Parent = root
		}
	}
	root.IsRoot = true
	return root
}

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

func unparse(root *TreeNode) string {
	var statement strings.Builder

	current := getLeftMostLeaf(root)
	statement.WriteString(current.Leaf)

	for {
		if current == root {
			break
		}
		current = current.Parent
		if current.Right.IsLeaf {
			statement.WriteString(unparse(current.Right))
		} else {
			statement.WriteString(fmt.Sprintf("(%s)", unparse(current.Right)))
		}
	}

	return statement.String()
}

func isWellDefined(statement string) error {
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
			if prevCharOpenParen {
				return errors.New("parens cannot be empty ()")
			}
			if count == 0 {
				return errors.New("closed paren must have open paren")
			}
			count--
		default:
			prevCharOpenParen = false
		}
	}
	if count != 0 {
		return errors.New("parens do not match")
	}
	return nil
}
