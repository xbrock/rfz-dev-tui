// Package app provides the main RFZ CLI application model.
//
// This file contains shared message types used across the application.
package app

import "time"

// TickMsg is sent every second to update the clock display.
type TickMsg time.Time

// NavigateMsg requests navigation to a specific screen index.
type NavigateMsg struct {
	Screen int
}
