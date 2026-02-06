// Package demo provides demo screens for showcasing UI components.
//
// This file contains the LayoutGallery interactive demo for layout/navigation components.
package demo

import (
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// section identifies which interactive section currently has focus.
type section int

const (
	sectionNav section = iota
	sectionTree
	sectionTabs
	sectionTable
)

const sectionCount = 4

// LayoutGallery is a Bubble Tea model showcasing all layout/navigation components.
type LayoutGallery struct {
	width  int
	height int
	ready  bool

	// Active section for keyboard focus
	activeSection section

	// TuiNavigation state
	navCursor int
	navActive int

	// TuiTree state
	treeNodes  []components.TuiTreeNode
	treeCursor int

	// TuiTabs state
	tabs       []components.TuiTab
	activeTab  int
	focusedTab int

	// TuiTable state (bubbles/table is stateful)
	tableModel table.Model

	// TuiModal state
	modalOpen bool
	modalFocus int

	// StatusBar hints change based on active section
}

// NewLayoutGallery creates a new LayoutGallery model.
func NewLayoutGallery() LayoutGallery {
	treeNodes := []components.TuiTreeNode{
		{
			Label:    "src",
			Expanded: true,
			Children: []components.TuiTreeNode{
				{
					Label:    "internal",
					Expanded: true,
					Children: []components.TuiTreeNode{
						{Label: "config.go", Metadata: "2.1 KB"},
						{Label: "server.go", Metadata: "4.3 KB"},
					},
				},
				{
					Label:    "ui",
					Expanded: false,
					Children: []components.TuiTreeNode{
						{Label: "app.go", Metadata: "1.8 KB"},
						{Label: "styles.go", Metadata: "3.2 KB"},
					},
				},
			},
		},
		{Label: "go.mod", Metadata: "0.5 KB"},
		{Label: "main.go", Metadata: "0.3 KB"},
	}

	tabs := []components.TuiTab{
		{Label: "Overview", Active: true},
		{Label: "Build", Badge: 3},
		{Label: "Config"},
		{Label: "Logs", Badge: 12},
	}

	columns := []table.Column{
		{Title: "Component", Width: 20},
		{Title: "Type", Width: 15},
		{Title: "Status", Width: 12},
	}
	rows := []table.Row{
		{"TuiNavigation", "Layout", "Done"},
		{"TuiModal", "Overlay", "Done"},
		{"TuiKeyHints", "Display", "Done"},
		{"TuiTable", "Data", "Done"},
		{"TuiTree", "Hierarchy", "Done"},
		{"TuiTabs", "Navigation", "Done"},
		{"TuiStatusBar", "Layout", "Done"},
	}
	tableModel := components.NewTuiTable(components.TuiTableConfig{
		Columns: columns,
		Rows:    rows,
		Height:  7,
		Focused: false,
	})

	return LayoutGallery{
		navCursor:     0,
		navActive:     0,
		treeNodes:     treeNodes,
		treeCursor:    0,
		tabs:          tabs,
		activeTab:     0,
		focusedTab:    0,
		tableModel:    tableModel,
		activeSection: sectionNav,
	}
}

// Init implements tea.Model.
func (m LayoutGallery) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m LayoutGallery) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case tea.KeyMsg:
		// Modal takes priority when open
		if m.modalOpen {
			return m.updateModal(msg)
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "m":
			m.modalOpen = true
			m.modalFocus = 0
			return m, nil
		case "tab":
			m.activeSection = (m.activeSection + 1) % sectionCount
			m.updateTableFocus()
			return m, nil
		case "shift+tab":
			m.activeSection = (m.activeSection - 1 + sectionCount) % sectionCount
			m.updateTableFocus()
			return m, nil
		}

		// Dispatch to active section
		switch m.activeSection {
		case sectionNav:
			return m.updateNav(msg)
		case sectionTree:
			return m.updateTree(msg)
		case sectionTabs:
			return m.updateTabs(msg)
		case sectionTable:
			return m.updateTable(msg)
		}
	}

	return m, nil
}

func (m LayoutGallery) updateModal(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.modalOpen = false
	case "tab":
		m.modalFocus = (m.modalFocus + 1) % 2
	case "shift+tab":
		m.modalFocus = (m.modalFocus - 1 + 2) % 2
	case "enter":
		m.modalOpen = false
	case "q", "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func (m LayoutGallery) updateNav(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	navItems := m.navItems()
	switch msg.String() {
	case "j", "down":
		if m.navCursor < len(navItems)-1 {
			m.navCursor++
		}
	case "k", "up":
		if m.navCursor > 0 {
			m.navCursor--
		}
	case "enter":
		m.navActive = m.navCursor
	}
	return m, nil
}

