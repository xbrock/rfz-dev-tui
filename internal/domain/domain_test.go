package domain

import "testing"

// Scenario 1: Mock-Provider liefert alle RFZ-Komponenten
func TestMockProviderReturns13Components(t *testing.T) {
	provider := MockComponentProvider{}
	components := provider.Components()
	if got := len(components); got != 13 {
		t.Errorf("expected 13 components, got %d", got)
	}
}

// Scenario 1: jede Komponente hat einen Namen und einen Typ
func TestComponentsHaveNameAndType(t *testing.T) {
	provider := MockComponentProvider{}
	for _, c := range provider.Components() {
		if c.Name == "" {
			t.Error("component has empty name")
		}
		if c.Type.String() == "Unknown" {
			t.Errorf("component %q has unknown type", c.Name)
		}
	}
}

// Scenario 2: Komponenten haben den korrekten Typ
func TestComponentTypes(t *testing.T) {
	tests := []struct {
		name     string
		wantType ComponentType
	}{
		{"boss", ComponentTypeCore},
		{"fistiv", ComponentTypeCore},
		{"audiocon", ComponentTypeStandalone},
		{"traktion", ComponentTypeCore},
		{"signalsteuerung", ComponentTypeCore},
		{"weichensteuerung", ComponentTypeCore},
		{"simkern", ComponentTypeSimulation},
		{"fahrdynamik", ComponentTypeSimulation},
		{"energierechnung", ComponentTypeSimulation},
		{"zuglauf", ComponentTypeSimulation},
		{"stellwerk", ComponentTypeCore},
		{"diagnose", ComponentTypeStandalone},
		{"konfiguration", ComponentTypeStandalone},
	}

	provider := MockComponentProvider{}
	components := provider.Components()
	byName := make(map[string]Component, len(components))
	for _, c := range components {
		byName[c.Name] = c
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, ok := byName[tt.name]
			if !ok {
				t.Fatalf("component %q not found", tt.name)
			}
			if c.Type != tt.wantType {
				t.Errorf("component %q: want type %s, got %s", tt.name, tt.wantType, c.Type)
			}
		})
	}
}

// Scenario 3: Standard-Build-Konfiguration erzeugen
func TestBuildConfigToCommand(t *testing.T) {
	cfg := BuildConfig{
		Goal:      GoalCleanInstall,
		Profiles:  []string{"generate_local_config_files"},
		Port:      11090,
		SkipTests: true,
	}
	want := "mvn clean install -Pgenerate_local_config_files,use_traktion_11090 -DskipTests"
	if got := cfg.ToCommand(); got != want {
		t.Errorf("ToCommand()\n got: %s\nwant: %s", got, want)
	}
}

// Edge case: Build-Konfiguration ohne Profile
func TestBuildConfigWithoutProfiles(t *testing.T) {
	cfg := BuildConfig{
		Goal: GoalInstall,
	}
	want := "mvn install"
	if got := cfg.ToCommand(); got != want {
		t.Errorf("ToCommand()\n got: %s\nwant: %s", got, want)
	}
}

// Scenario 4: Build-Phasen sind vollstaendig definiert
func TestBuildPhasesComplete(t *testing.T) {
	phases := []struct {
		phase BuildPhase
		name  string
	}{
		{PhasePending, "Pending"},
		{PhaseCompiling, "Compiling"},
		{PhaseTesting, "Testing"},
		{PhasePackaging, "Packaging"},
		{PhaseInstalling, "Installing"},
		{PhaseDone, "Done"},
		{PhaseFailed, "Failed"},
	}
	for _, tt := range phases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.phase.String(); got != tt.name {
				t.Errorf("phase %d: want %q, got %q", tt.phase, tt.name, got)
			}
		})
	}
}

// Test ComponentType.String()
func TestComponentTypeString(t *testing.T) {
	tests := []struct {
		ct   ComponentType
		want string
	}{
		{ComponentTypeCore, "Core"},
		{ComponentTypeSimulation, "Simulation"},
		{ComponentTypeStandalone, "Standalone"},
		{ComponentType(99), "Unknown"},
	}
	for _, tt := range tests {
		if got := tt.ct.String(); got != tt.want {
			t.Errorf("ComponentType(%d).String() = %q, want %q", tt.ct, got, tt.want)
		}
	}
}

// Test that MockComponentProvider satisfies ComponentProvider interface
func TestMockProviderImplementsInterface(t *testing.T) {
	var _ ComponentProvider = MockComponentProvider{}
}

// Additional ToCommand tests
func TestBuildConfigOnlyPort(t *testing.T) {
	cfg := BuildConfig{
		Goal: GoalCompile,
		Port: 8080,
	}
	want := "mvn compile -Puse_traktion_8080"
	if got := cfg.ToCommand(); got != want {
		t.Errorf("ToCommand()\n got: %s\nwant: %s", got, want)
	}
}

func TestBuildConfigSkipTestsOnly(t *testing.T) {
	cfg := BuildConfig{
		Goal:      GoalTest,
		SkipTests: true,
	}
	want := "mvn test -DskipTests"
	if got := cfg.ToCommand(); got != want {
		t.Errorf("ToCommand()\n got: %s\nwant: %s", got, want)
	}
}
