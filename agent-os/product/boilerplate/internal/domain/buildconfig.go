// Package domain contains the core domain entities and value objects.
//
// This file defines the BuildConfig value object which represents
// the configuration for a Maven build operation.
package domain

// BuildConfig represents the configuration for a build operation.
//
// This is a value object - it has no identity beyond its attributes
// and should be treated as immutable.
type BuildConfig struct {
	// Components is the list of component IDs to build.
	Components []string

	// Profiles is the list of Maven profiles to activate.
	Profiles []string

	// SkipTests indicates whether to skip running tests.
	SkipTests bool

	// CleanBuild indicates whether to run a clean build.
	CleanBuild bool

	// Parallelism is the number of parallel builds (-T option).
	Parallelism int

	// Offline indicates whether to run in offline mode.
	Offline bool

	// UpdateSnapshots indicates whether to update snapshots.
	UpdateSnapshots bool

	// CustomArgs contains any additional Maven arguments.
	CustomArgs []string
}

// DefaultBuildConfig returns a BuildConfig with sensible defaults.
func DefaultBuildConfig() BuildConfig {
	return BuildConfig{
		Components:  make([]string, 0),
		Profiles:    []string{"base"},
		SkipTests:   false,
		CleanBuild:  true,
		Parallelism: 4,
		Offline:     false,
		CustomArgs:  make([]string, 0),
	}
}

// WithComponents returns a new BuildConfig with the given components.
func (c BuildConfig) WithComponents(components []string) BuildConfig {
	c.Components = components
	return c
}

// WithProfiles returns a new BuildConfig with the given profiles.
func (c BuildConfig) WithProfiles(profiles []string) BuildConfig {
	c.Profiles = profiles
	return c
}

// WithSkipTests returns a new BuildConfig with skip tests enabled/disabled.
func (c BuildConfig) WithSkipTests(skip bool) BuildConfig {
	c.SkipTests = skip
	return c
}

// ToMavenArgs converts the configuration to Maven command-line arguments.
func (c BuildConfig) ToMavenArgs() []string {
	args := make([]string, 0)

	// Clean goal
	if c.CleanBuild {
		args = append(args, "clean")
	}

	// Install goal
	args = append(args, "install")

	// Profiles
	if len(c.Profiles) > 0 {
		profileArg := "-P"
		for i, p := range c.Profiles {
			if i > 0 {
				profileArg += ","
			}
			profileArg += p
		}
		args = append(args, profileArg)
	}

	// Skip tests
	if c.SkipTests {
		args = append(args, "-DskipTests")
	}

	// Parallelism
	if c.Parallelism > 1 {
		args = append(args, "-T", string(rune('0'+c.Parallelism)))
	}

	// Offline mode
	if c.Offline {
		args = append(args, "-o")
	}

	// Update snapshots
	if c.UpdateSnapshots {
		args = append(args, "-U")
	}

	// Custom args
	args = append(args, c.CustomArgs...)

	return args
}
