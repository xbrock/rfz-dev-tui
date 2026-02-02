# Design System: RFZ-CLI

> Last Updated: 2026-02-02
> Version: 1.0.0
> Source: Prototype Screenshots (references/prototype-screenshots/)

## Overview

Design system extracted from RFZ-CLI prototype screenshots for Go/Bubble Tea TUI implementation. This document defines the visual language, component patterns, and design tokens using Lip Gloss styles for consistent terminal UI implementation.

**Target Framework:** Go with charmbracelet/bubbletea + charmbracelet/lipgloss

---

## Color Tokens (Lip Gloss)

### Go Package Setup

```go
package styles

import "github.com/charmbracelet/lipgloss"

// ============================================================================
// COLOR TOKENS
// ============================================================================

// Base Colors - Background and Surface
var (
    // Background - very dark blue-gray (near black)
    ColorBackground = lipgloss.Color("#1e1e2e")

    // Card - slightly lighter than background
    ColorCard = lipgloss.Color("#2a2a3e")

    // Secondary - muted gray for secondary surfaces
    ColorSecondary = lipgloss.Color("#3a3a4e")

    // Border - subtle gray for borders
    ColorBorder = lipgloss.Color("#4a4a5e")
)

// Accent Colors - Interactive States
var (
    // TUI Cyan - Primary focus/interactive color
    ColorCyan = lipgloss.Color("#0891b2")

    // TUI Green - Success states
    ColorGreen = lipgloss.Color("#22c55e")

    // TUI Yellow - Warning states
    ColorYellow = lipgloss.Color("#eab308")

    // Destructive - Error/danger states (red)
    ColorDestructive = lipgloss.Color("#ef4444")

    // Brand - DB Red (Deutsche Bahn brand color)
    ColorBrand = lipgloss.Color("#ec0016")
)

// Text Colors
var (
    // Primary text - high contrast white
    ColorTextPrimary = lipgloss.Color("#f4f4f5")

    // Secondary text - muted gray
    ColorTextSecondary = lipgloss.Color("#a1a1aa")

    // Muted text - very subtle gray
    ColorTextMuted = lipgloss.Color("#71717a")

    // Disabled text - low contrast
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

    ColorAdaptiveCyan = lipgloss.AdaptiveColor{
        Light: "#0891b2",
        Dark:  "#0891b2",
    }
)
```

### Color Palette Reference

| Token | Hex | Lip Gloss | Usage |
|-------|-----|-----------|-------|
| background | #1e1e2e | `lipgloss.Color("#1e1e2e")` | Main background |
| card | #2a2a3e | `lipgloss.Color("#2a2a3e")` | Card/panel backgrounds |
| secondary | #3a3a4e | `lipgloss.Color("#3a3a4e")` | Secondary surfaces |
| border | #4a4a5e | `lipgloss.Color("#4a4a5e")` | Default borders |
| tui-cyan | #0891b2 | `lipgloss.Color("#0891b2")` | Focus, links, interactive |
| tui-green | #22c55e | `lipgloss.Color("#22c55e")` | Success, ready states |
| tui-yellow | #eab308 | `lipgloss.Color("#eab308")` | Warnings, caution |
| destructive | #ef4444 | `lipgloss.Color("#ef4444")` | Errors, destructive actions |
| brand | #ec0016 | `lipgloss.Color("#ec0016")` | DB Red brand accent |

---

## Typography

### Font Family

**Terminal Font:** JetBrains Mono (or system monospace)

In TUI applications, typography is determined by the user's terminal font. The prototype uses JetBrains Mono as the recommended font for best appearance.

### Text Styles (Lip Gloss)

