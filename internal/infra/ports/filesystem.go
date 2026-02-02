// Package ports defines interfaces (ports) for external dependencies.
//
// This file defines the FileSystem interface for filesystem operations.
package ports

import "io"

// FileInfo represents information about a file.
type FileInfo struct {
	// Name is the file name (not the full path).
	Name string

	// Path is the absolute file path.
	Path string

	// IsDir indicates if this is a directory.
	IsDir bool

	// Size is the file size in bytes.
	Size int64

	// ModTime is the modification time as Unix timestamp.
	ModTime int64
}

// FileSystem defines the interface for filesystem operations.
//
// This port abstracts filesystem access so that:
// - Production code can use real filesystem operations
// - Tests can use mock implementations (in-memory filesystem)
// - Component scanning can be tested without real files
type FileSystem interface {
	// ReadFile reads the contents of a file.
	ReadFile(path string) ([]byte, error)

	// WriteFile writes data to a file.
	WriteFile(path string, data []byte) error

	// Exists checks if a file or directory exists.
	Exists(path string) bool

	// IsDir checks if the path is a directory.
	IsDir(path string) bool

	// ListDir lists the contents of a directory.
	ListDir(path string) ([]FileInfo, error)

	// Walk walks the directory tree rooted at path.
	//
	// The callback is called for each file and directory.
	// Return an error from the callback to stop walking.
	Walk(path string, callback func(path string, info FileInfo) error) error

	// MkdirAll creates a directory and all parent directories.
	MkdirAll(path string) error

	// Remove deletes a file or empty directory.
	Remove(path string) error

	// RemoveAll deletes a directory and all its contents.
	RemoveAll(path string) error

	// Open opens a file for reading.
	Open(path string) (io.ReadCloser, error)

	// Create creates a file for writing.
	Create(path string) (io.WriteCloser, error)
}

// FileScanner defines the interface for scanning for component files.
//
// This is a higher-level abstraction for finding RFZ components.
type FileScanner interface {
	// ScanForPomFiles finds all pom.xml files under the given path.
	ScanForPomFiles(rootPath string) ([]string, error)

	// ParsePomFile parses a pom.xml file and returns component metadata.
	ParsePomFile(pomPath string) (*PomMetadata, error)

	// DetectComponentType determines the component type from metadata.
	DetectComponentType(metadata *PomMetadata) string
}

// PomMetadata contains information extracted from a pom.xml file.
type PomMetadata struct {
	// GroupID is the Maven group ID.
	GroupID string

	// ArtifactID is the Maven artifact ID.
	ArtifactID string

	// Version is the component version.
	Version string

	// Packaging is the Maven packaging type.
	Packaging string

	// Name is the project name.
	Name string

	// Description is the project description.
	Description string

	// Dependencies is the list of dependency coordinates.
	Dependencies []string

	// Properties contains Maven properties.
	Properties map[string]string
}
