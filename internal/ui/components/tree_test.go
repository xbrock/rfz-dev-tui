package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

// TuiTreeItem tests

func TestTreeItem_Leaf(t *testing.T) {
	node := components.TuiTreeNode{Label: "main.go", Metadata: ""}
	output := components.TuiTreeItem(node, 0, false, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTreeItem_Expanded(t *testing.T) {
	node := components.TuiTreeNode{
		Label:    "src",
		Expanded: true,
		Children: []components.TuiTreeNode{{Label: "child"}},
	}
	output := components.TuiTreeItem(node, 0, false, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTreeItem_Collapsed(t *testing.T) {
	node := components.TuiTreeNode{
		Label:    "vendor",
		Expanded: false,
		Children: []components.TuiTreeNode{{Label: "child"}},
	}
	output := components.TuiTreeItem(node, 0, false, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTreeItem_Cursor_Focused(t *testing.T) {
	node := components.TuiTreeNode{Label: "build.go", Metadata: "modified"}
	output := components.TuiTreeItem(node, 1, true, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTreeItem_WithMetadata(t *testing.T) {
	node := components.TuiTreeNode{Label: "boss", Metadata: "clean"}
	output := components.TuiTreeItem(node, 0, false, false)
	golden.RequireEqual(t, []byte(output))
}

// TuiTree tests

func TestTree_Empty(t *testing.T) {
	output := components.TuiTree(nil, -1, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTree_FlatList(t *testing.T) {
	nodes := []components.TuiTreeNode{
		{Label: "README.md"},
		{Label: "go.mod"},
		{Label: "main.go"},
	}
	output := components.TuiTree(nodes, -1, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTree_NestedExpanded(t *testing.T) {
	nodes := []components.TuiTreeNode{
		{
			Label:    "project",
			Expanded: true,
			Children: []components.TuiTreeNode{
				{
					Label:    "src",
					Expanded: true,
					Children: []components.TuiTreeNode{
						{Label: "main.go"},
						{Label: "util.go"},
					},
				},
				{Label: "README.md"},
			},
		},
	}
	output := components.TuiTree(nodes, -1, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTree_MixedExpandCollapse(t *testing.T) {
	nodes := []components.TuiTreeNode{
		{
			Label:    "root",
			Expanded: true,
			Children: []components.TuiTreeNode{
				{
					Label:    "open-dir",
					Expanded: true,
					Children: []components.TuiTreeNode{
						{Label: "file-a.go"},
					},
				},
				{
					Label:    "closed-dir",
					Expanded: false,
					Children: []components.TuiTreeNode{
						{Label: "hidden.go"},
					},
				},
				{Label: "leaf.txt"},
			},
		},
	}
	output := components.TuiTree(nodes, -1, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTree_WithCursor(t *testing.T) {
	nodes := []components.TuiTreeNode{
		{
			Label:    "project",
			Expanded: true,
			Children: []components.TuiTreeNode{
				{Label: "main.go"},
				{Label: "util.go"},
			},
		},
	}
	// Cursor on "main.go" (index 1 in flattened list)
	output := components.TuiTree(nodes, 1, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTree_WithMetadata(t *testing.T) {
	nodes := []components.TuiTreeNode{
		{Label: "boss", Metadata: "clean"},
		{Label: "worker-1", Metadata: "building"},
		{Label: "worker-2", Metadata: "idle"},
	}
	output := components.TuiTree(nodes, -1, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTree_VisibleNodeCount(t *testing.T) {
	nodes := []components.TuiTreeNode{
		{
			Label:    "root",
			Expanded: true,
			Children: []components.TuiTreeNode{
				{Label: "child-1"},
				{
					Label:    "child-2",
					Expanded: false,
					Children: []components.TuiTreeNode{
						{Label: "hidden"},
					},
				},
			},
		},
	}
	count := components.VisibleNodeCount(nodes)
	if count != 3 {
		t.Errorf("expected 3 visible nodes, got %d", count)
	}
}
