# Aufwandsschätzung: Layout & Navigation Components

**Erstellt:** 2026-02-06
**Spec:** 2026-02-06-layout-navigation
**Anzahl Stories:** 12

---

## Zusammenfassung

| Metrik | Human-only | Human + KI Agent | Ersparnis |
|--------|------------|------------------|-----------|
| **Stunden** | 46h | 14h | 32h (70%) |
| **Arbeitstage** | 5.8d | 1.8d | 4.0d |
| **Arbeitswochen** | 1.2w | 0.4w | 0.8w |

### Was bedeutet das?

**Human-only:** So lange würde die Implementierung dauern, wenn ein Entwickler komplett manuell arbeitet (ohne KI-Unterstützung).

**Human + KI Agent:** So lange dauert es realistisch mit modernen KI-Werkzeugen (Claude Code, Cursor, GitHub Copilot, etc.). Der Entwickler bleibt verantwortlich für Architektur, Code-Review und Qualitätssicherung.

---

## Schätzung pro Story

| ID | Story | Komplexität | Human (h) | KI-Faktor | KI-Adjusted (h) | Ersparnis |
|----|-------|-------------|-----------|-----------|-----------------|-----------|
| LAYOUT-001 | TuiNavigation + TuiNavItem | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-002 | TuiModal | M | 12h | medium (0.40) | 4.8h | 7.2h |
| LAYOUT-003 | TuiKeyHints | XS | 3h | high (0.20) | 0.6h | 2.4h |
| LAYOUT-004 | TuiTable | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-005 | TuiTree | M | 12h | medium (0.40) | 4.8h | 7.2h |
| LAYOUT-006 | TuiTabs | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-007 | TuiStatusBar | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-008 | Layout Demo | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-009 | Visual Tests | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-997 | Code Review | S | 4h | none (1.00) | 4.0h | 0h |
| LAYOUT-998 | Integration Validation | XS | 2h | high (0.20) | 0.4h | 1.6h |
| LAYOUT-999 | Finalize PR | XS | 2h | medium (0.40) | 0.8h | 1.2h |
| **TOTAL** | | | **71h** | | **22.6h** | **48.4h** |

---

## KI-Beschleunigung nach Kategorie

| Kategorie | Stories | Human (h) | KI-Adjusted (h) | Reduktion |
|-----------|---------|-----------|-----------------|-----------|
| **High** (80% schneller) | 8 | 41h | 8.2h | -80% |
| **Medium** (60% schneller) | 3 | 26h | 10.4h | -60% |
| **Low** (30% schneller) | 0 | 0h | 0h | -30% |
| **None** (keine Beschleunigung) | 1 | 4h | 4h | 0% |

### Erklärung der Kategorien

- **High (Faktor 0.20):** TUI-Komponenten nach etabliertem Pattern, Tests, Boilerplate - KI kann 5x schneller helfen
- **Medium (Faktor 0.40):** Komplexere Komponenten (Modal, Tree) mit State-Logik - KI hilft 2.5x schneller
- **Low (Faktor 0.70):** Neue Technologien, komplexe Bugs, Architektur - KI hilft 1.4x schneller
- **None (Faktor 1.00):** Code Review durch Opus - menschliches/LLM-Urteil erforderlich, keine Beschleunigung

---

## Annahmen & Hinweise

- Schätzungen basieren auf der Komplexitätsbewertung des Architects
- KI-Faktoren setzen aktive Nutzung von AI-Tools voraus (Claude Code, Cursor, etc.)
- Qualitätssicherung und Code-Review bleiben unverändert wichtig
- Unvorhergesehene Probleme können Aufwand erhöhen (+20-30% Puffer empfohlen)
- **High KI-Faktor:** TUI-Komponenten folgen etablierten Patterns (list.go, box.go), die KI gut generieren kann
- **Medium KI-Faktor:** Modal (Focus-Trapping) und Tree (Rekursion) erfordern mehr manuelles Denken

---

## Empfehlung

**Geplanter Aufwand:** 22.6h (2.8d / 0.6w)
**Mit Puffer (+25%):** 28.3h (3.5d / 0.7w)

---

*Erstellt mit Agent OS /create-spec v3.0*
