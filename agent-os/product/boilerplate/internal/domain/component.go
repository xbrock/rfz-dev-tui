// Package domain contains the core domain entities and value objects.
//
// This file defines the Component entity which represents an RFZ component
// that can be built, discovered, and managed.
package domain

// ComponentType represents the type of RFZ component.
type ComponentType string

const (
	ComponentTypeCore       ComponentType = "core"
	ComponentTypeService    ComponentType = "service"
	ComponentTypeSimulator  ComponentType = "simulator"
	ComponentTypeStandalone ComponentType = "standalone"
	ComponentTypeLibrary    ComponentType = "library"
)

// Component represents an RFZ component entity.
//
// Components are the primary unit of work in the RFZ Developer CLI.
// They can be discovered from the filesystem, added to a registry,
// and built using Maven.
type Component struct {
	// ID is the unique identifier for this component.
	// Format: groupId:artifactId
	ID string

	// Name is the display name of the component.
	Name string

	// Type is the component type (core, service, simulator, etc.)
	Type ComponentType

	// Path is the absolute filesystem path to the component.
	Path string

	// ArtifactID is the Maven artifact ID.
	ArtifactID string

	// GroupID is the Maven group ID.
	GroupID string

	// Version is the component version.
	Version string

	// Dependencies lists the component IDs this component depends on.
	Dependencies []string

	// Registered indicates if this component is in the registry.
	Registered bool

	// LastBuildTime is the timestamp of the last successful build.
	LastBuildTime string

	// LastBuildSuccess indicates if the last build was successful.
	LastBuildSuccess bool
}

// NewComponent creates a new Component with the given properties.
func NewComponent(id, name string, compType ComponentType, path string) Component {
	return Component{
		ID:           id,
		Name:         name,
		Type:         compType,
		Path:         path,
		Dependencies: make([]string, 0),
	}
}

// HasDependency checks if this component depends on the given component.
func (c Component) HasDependency(componentID string) bool {
	for _, dep := range c.Dependencies {
		if dep == componentID {
			return true
		}
	}
	return false
}

// IsCore returns true if this is a core component.
func (c Component) IsCore() bool {
	return c.Type == ComponentTypeCore
}

// IsService returns true if this is a service component.
func (c Component) IsService() bool {
	return c.Type == ComponentTypeService
}
