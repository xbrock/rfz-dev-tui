package build

import (
	"time"

	"rfz-cli/internal/domain"
)

// TestExecutingState returns a Model in the executing phase with deterministic build states.
// This is exported for use in golden file tests at the app level.
func TestExecutingState(width, height, termW, termH int) Model {
	m := New(width, height)
	m.termW = termW
	m.termH = termH
	m.phase = phaseExecuting
	m.selectedComponents = []string{"boss", "konfiguration", "strecke"}
	m.config = domain.BuildConfig{
		Goal:      domain.GoalCleanInstall,
		Profiles:  []string{"target_env_dev"},
		Port:      11090,
		SkipTests: true,
	}
	m.buildStates = []componentBuildState{
		{Name: "boss", Phase: domain.PhasePackaging, StartTime: time.Time{}, Elapsed: 12 * time.Second, Progress: 0.6},
		{Name: "konfiguration", Phase: domain.PhaseCompiling, StartTime: time.Time{}, Elapsed: 5 * time.Second, Progress: 0.3},
		{Name: "strecke", Phase: domain.PhasePending, StartTime: time.Time{}, Elapsed: 0, Progress: 0},
	}
	m.buildCursor = 0
	return m
}

// TestCompletedState returns a Model in the completed phase with deterministic build states.
// This is exported for use in golden file tests at the app level.
func TestCompletedState(width, height, termW, termH int) Model {
	m := New(width, height)
	m.termW = termW
	m.termH = termH
	m.phase = phaseCompleted
	m.selectedComponents = []string{"boss", "konfiguration", "strecke"}
	m.config = domain.BuildConfig{
		Goal:      domain.GoalCleanInstall,
		Profiles:  []string{"target_env_dev"},
		Port:      11090,
		SkipTests: true,
	}
	m.buildStates = []componentBuildState{
		{Name: "boss", Phase: domain.PhaseDone, StartTime: time.Time{}, Elapsed: 45 * time.Second, Progress: 1},
		{Name: "konfiguration", Phase: domain.PhaseDone, StartTime: time.Time{}, Elapsed: 38 * time.Second, Progress: 1},
		{Name: "strecke", Phase: domain.PhaseFailed, StartTime: time.Time{}, Elapsed: 22 * time.Second, Progress: 1},
	}
	m.buildCursor = 0
	return m
}
