// Package app contains global key bindings for the application.
//
// Key bindings are organized into:
// - Global bindings (quit, navigation) - handled in app.go
// - Screen-specific bindings - defined in each screen's keys.go
package app

import "github.com/charmbracelet/bubbles/key"

// keyMap defines global key bindings for the application.
type keyMap struct {
	Quit     key.Binding
	Help     key.Binding
	NavUp    key.Binding
	NavDown  key.Binding
	Select   key.Binding
	Back     key.Binding
	Nav1     key.Binding
	Nav2     key.Binding
	Nav3     key.Binding
	Nav4     key.Binding
	Nav5     key.Binding
}

// defaultKeyMap returns the default key bindings.
func defaultKeyMap() keyMap {
	return keyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		NavUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("up/k", "move up"),
		),
		NavDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("down/j", "move down"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "back"),
		),
		Nav1: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("1", "welcome"),
		),
		Nav2: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("2", "build"),
		),
		Nav3: key.NewBinding(
			key.WithKeys("3"),
			key.WithHelp("3", "logs"),
		),
		Nav4: key.NewBinding(
			key.WithKeys("4"),
			key.WithHelp("4", "discover"),
		),
		Nav5: key.NewBinding(
			key.WithKeys("5"),
			key.WithHelp("5", "config"),
		),
	}
}

// ShortHelp returns key bindings for the short help view.
// Implements bubbles/help.KeyMap interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NavUp, k.NavDown, k.Select, k.Help, k.Quit}
}

// FullHelp returns key bindings for the full help view.
// Implements bubbles/help.KeyMap interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NavUp, k.NavDown, k.Select, k.Back},
		{k.Nav1, k.Nav2, k.Nav3, k.Nav4, k.Nav5},
		{k.Help, k.Quit},
	}
}
