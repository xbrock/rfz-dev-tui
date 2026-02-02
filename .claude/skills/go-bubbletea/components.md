# Bubbles Component Usage

Reference for using charmbracelet/bubbles components in RFZ Developer CLI.

## Available Components

### List (bubbles/list)

**Usage:** Component selection, navigation menus

```go
import "github.com/charmbracelet/bubbles/list"

// Item interface implementation
type componentItem struct {
    id   string
    name string
    desc string
}

func (i componentItem) Title() string       { return i.name }
func (i componentItem) Description() string { return i.desc }
func (i componentItem) FilterValue() string { return i.name }

// Create list
items := []list.Item{
    componentItem{id: "boss", name: "Boss", desc: "Core component"},
    componentItem{id: "fistiv", name: "Fistiv", desc: "Simulation"},
}

delegate := list.NewDefaultDelegate()
l := list.New(items, delegate, width, height)
l.Title = "Select Components"
l.SetShowStatusBar(true)
l.SetFilteringEnabled(true)
```

### Table (bubbles/table)

**Usage:** Component registry, detected components

```go
import "github.com/charmbracelet/bubbles/table"

columns := []table.Column{
    {Title: "ID", Width: 10},
    {Title: "Name", Width: 20},
    {Title: "Type", Width: 15},
    {Title: "Status", Width: 10},
}

rows := []table.Row{
    {"1", "boss", "Core", "Ready"},
    {"2", "fistiv", "Simulation", "Pending"},
}

t := table.New(
    table.WithColumns(columns),
    table.WithRows(rows),
    table.WithFocused(true),
    table.WithHeight(10),
)

// Style the table
s := table.DefaultStyles()
s.Header = s.Header.
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("#4a4a5e")).
    BorderBottom(true).
    Bold(true)
s.Selected = s.Selected.
    Foreground(lipgloss.Color("#0891b2")).
    Bold(true)
t.SetStyles(s)
```

### Viewport (bubbles/viewport)

**Usage:** Log viewer, scrollable content

```go
import "github.com/charmbracelet/bubbles/viewport"

vp := viewport.New(width, height)
vp.Style = lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("#4a4a5e"))

// Set content
vp.SetContent(strings.Join(logLines, "\n"))

// Enable word wrapping
vp.HighPerformanceRendering = false

// In Update
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    m.viewport, cmd = m.viewport.Update(msg)
    return m, cmd
}
```

### Progress (bubbles/progress)

**Usage:** Build progress indicators

```go
import "github.com/charmbracelet/bubbles/progress"

// Create with gradient
p := progress.New(progress.WithDefaultGradient())

// Or with solid color
p := progress.New(progress.WithSolidFill("#0891b2"))

// Or with custom width
p := progress.New(
    progress.WithWidth(40),
    progress.WithDefaultGradient(),
)

// Update percentage
cmd := p.SetPercent(0.65)

// In View
view := p.View()
```

### Spinner (bubbles/spinner)

**Usage:** Loading states

```go
import "github.com/charmbracelet/bubbles/spinner"

s := spinner.New()
s.Spinner = spinner.Dot  // or spinner.Line, spinner.MiniDot
s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#0891b2"))

// Must init the spinner
func (m Model) Init() tea.Cmd {
    return m.spinner.Tick
}

// In Update, forward tick messages
case spinner.TickMsg:
    var cmd tea.Cmd
    m.spinner, cmd = m.spinner.Update(msg)
    return m, cmd

// In View
if m.loading {
    return m.spinner.View() + " Loading..."
}
```

### TextInput (bubbles/textinput)

**Usage:** Configuration inputs, search

```go
import "github.com/charmbracelet/bubbles/textinput"

ti := textinput.New()
ti.Placeholder = "Enter path..."
ti.CharLimit = 256
ti.Width = 40
ti.Focus()

// Style
ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0891b2"))
ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f4f4f5"))
ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#71717a"))

// In Update
case tea.KeyMsg:
    switch msg.String() {
    case "enter":
        value := m.input.Value()
        // process value
    }
var cmd tea.Cmd
m.input, cmd = m.input.Update(msg)
return m, cmd
```

### Help (bubbles/help)

**Usage:** Keyboard shortcut display

```go
import "github.com/charmbracelet/bubbles/help"
import "github.com/charmbracelet/bubbles/key"

// Define key bindings
type keyMap struct {
    Up    key.Binding
    Down  key.Binding
    Enter key.Binding
    Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
    return []key.Binding{k.Up, k.Down, k.Enter, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {k.Up, k.Down},
        {k.Enter, k.Quit},
    }
}

var keys = keyMap{
    Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("k", "up")),
    Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("j", "down")),
    Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
    Quit:  key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
}

// Create help model
h := help.New()
h.Styles.ShortKey = lipgloss.NewStyle().Foreground(lipgloss.Color("#0891b2"))
h.Styles.ShortDesc = lipgloss.NewStyle().Foreground(lipgloss.Color("#a1a1aa"))

// In View
helpView := h.View(keys)
```

## Component Integration Tips

### Sizing Components

```go
// Handle WindowSizeMsg to resize components
case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height

    // Account for borders and padding
    contentWidth := msg.Width - 4   // 2 border + 2 padding each side
    contentHeight := msg.Height - 6 // header + footer

    m.list.SetSize(contentWidth, contentHeight)
    m.viewport.Width = contentWidth
    m.viewport.Height = contentHeight
```

### Delegating Updates

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Handle global messages first
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "q" {
            return m, tea.Quit
        }
    }

    // Delegate to focused component
    var cmd tea.Cmd
    switch m.focused {
    case FocusList:
        m.list, cmd = m.list.Update(msg)
    case FocusViewport:
        m.viewport, cmd = m.viewport.Update(msg)
    }
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}
```
