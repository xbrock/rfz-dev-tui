# Aufwandsschätzung: Interactive Components (Sprint 1.2)

**Erstellt:** 2026-02-03
**Spec:** 2026-02-03-interactive-components
**Anzahl Stories:** 11 (8 regular + 3 system)

---

## Zusammenfassung

| Metrik | Human-only | Human + KI Agent | Ersparnis |
|--------|------------|------------------|-----------|
| **Stunden** | 32h | 10h | 22h (69%) |
| **Arbeitstage** | 4d | 1.25d | 2.75d |
| **Arbeitswochen** | 0.8w | 0.25w | 0.55w |

### Was bedeutet das?

**Human-only:** So lange würde die Implementierung dauern, wenn ein Entwickler komplett manuell arbeitet (ohne KI-Unterstützung).

**Human + KI Agent:** So lange dauert es realistisch mit modernen KI-Werkzeugen (Claude Code, Cursor, GitHub Copilot, etc.). Der Entwickler bleibt verantwortlich für Architektur, Code-Review und Qualitätssicherung.

---

## Schätzung pro Story

| ID | Story | Komplexität | Human (h) | KI-Faktor | KI-Adjusted (h) | Ersparnis |
|----|-------|-------------|-----------|-----------|-----------------|-----------|
| INTER-002 | TuiCheckbox | XS | 3h | high | 0.6h | 2.4h |
| INTER-003 | TuiRadio | XS | 3h | high | 0.6h | 2.4h |
| INTER-004 | TuiTextInput | S | 4h | high | 0.8h | 3.2h |
| INTER-005 | TuiSpinner | S | 4h | high | 0.8h | 3.2h |
| INTER-006 | TuiProgress | S | 4h | high | 0.8h | 3.2h |
| INTER-001 | TuiList | M | 6h | medium | 2.4h | 3.6h |
| INTER-007 | Extend Gallery | S | 4h | high | 0.8h | 3.2h |
| INTER-008 | Visual Tests | S | 4h | high | 0.8h | 3.2h |
| INTER-997 | Code Review | S | 2h | low | 1.4h | 0.6h |
| INTER-998 | Integration Validation | XS | 1h | none | 1.0h | 0h |
| INTER-999 | Finalize PR | XS | 1h | none | 1.0h | 0h |
| **TOTAL** | | | **36h** | | **11h** | **25h** |

---

## KI-Beschleunigung nach Kategorie

| Kategorie | Stories | Human (h) | KI-Adjusted (h) | Reduktion |
|-----------|---------|-----------|-----------------|-----------|
| **High** (80% schneller) | 7 | 26h | 5.2h | -80% |
| **Medium** (60% schneller) | 1 | 6h | 2.4h | -60% |
| **Low** (30% schneller) | 1 | 2h | 1.4h | -30% |
| **None** (keine Beschleunigung) | 2 | 2h | 2.0h | 0% |

### Erklärung der Kategorien

- **High (Faktor 0.20):** Stateless render functions (checkbox, radio), Bubbles wrappers (spinner, progress, textinput), gallery sections, test boilerplate - KI kann 5x schneller helfen
- **Medium (Faktor 0.40):** Complex component (TuiList) with state management and integration - KI hilft 2.5x schneller
- **Low (Faktor 0.70):** Code review requiring human judgment - KI unterstützt bei Analyse
- **None (Faktor 1.00):** System tasks (integration validation, PR) - menschliches Urteil erforderlich

---

## Annahmen & Hinweise

- Schätzungen basieren auf der Komplexitätsbewertung des Architects
- KI-Faktoren setzen aktive Nutzung von AI-Tools voraus (Claude Code, Cursor, etc.)
- Qualitätssicherung und Code-Review bleiben unverändert wichtig
- Unvorhergesehene Probleme können Aufwand erhöhen (+20-30% Puffer empfohlen)
- Existing patterns from Sprint 1.1 (box.go, button.go) significantly reduce learning curve

---

## Empfehlung

**Geplanter Aufwand:** 11h (1.4d / 0.3w)
**Mit Puffer (+25%):** 14h (1.75d / 0.35w)

### Parallelisierung

Stories INTER-002 through INTER-006 können parallel ausgeführt werden, was die Gesamtdauer weiter reduzieren kann wenn mehrere Agenten gleichzeitig arbeiten.

---

*Erstellt mit Agent OS /create-spec v2.7*