```go
// ============================================================================
// TYPOGRAPHY STYLES
// ============================================================================

// Headings
var (
    // H1 - Page title (bold)
    StyleH1 = lipgloss.NewStyle().
        Bold(true).
        Foreground(ColorTextPrimary)

    // H2 - Section header (bold, cyan accent)
    StyleH2 = lipgloss.NewStyle().
        Bold(true).
        Foreground(ColorCyan)

    // H3 - Subsection header
    StyleH3 = lipgloss.NewStyle().
        Bold(true).
        Foreground(ColorTextSecondary)
)

// Body Text
var (
    // Body - Primary content text
    StyleBody = lipgloss.NewStyle().
        Foreground(ColorTextPrimary)

    // Body Secondary - Less important text
    StyleBodySecondary = lipgloss.NewStyle().
        Foreground(ColorTextSecondary)

    // Body Muted - Hints, captions
    StyleBodyMuted = lipgloss.NewStyle().
        Foreground(ColorTextMuted).
        Italic(true)
)

// Special Text
var (
    // Code/Monospace - Already monospace in terminal
    StyleCode = lipgloss.NewStyle().
        Background(ColorSecondary).
        Padding(0, 1)

    // Link-like text
    StyleLink = lipgloss.NewStyle().
        Foreground(ColorCyan).
        Underline(true)

    // Keyboard shortcut
    StyleKeyboard = lipgloss.NewStyle().
        Foreground(ColorTextSecondary).
        Bold(true)
)
```

### Text Hierarchy

| Element | Style | Example Usage |
|---------|-------|---------------|
| Page Title | Bold, Primary | "RFZ-CLI v1.0.0" |
| Section Title | Bold, Cyan | "Navigation", "Color Tokens" |
| Subsection | Bold, Secondary | "TuiBox", "TuiNavItem" |
| Body | Regular, Primary | Main content |
| Muted | Italic, Muted | Descriptions, hints |
| Keyboard | Bold, Secondary | "Enter Select" |

---

## Spacing System

### TUI Spacing Scale

```go
// ============================================================================
// SPACING CONSTANTS
// ============================================================================

const (
    SpaceNone   = 0
    SpaceXS     = 1  // Tight: between related items
    SpaceSM     = 2  // Compact: list items
    SpaceMD     = 3  // Default: component padding
    SpaceLG     = 4  // Generous: section spacing
    SpaceXL     = 6  // Large: major sections
    Space2XL    = 8  // Extra large: page margins
)
```

### Lip Gloss Padding/Margin Usage

```go
// Padding examples
var StyleWithPaddingSM = lipgloss.NewStyle().Padding(0, SpaceSM)     // Horizontal only
var StyleWithPaddingMD = lipgloss.NewStyle().Padding(SpaceSM, SpaceMD) // Vert, Horiz
var StyleWithPaddingAll = lipgloss.NewStyle().Padding(SpaceMD)        // All sides

// Margin examples
var StyleWithMarginTop = lipgloss.NewStyle().MarginTop(SpaceSM)
var StyleWithMarginBottom = lipgloss.NewStyle().MarginBottom(SpaceMD)
var StyleWithMarginLeft = lipgloss.NewStyle().MarginLeft(SpaceLG)
```

### Spacing Guidelines

| Context | Padding | Margin |
|---------|---------|--------|
| Box content | 1-2 chars | - |
| List items | 0, 1 char horizontal | 0 |
| Sections | 1-2 lines vertical | 1 line between |
| Navigation panel | 1 char all sides | - |
| Help footer | 0, 1 char | - |

---

## Border Styles

### Lip Gloss Border Types

```go
// ============================================================================
// BORDER STYLES
// ============================================================================

// Border type references
var (
    BorderSingle  = lipgloss.NormalBorder()   // Single line: ┌─┐
    BorderDouble  = lipgloss.DoubleBorder()   // Double line: ╔═╗
    BorderRounded = lipgloss.RoundedBorder()  // Rounded: ╭─╮
    BorderHeavy   = lipgloss.ThickBorder()    // Heavy/thick: ┏━┓
)

// Box styles with borders
var (
    // Default box - single border
    StyleBoxDefault = lipgloss.NewStyle().
        Border(BorderSingle).
        BorderForeground(ColorBorder).
        Padding(1, 2)

    // Double border box
    StyleBoxDouble = lipgloss.NewStyle().
        Border(BorderDouble).
        BorderForeground(ColorBorder).
        Padding(1, 2)

    // Heavy border box
    StyleBoxHeavy = lipgloss.NewStyle().
        Border(BorderHeavy).
        BorderForeground(ColorBorder).
        Padding(1, 2)

    // Focused box - cyan border highlight
    StyleBoxFocused = lipgloss.NewStyle().
        Border(BorderSingle).
        BorderForeground(ColorCyan).
        Padding(1, 2)

    // Rounded box (for softer appearance)
    StyleBoxRounded = lipgloss.NewStyle().
        Border(BorderRounded).
        BorderForeground(ColorBorder).
        Padding(1, 2)
)
```

