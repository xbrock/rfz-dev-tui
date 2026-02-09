// Package components provides shared UI components and styles.
//
// This file contains all Lip Gloss style definitions extracted from design-system.md.
// CRITICAL: All visual styling MUST use these Lip Gloss styles.
// Never use manual string padding, ANSI codes, or custom border drawing.
package components

import "github.com/charmbracelet/lipgloss"

// ============================================================================
// COLOR TOKENS
// ============================================================================

// Base Colors - Background and Surface
var (
	// ColorBackground is the main background color (very dark blue-gray)
	ColorBackground = lipgloss.Color("#1e1e2e")

	// ColorCard is for card/panel backgrounds (slightly lighter)
	ColorCard = lipgloss.Color("#2a2a3e")

	// ColorSecondary is for secondary surfaces (muted gray)
	ColorSecondary = lipgloss.Color("#3a3a4e")

	// ColorBorder is the default border color (subtle gray)
	ColorBorder = lipgloss.Color("#4a4a5e")
)

// Accent Colors - Interactive States
var (
	// ColorCyan is the primary focus/interactive color
	ColorCyan = lipgloss.Color("#0891b2")

	// ColorNavActiveBg is the light blue background for active navigation items
	ColorNavActiveBg = lipgloss.Color("#164e63")

	// ColorGreen is for success states
	ColorGreen = lipgloss.Color("#22c55e")

	// ColorYellow is for warning states
	ColorYellow = lipgloss.Color("#eab308")

	// ColorDestructive is for error/danger states (red)
	ColorDestructive = lipgloss.Color("#ef4444")

	// ColorBrand is the DB Red brand accent
	ColorBrand = lipgloss.Color("#ec0016")
)

// Text Colors
var (
	// ColorTextPrimary is high contrast white for primary text
	ColorTextPrimary = lipgloss.Color("#f4f4f5")

	// ColorTextSecondary is muted gray for secondary text
	ColorTextSecondary = lipgloss.Color("#a1a1aa")

	// ColorTextMuted is very subtle gray for hints
	ColorTextMuted = lipgloss.Color("#71717a")

	// ColorTextDisabled is low contrast for disabled text
	ColorTextDisabled = lipgloss.Color("#52525b")
)

// Adaptive Colors (for light/dark terminal support)
var (
	ColorAdaptiveBackground = lipgloss.AdaptiveColor{
		Light: "#ffffff",
		Dark:  "#1e1e2e",
	}

	ColorAdaptiveForeground = lipgloss.AdaptiveColor{
		Light: "#1e1e2e",
		Dark:  "#f4f4f5",
	}
)

// ============================================================================
// UNICODE SYMBOLS
// ============================================================================

const (
	// Checkbox symbols (charm-style ballot box)
	SymbolCheckboxUnchecked = "☐" // U+2610 Ballot Box
	SymbolCheckboxChecked   = "☑" // U+2611 Ballot Box with Check

	// Radio symbols (charm-style circle)
	SymbolRadioUnselected = "◯" // U+25EF Large Circle
	SymbolRadioSelected   = "◉" // U+25C9 Fisheye (filled circle)

	// Circle select symbols (for component multi-select lists)
	SymbolCircleUnselected = "○" // U+25CB White Circle
	SymbolCircleSelected   = "◉" // U+25C9 Fisheye (filled circle)

	// List/navigation symbols
	SymbolListPointer = "›" // U+203A Single Right-Pointing Angle Quotation Mark

	// Tree expand/collapse symbols
	SymbolExpanded  = "▼" // U+25BC Black Down-Pointing Triangle
	SymbolCollapsed = "▶" // U+25B6 Black Right-Pointing Triangle
)

// ============================================================================
// SPACING CONSTANTS
// ============================================================================

const (
	SpaceNone = 0
	SpaceXS   = 1 // Tight: between related items
	SpaceSM   = 2 // Compact: list items
	SpaceMD   = 3 // Default: component padding
	SpaceLG   = 4 // Generous: section spacing
	SpaceXL   = 6 // Large: major sections
	Space2XL  = 8 // Extra large: page margins
)

// ============================================================================
// BORDER STYLES
// ============================================================================

var (
	BorderSingle  = lipgloss.NormalBorder()  // Single line: +--+
	BorderDouble  = lipgloss.DoubleBorder()  // Double line
	BorderRounded = lipgloss.RoundedBorder() // Rounded corners
	BorderHeavy   = lipgloss.ThickBorder()   // Heavy/thick
)

