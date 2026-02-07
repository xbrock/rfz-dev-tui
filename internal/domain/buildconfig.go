package domain

import (
	"fmt"
	"strings"
)

// MavenGoal represents a Maven build goal.
type MavenGoal string

const (
	GoalInstall      MavenGoal = "install"
	GoalCleanInstall MavenGoal = "clean install"
	GoalPackage      MavenGoal = "package"
	GoalCompile      MavenGoal = "compile"
	GoalTest         MavenGoal = "test"
)

// BuildPhase represents a stage in the build lifecycle.
type BuildPhase int

const (
	PhasePending    BuildPhase = iota // Not started
	PhaseCompiling                    // Compiling sources
	PhaseTesting                      // Running tests
	PhasePackaging                    // Packaging artifacts
	PhaseInstalling                   // Installing to local repo
	PhaseDone                         // Build finished successfully
	PhaseFailed                       // Build failed
)

func (p BuildPhase) String() string {
	switch p {
	case PhasePending:
		return "Pending"
	case PhaseCompiling:
		return "Compiling"
	case PhaseTesting:
		return "Testing"
	case PhasePackaging:
		return "Packaging"
	case PhaseInstalling:
		return "Installing"
	case PhaseDone:
		return "Done"
	case PhaseFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

// BuildConfig holds the parameters for a Maven build.
type BuildConfig struct {
	Goal      MavenGoal
	Profiles  []string
	Port      int
	SkipTests bool
}

// ToCommand generates the Maven command string from this configuration.
func (c BuildConfig) ToCommand() string {
	var parts []string
	parts = append(parts, "mvn")
	parts = append(parts, string(c.Goal))

	var profiles []string
	profiles = append(profiles, c.Profiles...)
	if c.Port > 0 {
		profiles = append(profiles, portProfile(c.Port))
	}
	if len(profiles) > 0 {
		parts = append(parts, "-P"+strings.Join(profiles, ","))
	}

	if c.SkipTests {
		parts = append(parts, "-DskipTests")
	}

	return strings.Join(parts, " ")
}

func portProfile(port int) string {
	return fmt.Sprintf("use_traktion_%d", port)
}