func (m LayoutGallery) updateTree(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	visibleCount := components.VisibleNodeCount(m.treeNodes)
	switch msg.String() {
	case "j", "down":
		if m.treeCursor < visibleCount-1 {
			m.treeCursor++
		}
	case "k", "up":
		if m.treeCursor > 0 {
			m.treeCursor--
		}
	case "enter":
		m.toggleTreeNode()
	}
	return m, nil
}

func (m *LayoutGallery) toggleTreeNode() {
	idx := 0
	toggleAt(m.treeNodes, m.treeCursor, &idx)
}

func toggleAt(nodes []components.TuiTreeNode, target int, idx *int) bool {
	for i := range nodes {
		if *idx == target {
			if len(nodes[i].Children) > 0 {
				nodes[i].Expanded = !nodes[i].Expanded
			}
			return true
		}
		*idx++
		if len(nodes[i].Children) > 0 && nodes[i].Expanded {
			if toggleAt(nodes[i].Children, target, idx) {
				return true
			}
		}
	}
	return false
}

func (m LayoutGallery) updateTabs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "l", "right":
		if m.focusedTab < len(m.tabs)-1 {
			m.focusedTab++
		}
	case "h", "left":
		if m.focusedTab > 0 {
			m.focusedTab--
		}
	case "enter":
		// Set active tab
		for i := range m.tabs {
			m.tabs[i].Active = false
		}
		m.activeTab = m.focusedTab
		m.tabs[m.activeTab].Active = true
	case "1", "2", "3", "4":
		idx := int(msg.String()[0] - '1')
		if idx < len(m.tabs) {
			for i := range m.tabs {
				m.tabs[i].Active = false
			}
			m.activeTab = idx
			m.focusedTab = idx
			m.tabs[idx].Active = true
		}
	}
	return m, nil
}

func (m LayoutGallery) updateTable(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.tableModel, cmd = m.tableModel.Update(msg)
	return m, cmd
}

func (m *LayoutGallery) updateTableFocus() {
	if m.activeSection == sectionTable {
		m.tableModel.Focus()
	} else {
		m.tableModel.Blur()
	}
}

// View implements tea.Model.
func (m LayoutGallery) View() string {
	if !m.ready {
		return "Initializing..."
	}

	// Render modal overlay when open
	if m.modalOpen {
		return m.renderModal()
	}

	var rows []string

	// Header
	rows = append(rows, m.renderHeader())

	// Main content: Navigation (left) | Showcase (right)
	rows = append(rows, m.renderMainContent())

	// Status bar at bottom
	rows = append(rows, m.renderStatusBar())

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m LayoutGallery) renderHeader() string {
	title := components.StyleH1.Render("Layout Component Gallery")
	subtitle := lipgloss.NewStyle().
		Foreground(components.ColorTextMuted).
		Render("  Interactive Demo — Tab to switch sections, m for modal")

	return title + subtitle + "\n" +
		components.TuiDivider(components.DividerSingle, m.width) + "\n"
}

func (m LayoutGallery) renderMainContent() string {
	navWidth := 30
	rightWidth := m.width - navWidth - 3 // 3 for separator spacing
	if rightWidth < 20 {
		rightWidth = 20
	}

	// Left: Navigation
	navContent := m.renderNavSection(navWidth)

	// Separator
	sep := lipgloss.NewStyle().
		Foreground(components.ColorBorder).
		Render(" │ ")

	// Right: Component showcase sections
	rightContent := m.renderShowcaseSections(rightWidth)

	return lipgloss.JoinHorizontal(lipgloss.Top, navContent, sep, rightContent)
}

func (m LayoutGallery) renderNavSection(width int) string {
	var sb strings.Builder

	sectionLabel := m.sectionLabel(sectionNav)
	sb.WriteString(sectionLabel)
	sb.WriteString("\n\n")

	navItems := m.navItems()
	focused := m.activeSection == sectionNav
	nav := components.TuiNavigation(navItems, m.navCursor, m.navActive, focused, "Screens", "", width)
	sb.WriteString(nav)

	// Pad to fill vertical space
	content := sb.String()
	contentHeight := lipgloss.Height(content)
	targetHeight := m.height - 5 // account for header + statusbar
	if targetHeight > contentHeight {
		content += strings.Repeat("\n", targetHeight-contentHeight)
	}

	return lipgloss.NewStyle().Width(width).Render(content)
}

