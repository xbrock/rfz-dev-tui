# Lip Gloss Styling Patterns

Styling patterns for RFZ Developer CLI using charmbracelet/lipgloss.

## Color Palette

### From Design System

```go
package styles

import "github.com/charmbracelet/lipgloss"

// Background Colors
var (
    ColorBackground = lipgloss.Color("#1e1e2e") // Main background
    ColorCard       = lipgloss.Color("#2a2a3e") // Card/panel backgrounds
    ColorSecondary  = lipgloss.Color("#3a3a4e") // Secondary surfaces
    ColorBorder     = lipgloss.Color("#4a4a5e") // Default borders
)

// Accent Colors
var (
    ColorCyan        = lipgloss.Color("#0891b2") // Focus, interactive
    ColorGreen       = lipgloss.Color("#22c55e") // Success states
    ColorYellow      = lipgloss.Color("#eab308") // Warnings
    ColorDestructive = lipgloss.Color("#ef4444") // Errors
    ColorBrand       = lipgloss.Color("#ec0016") // DB Red
)

// Text Colors
var (
    ColorTextPrimary   = lipgloss.Color("#f4f4f5") // High contrast
    ColorTextSecondary = lipgloss.Color("#a1a1aa") // Muted
    ColorTextMuted     = lipgloss.Color("#71717a") // Very subtle
    ColorTextDisabled  = lipgloss.Color("#52525b") // Disabled
)
```

## Border Styles

### Available Borders

```go
// Single line (default)
lipgloss.NormalBorder()    // ┌─┐

// Double line (modals)
lipgloss.DoubleBorder()    // ╔═╗

// Rounded (friendly)
lipgloss.RoundedBorder()   // ╭─╮

// Heavy/thick
lipgloss.ThickBorder()     // ┏━┓

// No border
lipgloss.HiddenBorder()
```

### Box Styles

```go
var (
    // Default box
    StyleBoxDefault = lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(ColorBorder).
        Padding(1, 2)

    // Focused box (cyan border)
    StyleBoxFocused = lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(ColorCyan).
        Padding(1, 2)

    // Modal box (double border)
    StyleBoxModal = lipgloss.NewStyle().
        Border(lipgloss.DoubleBorder()).
        BorderForeground(ColorTextPrimary).
        Padding(1, 2)

    // Rounded box
    StyleBoxRounded = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(ColorBorder).
        Padding(1, 2)
)
```

## Layout Functions

### Joining Content

```go
// Horizontal (side by side)
lipgloss.JoinHorizontal(
    lipgloss.Top,     // Align at top
    leftPanel,
    rightPanel,
)

// Vertical (stacked)
lipgloss.JoinVertical(
    lipgloss.Left,    // Align at left
    header,
    content,
    footer,
)
```

### Centering Content

```go
// Center in available space
lipgloss.Place(
    width, height,
    lipgloss.Center, lipgloss.Center,
    content,
)

// Or with style
style := lipgloss.NewStyle().
    Width(width).
    Height(height).
    Align(lipgloss.Center, lipgloss.Center)

centered := style.Render(content)
```

### Fixed Dimensions

```go
// Fixed width
style := lipgloss.NewStyle().Width(40)

// Fixed height
style := lipgloss.NewStyle().Height(10)

// Both
style := lipgloss.NewStyle().Width(40).Height(10)

// Max width/height
style := lipgloss.NewStyle().MaxWidth(80).MaxHeight(20)
```

## Text Styling

### Typography

```go
var (
    // Headings
    StyleH1 = lipgloss.NewStyle().
        Bold(true).
        Foreground(ColorTextPrimary)

    StyleH2 = lipgloss.NewStyle().
        Bold(true).
        Foreground(ColorCyan)

    // Body text
    StyleBody = lipgloss.NewStyle().
        Foreground(ColorTextPrimary)

    StyleBodyMuted = lipgloss.NewStyle().
        Foreground(ColorTextMuted).
        Italic(true)

    // Code
    StyleCode = lipgloss.NewStyle().
        Background(ColorSecondary).
        Padding(0, 1)
)
```

