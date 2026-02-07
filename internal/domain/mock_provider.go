package domain

// MockComponentProvider returns hardcoded RFZ components for development and testing.
// Will be replaced by a real file-system scanner in Phase 3.
type MockComponentProvider struct{}

// Components returns the 13 RFZ components matching the prototype.
func (m MockComponentProvider) Components() []Component {
	return []Component{
		{Name: "boss", Type: ComponentTypeCore},
		{Name: "fistiv", Type: ComponentTypeCore},
		{Name: "audiocon", Type: ComponentTypeStandalone},
		{Name: "traktion", Type: ComponentTypeCore},
		{Name: "signalsteuerung", Type: ComponentTypeCore},
		{Name: "weichensteuerung", Type: ComponentTypeCore},
		{Name: "simkern", Type: ComponentTypeSimulation},
		{Name: "fahrdynamik", Type: ComponentTypeSimulation},
		{Name: "energierechnung", Type: ComponentTypeSimulation},
		{Name: "zuglauf", Type: ComponentTypeSimulation},
		{Name: "stellwerk", Type: ComponentTypeCore},
		{Name: "diagnose", Type: ComponentTypeStandalone},
		{Name: "konfiguration", Type: ComponentTypeStandalone},
	}
}
