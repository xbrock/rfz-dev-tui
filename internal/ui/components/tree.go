// Package components provides shared UI components and styles.
//
// This file contains the TuiTree hierarchical view component.
package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// maxTreeDepth is the maximum nesting depth before showing "..." indicator.
const maxTreeDepth = 15

// treeIndentWidth is the number of spaces per indentation level.
const treeIndentWidth = 2

// TuiTreeNode represents a single node in a hierarchical tree.
type TuiTreeNode struct {
	Label    string         // Display text for the node
	Metadata string         // Optional metadata shown right of label (e.g., status)
	Children []TuiTreeNode  // Child nodes (empty = leaf node)
	Expanded bool           // Whether children are visible (ignored for leaf nodes)
}

// flatNode is a pre-computed visible node used for rendering.
type flatNode struct {
	label    string
	metadata string
	depth    int
	isLeaf   bool
	expanded bool
	truncated bool // depth exceeded maxTreeDepth
}

// flattenTree converts a tree into a flat list of visible nodes.
func flattenTree(nodes []TuiTreeNode, depth int) []flatNode {
	var result []flatNode
	for _, node := range nodes {
		isLeaf := len(node.Children) == 0

		if depth >= maxTreeDepth {
			result = append(result, flatNode{
				label:     node.Label,
				metadata:  node.Metadata,
				depth:     depth,
				isLeaf:    isLeaf,
				expanded:  node.Expanded,
				truncated: true,
			})
			continue
		}

		result = append(result, flatNode{
			label:    node.Label,
			metadata: node.Metadata,
			depth:    depth,
			isLeaf:   isLeaf,
			expanded: node.Expanded,
		})

		if !isLeaf && node.Expanded {
			result = append(result, flattenTree(node.Children, depth+1)...)
		}
	}
	return result
}

// TuiTreeItem renders a single tree node line.
// cursor: whether the cursor is on this node
// focused: whether the tree has focus
func TuiTreeItem(node TuiTreeNode, depth int, cursor bool, focused bool) string {
	isLeaf := len(node.Children) == 0
	return renderFlatNode(flatNode{
		label:    node.Label,
		metadata: node.Metadata,
		depth:    depth,
		isLeaf:   isLeaf,
		expanded: node.Expanded,
	}, cursor, focused)
}

// renderFlatNode renders a single flat node line.
func renderFlatNode(fn flatNode, cursor bool, focused bool) string {
	var parts []string

	// Cursor indicator
	if cursor {
		cursorStyle := lipgloss.NewStyle().Foreground(ColorCyan).Bold(true)
		parts = append(parts, cursorStyle.Render(SymbolListPointer))
	} else {
		parts = append(parts, " ")
	}

	// Indentation
	indent := strings.Repeat(" ", fn.depth*treeIndentWidth)
	parts = append(parts, indent)

	// Expand/collapse icon or leaf spacing
	if fn.truncated {
		iconStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)
		parts = append(parts, iconStyle.Render("..."))
	} else if fn.isLeaf {
		// Leaf nodes get spacing to align with parent labels
		parts = append(parts, "  ")
	} else if fn.expanded {
		iconStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary)
		parts = append(parts, iconStyle.Render(SymbolExpanded))
	} else {
		iconStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary)
		parts = append(parts, iconStyle.Render(SymbolCollapsed))
	}

	// Space between icon and label
	parts = append(parts, " ")

	// Label
	var labelStyle lipgloss.Style
	if cursor && focused {
		labelStyle = lipgloss.NewStyle().Foreground(ColorCyan).Bold(true)
	} else {
		labelStyle = lipgloss.NewStyle().Foreground(ColorTextPrimary)
	}
	parts = append(parts, labelStyle.Render(fn.label))

	// Metadata
	if fn.metadata != "" {
		metaStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)
		parts = append(parts, " "+metaStyle.Render(fn.metadata))
	}

	return strings.Join(parts, "")
}

// TuiTree renders a complete hierarchical tree view.
// nodes: root-level tree nodes
// cursorIndex: index in the visible (flattened) node list (-1 for no cursor)
// focused: whether the tree has focus
func TuiTree(nodes []TuiTreeNode, cursorIndex int, focused bool) string {
	if len(nodes) == 0 {
		emptyStyle := lipgloss.NewStyle().Foreground(ColorTextMuted).Italic(true)
		return emptyStyle.Render("No items")
	}

	visible := flattenTree(nodes, 0)

	var lines []string
	for i, fn := range visible {
		isCursor := i == cursorIndex
		lines = append(lines, renderFlatNode(fn, isCursor, focused))
	}

	return strings.Join(lines, "\n")
}

// VisibleNodeCount returns the number of visible nodes in the tree
// (useful for cursor bounds checking).
func VisibleNodeCount(nodes []TuiTreeNode) int {
	return len(flattenTree(nodes, 0))
}