### Text Modifiers

```go
style := lipgloss.NewStyle().
    Bold(true).
    Italic(true).
    Underline(true).
    Strikethrough(true).
    Faint(true).
    Blink(true)
```

## Spacing

### Padding

```go
// All sides
style.Padding(2)

// Vertical, Horizontal
style.Padding(1, 2)  // 1 top/bottom, 2 left/right

// Top, Horizontal, Bottom
style.Padding(1, 2, 1)

// Top, Right, Bottom, Left
style.Padding(1, 2, 1, 2)

// Individual
style.PaddingTop(1)
style.PaddingRight(2)
style.PaddingBottom(1)
style.PaddingLeft(2)
```

### Margin

```go
// Same pattern as padding
style.Margin(2)
style.Margin(1, 2)
style.MarginTop(1)
style.MarginLeft(2)
```

## Status Styles

### Build Status Badges

```go
var (
    StyleStatusPending = lipgloss.NewStyle().
        Foreground(ColorTextMuted)

    StyleStatusRunning = lipgloss.NewStyle().
        Foreground(ColorYellow)

    StyleStatusSuccess = lipgloss.NewStyle().
        Foreground(ColorGreen).
        Bold(true)

    StyleStatusFailed = lipgloss.NewStyle().
        Foreground(ColorDestructive).
        Bold(true)
)

func RenderStatus(status string) string {
    switch status {
    case "pending":
        return StyleStatusPending.Render("o")
    case "running":
        return StyleStatusRunning.Render(".")
    case "success":
        return StyleStatusSuccess.Render("v")
    case "failed":
        return StyleStatusFailed.Render("x")
    }
    return status
}
```

### Log Level Styles

```go
var (
    StyleLogInfo    = lipgloss.NewStyle().Foreground(ColorCyan)
    StyleLogWarn    = lipgloss.NewStyle().Foreground(ColorYellow)
    StyleLogError   = lipgloss.NewStyle().Foreground(ColorDestructive)
    StyleLogDebug   = lipgloss.NewStyle().Foreground(ColorTextMuted)
    StyleLogSuccess = lipgloss.NewStyle().Foreground(ColorGreen)
)
```

## Navigation Styles

```go
var (
    // Normal nav item
    StyleNavItem = lipgloss.NewStyle().
        Foreground(ColorTextPrimary).
        PaddingLeft(2)

    // Focused (cursor on it)
    StyleNavItemFocused = lipgloss.NewStyle().
        Foreground(ColorCyan).
        Bold(true)

    // Active (current screen)
    StyleNavItemActive = lipgloss.NewStyle().
        Background(ColorSecondary).
        Foreground(ColorTextPrimary).
        Bold(true).
        Padding(0, 2)
)
```

## Adaptive Colors

For light/dark terminal support:

```go
var ColorAdaptiveBackground = lipgloss.AdaptiveColor{
    Light: "#ffffff",
    Dark:  "#1e1e2e",
}

var ColorAdaptiveForeground = lipgloss.AdaptiveColor{
    Light: "#1e1e2e",
    Dark:  "#f4f4f5",
}

// Use in style
style := lipgloss.NewStyle().
    Background(ColorAdaptiveBackground).
    Foreground(ColorAdaptiveForeground)
```

## Common Patterns

### Focus Toggle

```go
func (m *Model) getFocusStyle() lipgloss.Style {
    if m.focused {
        return StyleBoxFocused
    }
    return StyleBoxDefault
}
```

### Dynamic Width

```go
func (m Model) View() string {
    // Use available width minus borders
    contentWidth := m.width - 4

    style := lipgloss.NewStyle().
        Width(contentWidth).
        Border(lipgloss.NormalBorder()).
        BorderForeground(ColorBorder)

    return style.Render(m.content)
}
```

### Truncation

```go
// Lip Gloss handles overflow with MaxWidth
style := lipgloss.NewStyle().
    MaxWidth(40)

text := style.Render("Very long text that needs truncation...")
```
