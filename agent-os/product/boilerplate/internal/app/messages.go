// Package app contains shared message types used across the application.
//
// Messages are typed structs used for component communication in Bubble Tea.
// This file defines messages that are shared between screens.
// Screen-specific messages should be defined in their respective packages.
package app

// NavigateToScreenMsg requests navigation to a different screen.
// Sent by child screens when they need to trigger navigation.
type NavigateToScreenMsg struct {
	Screen Screen
}

// ShowModalMsg requests display of a modal dialog.
// The app model handles modal overlay rendering.
type ShowModalMsg struct {
	Modal ModalType
}

// HideModalMsg requests hiding the current modal.
type HideModalMsg struct{}

// ErrorMsg represents an error that occurred during an operation.
// Can be displayed in a notification or error modal.
type ErrorMsg struct {
	Err     error
	Context string // Additional context about what operation failed
}

// Error implements the error interface.
func (e ErrorMsg) Error() string {
	if e.Context != "" {
		return e.Context + ": " + e.Err.Error()
	}
	return e.Err.Error()
}

// NotificationMsg requests display of a transient notification.
type NotificationMsg struct {
	Message string
	Type    NotificationType
}

// NotificationType defines the visual style of a notification.
type NotificationType int

const (
	NotificationInfo NotificationType = iota
	NotificationSuccess
	NotificationWarning
	NotificationError
)

// RefreshDataMsg requests a refresh of data from external sources.
// Screens can respond to this by re-fetching their data.
type RefreshDataMsg struct{}