### Border Usage Guide

| Border Type | Lip Gloss | Use Case |
|-------------|-----------|----------|
| Single | `lipgloss.NormalBorder()` | Default containers, panels |
| Double | `lipgloss.DoubleBorder()` | Emphasized sections, dialogs |
| Rounded | `lipgloss.RoundedBorder()` | Softer UI, cards |
| Heavy | `lipgloss.ThickBorder()` | Important highlights |

### Focused State

```go
// Focus border adds cyan highlight
func FocusBorder(style lipgloss.Style) lipgloss.Style {
    return style.BorderForeground(ColorCyan)
}

// Example usage
var BoxNormal = StyleBoxDefault
var BoxFocused = FocusBorder(StyleBoxDefault)
```

---

## Layout Components

### TuiBox - Container Component

```go
// ============================================================================
// LAYOUT COMPONENTS
// ============================================================================

// TuiBox creates a bordered container
// Supports: single, double, rounded, heavy borders
type TuiBoxStyle struct {
    style lipgloss.Style
}

func NewTuiBox(borderType lipgloss.Border) TuiBoxStyle {
    return TuiBoxStyle{
        style: lipgloss.NewStyle().
            Border(borderType).
            BorderForeground(ColorBorder).
            Padding(1, 2),
    }
}

func (b TuiBoxStyle) Focused() TuiBoxStyle {
    return TuiBoxStyle{
        style: b.style.BorderForeground(ColorCyan),
    }
}

func (b TuiBoxStyle) Render(content string) string {
    return b.style.Render(content)
}

// Pre-built box styles
var (
    BoxSingle  = NewTuiBox(BorderSingle)
    BoxDouble  = NewTuiBox(BorderDouble)
    BoxRounded = NewTuiBox(BorderRounded)
    BoxHeavy   = NewTuiBox(BorderHeavy)
)
```

### TuiDivider - Separator Component

```go
// TuiDivider creates horizontal separator lines
func DividerSingle(width int) string {
    return lipgloss.NewStyle().
        Foreground(ColorBorder).
        Render(strings.Repeat("─", width))
}

func DividerDouble(width int) string {
    return lipgloss.NewStyle().
        Foreground(ColorBorder).
        Render(strings.Repeat("═", width))
}
```

---

## Navigation Components

### TuiNavItem - Menu Item

```go
// ============================================================================
// NAVIGATION COMPONENTS
// ============================================================================

// Navigation item states
var (
    // Normal navigation item
    StyleNavItem = lipgloss.NewStyle().
        Foreground(ColorTextPrimary).
        PaddingLeft(2)

    // Focused navigation item (cursor on it)
    StyleNavItemFocused = lipgloss.NewStyle().
        Foreground(ColorCyan).
        Bold(true).
        PaddingLeft(0)

    // Active/selected navigation item (current view)
    StyleNavItemActive = lipgloss.NewStyle().
        Foreground(ColorTextPrimary).
        Background(ColorSecondary).
        Bold(true).
        PaddingLeft(2).
        PaddingRight(2)
)

// NavItem renders a navigation menu item
// Format: "> 1. Build Components  1" (focused) or "  2. View Logs  2" (normal)
func NavItem(index int, label string, shortcut string, focused bool, active bool) string {
    prefix := "  "
    if focused {
        prefix = "> "
    }

    text := fmt.Sprintf("%s%d. %s", prefix, index, label)

    var style lipgloss.Style
    switch {
    case focused:
        style = StyleNavItemFocused
    case active:
        style = StyleNavItemActive
    default:
        style = StyleNavItem
    }

    // Add shortcut on right side
    rendered := style.Render(text)
    shortcutRendered := lipgloss.NewStyle().
        Foreground(ColorTextMuted).
        Render(shortcut)

    return lipgloss.JoinHorizontal(lipgloss.Top, rendered, "  ", shortcutRendered)
}
```