func (m LayoutGallery) renderShowcaseSections(width int) string {
	var sections []string

	// Tabs section
	sections = append(sections, m.renderTabsSection(width))

	// Tree section
	sections = append(sections, m.renderTreeSection(width))

	// Table section
	sections = append(sections, m.renderTableSection(width))

	// KeyHints standalone demo
	sections = append(sections, m.renderKeyHintsSection(width))

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m LayoutGallery) renderTabsSection(width int) string {
	var sb strings.Builder

	sectionLabel := m.sectionLabel(sectionTabs)
	sb.WriteString(sectionLabel)
	sb.WriteString("\n")

	focusIdx := -1
	if m.activeSection == sectionTabs {
		focusIdx = m.focusedTab
	}
	sb.WriteString(components.TuiTabs(m.tabs, focusIdx, width))
	sb.WriteString("\n\n")

	return sb.String()
}

func (m LayoutGallery) renderTreeSection(width int) string {
	var sb strings.Builder

	sectionLabel := m.sectionLabel(sectionTree)
	sb.WriteString(sectionLabel)
	sb.WriteString("\n")

	focused := m.activeSection == sectionTree
	cursorIdx := -1
	if focused {
		cursorIdx = m.treeCursor
	}
	tree := components.TuiTree(m.treeNodes, cursorIdx, focused)
	sb.WriteString(tree)
	sb.WriteString("\n\n")

	return sb.String()
}

func (m LayoutGallery) renderTableSection(width int) string {
	var sb strings.Builder

	sectionLabel := m.sectionLabel(sectionTable)
	sb.WriteString(sectionLabel)
	sb.WriteString("\n")

	sb.WriteString(m.tableModel.View())
	sb.WriteString("\n\n")

	return sb.String()
}

func (m LayoutGallery) renderKeyHintsSection(width int) string {
	var sb strings.Builder

	label := components.StyleH3.Render("TuiKeyHints")
	sb.WriteString(label)
	sb.WriteString("\n")

	hints := []components.KeyHint{
		{Key: "j/k", Label: "Navigate"},
		{Key: "Enter", Label: "Select"},
		{Key: "Tab", Label: "Section"},
		{Key: "m", Label: "Modal"},
		{Key: "q", Label: "Quit"},
	}
	sb.WriteString(components.TuiKeyHints(hints, width))
	sb.WriteString("\n")

	return sb.String()
}

func (m LayoutGallery) renderModal() string {
	config := components.TuiModalConfig{
		Title:   "Demo Modal",
		Content: "This is a demo modal dialog.\nPress Esc to close or Enter to confirm.\nUse Tab to switch between buttons.",
		Buttons: []components.TuiModalButton{
			{Label: "Confirm", Variant: components.ButtonPrimary, Shortcut: "Enter"},
			{Label: "Cancel", Variant: components.ButtonSecondary, Shortcut: "Esc"},
		},
		Width:        50,
		FocusedIndex: m.modalFocus,
	}
	return components.TuiModal(config, m.width, m.height)
}

func (m LayoutGallery) renderStatusBar() string {
	hints := []components.KeyHint{
		{Key: "Tab", Label: "Section"},
		{Key: "m", Label: "Modal"},
		{Key: "q", Label: "Quit"},
	}

	sectionNames := [sectionCount]string{"Navigation", "Tree", "Tabs", "Table"}

	return components.TuiStatusBar(components.TuiStatusBarConfig{
		Status:      "Layout Demo",
		StatusColor: components.ColorGreen,
		Info:        sectionNames[m.activeSection],
		Hints:       hints,
		Width:       m.width,
	})
}

func (m LayoutGallery) navItems() []components.TuiNavItem {
	return []components.TuiNavItem{
		{Label: "Welcome", Number: 1, Shortcut: "w"},
		{Label: "Build", Number: 2, Shortcut: "b"},
		{Label: "Config", Number: 3, Shortcut: "c"},
		{Label: "Logs", Number: 4, Shortcut: "l"},
		{Label: "Discover", Number: 5, Shortcut: "d"},
	}
}

func (m LayoutGallery) sectionLabel(s section) string {
	names := [sectionCount]string{"TuiNavigation", "TuiTree", "TuiTabs", "TuiTable"}
	style := components.StyleH3
	if m.activeSection == s {
		style = components.StyleH2
	}
	label := style.Render(names[s])

	if m.activeSection == s {
		indicator := lipgloss.NewStyle().Foreground(components.ColorCyan).Render(" ●")
		return label + indicator
	}
	return label
}