// ============================================================================
// TYPOGRAPHY STYLES
// ============================================================================

// Headings
var (
	// StyleH1 is for page titles (bold)
	StyleH1 = lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorTextPrimary)

	// StyleH2 is for section headers (bold, cyan accent)
	StyleH2 = lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorCyan)

	// StyleH3 is for subsection headers
	StyleH3 = lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorTextSecondary)
)

// Body Text
var (
	// StyleBody is for primary content text
	StyleBody = lipgloss.NewStyle().
		Foreground(ColorTextPrimary)

	// StyleBodySecondary is for less important text
	StyleBodySecondary = lipgloss.NewStyle().
		Foreground(ColorTextSecondary)

	// StyleBodyMuted is for hints and captions
	StyleBodyMuted = lipgloss.NewStyle().
		Foreground(ColorTextMuted).
		Italic(true)
)

// Special Text
var (
	// StyleCode is for code/monospace content
	StyleCode = lipgloss.NewStyle().
		Background(ColorSecondary).
		Padding(0, 1)

	// StyleKeyboard is for keyboard shortcuts
	StyleKeyboard = lipgloss.NewStyle().
		Foreground(ColorTextSecondary).
		Bold(true)
)

// ============================================================================
// BOX STYLES
// ============================================================================

var (
	// StyleBoxDefault is a single-bordered container
	StyleBoxDefault = lipgloss.NewStyle().
		Border(BorderSingle).
		BorderForeground(ColorBorder).
		Padding(1, 2)

	// StyleBoxDouble is a double-bordered container
	StyleBoxDouble = lipgloss.NewStyle().
		Border(BorderDouble).
		BorderForeground(ColorBorder).
		Padding(1, 2)

	// StyleBoxRounded is a rounded-bordered container
	StyleBoxRounded = lipgloss.NewStyle().
		Border(BorderRounded).
		BorderForeground(ColorBorder).
		Padding(1, 2)

	// StyleBoxFocused is a container with cyan focus highlight
	StyleBoxFocused = lipgloss.NewStyle().
		Border(BorderSingle).
		BorderForeground(ColorCyan).
		Padding(1, 2)
)

// ============================================================================
// NAVIGATION STYLES
// ============================================================================

var (
	// StyleNavItem is a normal navigation item
	StyleNavItem = lipgloss.NewStyle().
		Foreground(ColorTextPrimary).
		PaddingLeft(2)

	// StyleNavItemFocused is a focused navigation item (cursor on it)
	StyleNavItemFocused = lipgloss.NewStyle().
		Foreground(ColorCyan).
		Bold(true).
		PaddingLeft(0)

	// StyleNavItemActive is the currently active/selected navigation item
	StyleNavItemActive = lipgloss.NewStyle().
		Foreground(ColorTextPrimary).
		Background(ColorNavActiveBg).
		Bold(true).
		PaddingLeft(2).
		PaddingRight(2)
)

// ============================================================================
// TAB STYLES
// ============================================================================

var (
	// StyleTabNormal is for inactive, unfocused tabs
	StyleTabNormal = lipgloss.NewStyle().
		Foreground(ColorTextSecondary)

	// StyleTabActive is for the currently active/selected tab
	StyleTabActive = lipgloss.NewStyle().
		Foreground(ColorTextPrimary).
		Background(ColorSecondary).
		Bold(true).
		Padding(0, 1)

	// StyleTabFocused is for the keyboard-focused tab (distinct from active)
	StyleTabFocused = lipgloss.NewStyle().
		Foreground(ColorCyan).
		Bold(true).
		Underline(true)
)

// ============================================================================
// STATUS/BADGE STYLES
// ============================================================================

var (
	// StyleBadgeVersion is for version badges (like "v1.0.0")
	StyleBadgeVersion = lipgloss.NewStyle().
		Background(ColorBrand).
		Foreground(ColorTextPrimary).
		Bold(true).
		Padding(0, 1)

	// StyleBadgeInfo is for informational badges
	StyleBadgeInfo = lipgloss.NewStyle().
		Border(BorderSingle).
		BorderForeground(ColorBorder).
		Foreground(ColorTextSecondary).
		Padding(0, 1)
)

// ============================================================================
// BUTTON STYLES
// ============================================================================