### Navigation Panel

```go
// NavigationPanel - complete sidebar navigation with help footer
var StyleNavPanel = lipgloss.NewStyle().
    Border(BorderSingle).
    BorderForeground(ColorBorder).
    Width(30).
    Padding(1)

// Navigation panel header
var StyleNavPanelHeader = lipgloss.NewStyle().
    Foreground(ColorTextSecondary).
    MarginBottom(1)

// Help footer within navigation
var StyleNavHelp = lipgloss.NewStyle().
    Foreground(ColorTextMuted).
    MarginTop(1).
    BorderTop(true).
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(ColorBorder).
    PaddingTop(1)
```

---

## Status Components

### Status Bar / Footer

```go
// ============================================================================
// STATUS COMPONENTS
// ============================================================================

// Footer bar at bottom of screen
var StyleFooter = lipgloss.NewStyle().
    Background(ColorCard).
    Foreground(ColorTextSecondary).
    Padding(0, 1)

// Footer item (label + key)
func FooterItem(key string, label string) string {
    keyStyle := lipgloss.NewStyle().
        Foreground(ColorCyan).
        Bold(true)

    labelStyle := lipgloss.NewStyle().
        Foreground(ColorTextSecondary)

    return keyStyle.Render(key) + " " + labelStyle.Render(label)
}

// Active/current footer item
func FooterItemActive(label string) string {
    return lipgloss.NewStyle().
        Background(ColorCyan).
        Foreground(ColorBackground).
        Bold(true).
        Padding(0, 1).
        Render(label)
}
```

### Badges/Tags

```go
// Version badge (like "v1.0.0")
var StyleBadgeVersion = lipgloss.NewStyle().
    Background(ColorBrand).
    Foreground(ColorTextPrimary).
    Bold(true).
    Padding(0, 1)

// Info badge (like "Internal Tool")
var StyleBadgeInfo = lipgloss.NewStyle().
    Border(BorderSingle).
    BorderForeground(ColorBorder).
    Foreground(ColorTextSecondary).
    Padding(0, 1)

// Status badges
func StatusBadge(label string, status string) string {
    var bgColor lipgloss.TerminalColor
    switch status {
    case "success":
        bgColor = ColorGreen
    case "warning":
        bgColor = ColorYellow
    case "error":
        bgColor = ColorDestructive
    default:
        bgColor = ColorSecondary
    }

    return lipgloss.NewStyle().
        Background(bgColor).
        Foreground(ColorBackground).
        Bold(true).
        Padding(0, 1).
        Render(label)
}
```

---

## Header Component

```go
// ============================================================================
// HEADER COMPONENT
// ============================================================================

// Main application header bar
var StyleHeader = lipgloss.NewStyle().
    BorderBottom(true).
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(ColorBrand).
    Padding(0, 1).
    MarginBottom(1)

// Header title (left side)
var StyleHeaderTitle = lipgloss.NewStyle().
    Foreground(ColorTextPrimary)

// Header subtitle (like "Terminal Orchestration Tool")
var StyleHeaderSubtitle = lipgloss.NewStyle().
    Foreground(ColorCyan)

// Header right info (time, context)
var StyleHeaderInfo = lipgloss.NewStyle().
    Foreground(ColorTextSecondary).
    Align(lipgloss.Right)
```

---

## Content Area Components

### Welcome View

```go
// ============================================================================
// CONTENT COMPONENTS
// ============================================================================

// Main content area
var StyleContent = lipgloss.NewStyle().
    Border(BorderSingle).
    BorderForeground(ColorBorder).
    Padding(1, 2)

// Welcome ASCII art container
var StyleASCIIArt = lipgloss.NewStyle().
    Foreground(ColorBrand).   // DB Red for "RFZ" letters
    Align(lipgloss.Center)

// The "CLI" part uses cyan
var StyleASCIIArtCyan = lipgloss.NewStyle().
    Foreground(ColorCyan).
    Align(lipgloss.Center)

// Tagline/quote
var StyleTagline = lipgloss.NewStyle().
    Foreground(ColorTextMuted).
    Italic(true).
    Align(lipgloss.Center).
    MarginTop(1)

// Ready status indicator
var StyleReadyStatus = lipgloss.NewStyle().
    Foreground(ColorTextPrimary).
    Align(lipgloss.Center).
    MarginTop(2)

// Dollar sign prefix for commands
var StylePrompt = lipgloss.NewStyle().
    Foreground(ColorYellow).
    Bold(true)
```

