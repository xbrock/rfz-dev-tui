package build

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/domain"
)

// simulatorTickInterval is how often the simulator sends tick messages.
const simulatorTickInterval = 100 * time.Millisecond

// phaseDurations defines min/max ticks (at 100ms each) per build phase.
var phaseDurations = map[domain.BuildPhase][2]int{
	domain.PhaseCompiling:  {8, 20},  // 0.8-2.0s
	domain.PhaseTesting:    {10, 25}, // 1.0-2.5s
	domain.PhasePackaging:  {5, 15},  // 0.5-1.5s
	domain.PhaseInstalling: {5, 12},  // 0.5-1.2s
}

// failureChance is the probability of failure during the Testing phase (0-100).
const failureChance = 20

// simulatorState tracks per-component simulation progress.
type simulatorState struct {
	ticksInPhase  int
	ticksRequired int
}

// startSimulator initializes simulation state for all components and returns
// the first tick command.
func startSimulator(count int) ([]simulatorState, tea.Cmd) {
	states := make([]simulatorState, count)
	for i := range states {
		states[i] = simulatorState{
			ticksInPhase:  0,
			ticksRequired: randomDuration(domain.PhaseCompiling),
		}
	}
	return states, simulatorTick()
}

// simulatorTick returns a tea.Cmd that sends a BuildTickMsg after the tick interval.
func simulatorTick() tea.Cmd {
	return tea.Tick(simulatorTickInterval, func(t time.Time) tea.Msg {
		return BuildTickMsg(t)
	})
}

// advanceSimulation processes one tick for all components, returning phase
// transition messages and whether the simulation is complete.
func advanceSimulation(
	buildStates []componentBuildState,
	simStates []simulatorState,
) ([]BuildPhaseMsg, bool) {
	var msgs []BuildPhaseMsg
	allDone := true

	for i := range buildStates {
		phase := buildStates[i].Phase

		// Skip finished components
		if phase == domain.PhaseDone || phase == domain.PhaseFailed || phase == domain.PhasePending {
			if phase == domain.PhasePending {
				allDone = false
			}
			continue
		}

		allDone = false
		simStates[i].ticksInPhase++

		// Update progress within current phase
		if simStates[i].ticksRequired > 0 {
			buildStates[i].Progress = float64(simStates[i].ticksInPhase) / float64(simStates[i].ticksRequired)
			if buildStates[i].Progress > 1 {
				buildStates[i].Progress = 1
			}
		}

		// Check if phase is complete
		if simStates[i].ticksInPhase >= simStates[i].ticksRequired {
			nextPhase := nextBuildPhase(phase)

			// Random failure during testing
			if phase == domain.PhaseTesting && rand.Intn(100) < failureChance { //nolint:gosec // simulation only
				nextPhase = domain.PhaseFailed
			}

			msgs = append(msgs, BuildPhaseMsg{
				ComponentIndex: i,
				Phase:          nextPhase,
			})

			// Reset tick counter for next phase
			simStates[i].ticksInPhase = 0
			if nextPhase != domain.PhaseDone && nextPhase != domain.PhaseFailed {
				simStates[i].ticksRequired = randomDuration(nextPhase)
			}
		}
	}

	return msgs, allDone
}

// startPendingComponents kicks off the first N pending components (staggered start).
// Returns phase transition messages for components that should begin compiling.
func startPendingComponents(buildStates []componentBuildState, maxConcurrent int) []BuildPhaseMsg {
	var msgs []BuildPhaseMsg
	running := 0

	for _, s := range buildStates {
		if s.Phase != domain.PhasePending && s.Phase != domain.PhaseDone && s.Phase != domain.PhaseFailed {
			running++
		}
	}

	for i := range buildStates {
		if running >= maxConcurrent {
			break
		}
		if buildStates[i].Phase == domain.PhasePending {
			msgs = append(msgs, BuildPhaseMsg{
				ComponentIndex: i,
				Phase:          domain.PhaseCompiling,
			})
			running++
		}
	}

	return msgs
}

// nextBuildPhase returns the phase following the given phase.
func nextBuildPhase(p domain.BuildPhase) domain.BuildPhase {
	switch p {
	case domain.PhaseCompiling:
		return domain.PhaseTesting
	case domain.PhaseTesting:
		return domain.PhasePackaging
	case domain.PhasePackaging:
		return domain.PhaseInstalling
	case domain.PhaseInstalling:
		return domain.PhaseDone
	default:
		return domain.PhaseDone
	}
}

// randomDuration returns a random tick count for the given phase.
func randomDuration(phase domain.BuildPhase) int {
	bounds, ok := phaseDurations[phase]
	if !ok {
		return 10
	}
	min, max := bounds[0], bounds[1]
	return min + rand.Intn(max-min+1) //nolint:gosec // simulation only
}
