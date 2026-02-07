// Package main is the entry point for the RFZ Component Demo.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/ui/components/demo"
)

func main() {
	p := tea.NewProgram(
		demo.New(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