---

## Form Components

### Input Fields

```go
// ============================================================================
// FORM COMPONENTS
// ============================================================================

// Text input field
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

    StyleInputDisabled = lipgloss.NewStyle().
        Border(BorderSingle).
        BorderForeground(ColorBorder).
        Foreground(ColorTextDisabled).
        Padding(0, 1)
)

// Input label
var StyleInputLabel = lipgloss.NewStyle().
    Foreground(ColorTextSecondary).
    MarginBottom(0)

// Input helper text
var StyleInputHelper = lipgloss.NewStyle().
    Foreground(ColorTextMuted).
    Italic(true).
    MarginTop(0)

// Input error message
var StyleInputErrorText = lipgloss.NewStyle().
    Foreground(ColorDestructive).
    MarginTop(0)
```

### Buttons

```go
// Button variants
var (
    // Primary button (filled cyan)
    StyleButtonPrimary = lipgloss.NewStyle().
        Background(ColorCyan).
        Foreground(ColorBackground).
        Bold(true).
        Padding(0, 2)

    // Secondary button (outlined)
    StyleButtonSecondary = lipgloss.NewStyle().
        Border(BorderSingle).
        BorderForeground(ColorBorder).
        Foreground(ColorTextPrimary).
        Padding(0, 2)

    // Destructive button (for dangerous actions)
    StyleButtonDestructive = lipgloss.NewStyle().
        Background(ColorDestructive).
        Foreground(ColorTextPrimary).
        Bold(true).
        Padding(0, 2)

    // Ghost button (minimal)
    StyleButtonGhost = lipgloss.NewStyle().
        Foreground(ColorCyan).
        Padding(0, 1)
)

// Button focused state
func ButtonFocused(style lipgloss.Style) lipgloss.Style {
    return style.Bold(true).Underline(true)
}
```

---

## List Components

### Selection List

```go
// ============================================================================
// LIST COMPONENTS
// ============================================================================

// List item
var StyleListItem = lipgloss.NewStyle().
    Foreground(ColorTextPrimary).
    PaddingLeft(2)

// Selected list item
var StyleListItemSelected = lipgloss.NewStyle().
    Background(ColorSecondary).
    Foreground(ColorTextPrimary).
    Bold(true).
    PaddingLeft(2).
    PaddingRight(2)

// List cursor indicator
var StyleListCursor = lipgloss.NewStyle().
    Foreground(ColorCyan).
    Bold(true)

// Example list rendering
func ListItem(label string, selected bool, cursor bool) string {
    prefix := "  "
    if cursor {
        prefix = lipgloss.NewStyle().Foreground(ColorCyan).Render("> ")
    }

    var style lipgloss.Style
    if selected {
        style = StyleListItemSelected
    } else {
        style = StyleListItem
    }

    return prefix + style.Render(label)
}
```

---

## Log/Output Components

```go
// ============================================================================
// LOG COMPONENTS
// ============================================================================

// Log levels
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
)

// Log timestamp
var StyleLogTimestamp = lipgloss.NewStyle().
    Foreground(ColorTextMuted)

// Log message
var StyleLogMessage = lipgloss.NewStyle().
    Foreground(ColorTextPrimary)

// Render log line
func LogLine(level string, timestamp string, message string) string {
    var levelStyle lipgloss.Style
    switch level {
    case "INFO":
        levelStyle = StyleLogInfo
    case "SUCCESS":
        levelStyle = StyleLogSuccess
    case "WARN":
        levelStyle = StyleLogWarning
    case "ERROR":
        levelStyle = StyleLogError
    default:
        levelStyle = StyleLogDebug
    }

    return lipgloss.JoinHorizontal(
        lipgloss.Top,
        StyleLogTimestamp.Render(timestamp),
        " ",
        levelStyle.Render(fmt.Sprintf("[%-7s]", level)),
        " ",
        StyleLogMessage.Render(message),
    )
}
```

