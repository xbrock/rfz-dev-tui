package domain

// ComponentType categorizes RFZ components.
type ComponentType int

const (
	ComponentTypeCore       ComponentType = iota // Core system components
	ComponentTypeSimulation                      // Simulation components
	ComponentTypeStandalone                      // Standalone tools
)

func (t ComponentType) String() string {
	switch t {
	case ComponentTypeCore:
		return "Core"
	case ComponentTypeSimulation:
		return "Simulation"
	case ComponentTypeStandalone:
		return "Standalone"
	default:
		return "Unknown"
	}
}

// Component represents an RFZ software component.
type Component struct {
	Name string
	Type ComponentType
}

// ComponentProvider abstracts component discovery.
// Mock implementation used now; replaced with real file-system scanner in Phase 3.
type ComponentProvider interface {
	Components() []Component
}