var (
	// StyleButtonPrimary is a filled primary button
	StyleButtonPrimary = lipgloss.NewStyle().
		Background(ColorCyan).
		Foreground(ColorBackground).
		Bold(true).
		Padding(0, 2)

	// StyleButtonSecondary is an outlined secondary button
	StyleButtonSecondary = lipgloss.NewStyle().
		Border(BorderSingle).
		BorderForeground(ColorBorder).
		Foreground(ColorTextPrimary).
		Padding(0, 2)

	// StyleButtonDestructive is for dangerous actions
	StyleButtonDestructive = lipgloss.NewStyle().
		Background(ColorDestructive).
		Foreground(ColorTextPrimary).
		Bold(true).
		Padding(0, 2)
)

// ============================================================================
// INPUT STYLES
// ============================================================================

var (
	StyleInputNormal = lipgloss.NewStyle().
			Border(BorderSingle).
			BorderForeground(ColorBorder).
			Padding(0, 1)

	StyleInputFocused = lipgloss.NewStyle().
				Border(BorderSingle).
				BorderForeground(ColorCyan).
				Padding(0, 1)

	StyleInputError = lipgloss.NewStyle().
			Border(BorderSingle).
			BorderForeground(ColorDestructive).
			Padding(0, 1)
)

// ============================================================================
// LOG STYLES
// ============================================================================

var (
	StyleLogInfo = lipgloss.NewStyle().
			Foreground(ColorCyan)

	StyleLogSuccess = lipgloss.NewStyle().
			Foreground(ColorGreen)

	StyleLogWarning = lipgloss.NewStyle().
			Foreground(ColorYellow)

	StyleLogError = lipgloss.NewStyle().
			Foreground(ColorDestructive)

	StyleLogDebug = lipgloss.NewStyle().
			Foreground(ColorTextMuted)

	StyleLogTimestamp = lipgloss.NewStyle().
				Foreground(ColorTextMuted)

	StyleLogMessage = lipgloss.NewStyle().
			Foreground(ColorTextPrimary)
)

// ============================================================================
// HEADER STYLES
// ============================================================================

var (
	// StyleHeader is the main application header bar
	StyleHeader = lipgloss.NewStyle().
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(ColorBrand).
			Padding(0, 1).
			MarginBottom(1)

	// StyleHeaderTitle is for the header title (left side)
	StyleHeaderTitle = lipgloss.NewStyle().
				Foreground(ColorTextPrimary)

	// StyleHeaderSubtitle is for the header subtitle
	StyleHeaderSubtitle = lipgloss.NewStyle().
				Foreground(ColorCyan)
)

// ============================================================================
// FOOTER STYLES
// ============================================================================

var (
	// StyleFooter is the footer bar at bottom of screen
	StyleFooter = lipgloss.NewStyle().
			Background(ColorCard).
			Foreground(ColorTextSecondary).
			Padding(0, 1)
)

// FooterItem renders a key+label footer item with styled key and muted label.
func FooterItem(key string, label string) string {
	keyStyle := lipgloss.NewStyle().
		Foreground(ColorCyan).
		Bold(true)

	labelStyle := lipgloss.NewStyle().
		Foreground(ColorTextSecondary)

	return keyStyle.Render(key) + " " + labelStyle.Render(label)
}

// FooterItemActive renders an active/current footer badge as a colored pill.
func FooterItemActive(label string) string {
	return lipgloss.NewStyle().
		Background(ColorCyan).
		Foreground(ColorBackground).
		Bold(true).
		Padding(0, 1).
		Render(label)
}

// ============================================================================
// CONTENT STYLES
// ============================================================================

var (
	// StyleContent is for the main content area
	StyleContent = lipgloss.NewStyle().
			Border(BorderSingle).
			BorderForeground(ColorBorder).
			Padding(1, 2)

	// StyleASCIIArt is for the welcome ASCII art (DB Red)
	StyleASCIIArt = lipgloss.NewStyle().
			Foreground(ColorBrand).
			Align(lipgloss.Center)

	// StyleASCIIArtCyan is for the "CLI" part of ASCII art
	StyleASCIIArtCyan = lipgloss.NewStyle().
				Foreground(ColorCyan).
				Align(lipgloss.Center)

	// StyleTagline is for taglines/quotes
	StyleTagline = lipgloss.NewStyle().
			Foreground(ColorTextMuted).
			Italic(true).
			Align(lipgloss.Center).
			MarginTop(1)

	// StylePrompt is for the command prompt symbol
	StylePrompt = lipgloss.NewStyle().
			Foreground(ColorYellow).
			Bold(true)
)