---

## Complete Style Package Structure

```go
// styles/styles.go - Main package file
package styles

import "github.com/charmbracelet/lipgloss"

// Re-export all styles for easy access
// Usage: styles.BoxDefault.Render("content")
```

### Recommended File Organization

```
internal/
  ui/
    styles/
      colors.go      // Color tokens
      typography.go  // Text styles
      borders.go     // Border styles
      components.go  // Component styles
      helpers.go     // Helper functions
      styles.go      // Main export file
```

---

## Bubbles Component Integration

### Using with Bubbles Library

```go
import (
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/bubbles/viewport"
)

// Configure textinput with RFZ styles
func NewStyledTextInput() textinput.Model {
    ti := textinput.New()
    ti.Prompt = StylePrompt.Render("$ ")
    ti.TextStyle = StyleBody
    ti.PlaceholderStyle = StyleBodyMuted
    ti.Cursor.Style = lipgloss.NewStyle().Foreground(ColorCyan)
    return ti
}

// Configure list with RFZ styles
func NewStyledList(items []list.Item, delegate list.ItemDelegate) list.Model {
    l := list.New(items, delegate, 0, 0)
    l.Styles.Title = StyleH2
    l.Styles.FilterPrompt = StyleBody
    l.Styles.FilterCursor = lipgloss.NewStyle().Foreground(ColorCyan)
    return l
}

// Configure viewport with RFZ styles
func NewStyledViewport(width, height int) viewport.Model {
    vp := viewport.New(width, height)
    vp.Style = StyleContent
    return vp
}
```

---

## Accessibility Notes

### Terminal Contrast

- Primary text (#f4f4f5) on background (#1e1e2e): High contrast
- Muted text (#71717a) on background: Meets WCAG AA for large text
- Cyan (#0891b2) on background: Good visibility

### Focus Indicators

- All focusable elements use cyan border highlight
- Cursor prefix (">") provides additional visual indicator
- Bold text for focused items

### Color Blindness Considerations

- Success (green) and Error (red) are also differentiated by context
- Warning uses yellow which is distinct from red/green
- Consider adding text labels alongside color indicators

---

## Screenshots Reference

Reference screenshots stored in:
- `references/prototype-screenshots/95-design-system-top.png` - Color tokens, TuiBox borders
- `references/prototype-screenshots/96-design-system-middle.png` - Navigation components
- `references/prototype-screenshots/97-design-system-bottom.png` - Additional components
- `references/prototype-screenshots/01-welcome-default.png` - Overall layout reference

---

## Notes and Assumptions

### Extracted Values

- Colors extracted visually from screenshots with color picker estimation
- Border styles mapped to closest lipgloss.Border equivalents
- Spacing values based on visual proportions (terminal characters)

### Assumptions Made

- Terminal uses JetBrains Mono or similar monospace font
- Terminal supports ANSI 256 colors or true color
- Dark theme is primary; light theme support via AdaptiveColor

### Areas for Clarification

- Exact hex values should be verified against design source
- Animation/transition timing (not applicable in TUI)
- Responsive breakpoints for different terminal sizes

---

## Quick Reference Card

```
COLORS:
  Background: #1e1e2e    Card: #2a2a3e        Border: #4a4a5e
  Cyan:       #0891b2    Green: #22c55e       Yellow: #eab308
  Red:        #ef4444    Brand: #ec0016       Text: #f4f4f5

BORDERS:
  Normal:  lipgloss.NormalBorder()   Focused: BorderForeground(ColorCyan)
  Double:  lipgloss.DoubleBorder()
  Rounded: lipgloss.RoundedBorder()
  Thick:   lipgloss.ThickBorder()

SPACING:
  XS=1  SM=2  MD=3  LG=4  XL=6  2XL=8

COMMON PATTERNS:
  Box:    Border(NormalBorder()).BorderForeground(ColorBorder).Padding(1,2)
  Focus:  .BorderForeground(ColorCyan)
  Nav:    "> " prefix + Bold + Cyan for focused
```
