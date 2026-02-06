// Package main provides the entry point for the layout component demo.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/ui/components/demo"
)

func main() {
	m := demo.NewLayoutGallery()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
