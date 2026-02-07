# Aufwandsschaetzung: Welcome & Navigation Screen

**Erstellt:** 2026-02-07
**Spec:** 2026-02-07-welcome-navigation-screen
**Anzahl Stories:** 6 (regulaer) + 3 (System)

---

## Zusammenfassung

| Metrik | Human-only | Human + KI Agent | Ersparnis |
|--------|------------|------------------|-----------|
| **Stunden** | 24h | 6.6h | 17.4h (73%) |
| **Arbeitstage** | 3.0d | 0.8d | 2.2d |
| **Arbeitswochen** | 0.6w | 0.2w | 0.4w |

### Was bedeutet das?

**Human-only:** So lange wuerde die Implementierung dauern, wenn ein Entwickler komplett manuell arbeitet (ohne KI-Unterstuetzung).

**Human + KI Agent:** So lange dauert es realistisch mit modernen KI-Werkzeugen (Claude Code, Cursor, GitHub Copilot, etc.). Der Entwickler bleibt verantwortlich fuer Architektur, Code-Review und Qualitaetssicherung.

---

## Schaetzung pro Story

| ID | Story | Komplexitaet | Human (h) | KI-Faktor | KI-Adjusted (h) | Ersparnis |
|----|-------|-------------|-----------|-----------|-----------------|-----------|
| WELCOME-001 | Entry Point & Demo Rename | XS | 3h | high (0.20) | 0.6h | 2.4h |
| WELCOME-002 | App Shell Model with Layout | S | 6h | high (0.20) | 1.2h | 4.8h |
| WELCOME-003 | Welcome Screen | XS | 3h | high (0.20) | 0.6h | 2.4h |
| WELCOME-004 | Screen Switching & Navigation | S | 6h | medium (0.40) | 2.4h | 3.6h |
| WELCOME-005 | Exit Confirmation Modal | XS | 3h | high (0.20) | 0.6h | 2.4h |
| WELCOME-006 | Visual Regression Tests | S | 6h | high (0.20) | 1.2h | 4.8h |
| **TOTAL** | | | **27h** | | **6.6h** | **20.4h** |

*Hinweis: System Stories (997-999) erfordern keinen manuellen Aufwand und werden automatisch ausgefuehrt.*

---

## KI-Beschleunigung nach Kategorie

| Kategorie | Stories | Human (h) | KI-Adjusted (h) | Reduktion |
|-----------|---------|-----------|-----------------|-----------|
| **High** (80% schneller) | 5 | 21h | 4.2h | -80% |
| **Medium** (60% schneller) | 1 | 6h | 2.4h | -60% |
| **Low** (30% schneller) | 0 | 0h | 0h | - |
| **None** (keine Beschleunigung) | 0 | 0h | 0h | - |

### Erklaerung der Kategorien

- **High (Faktor 0.20):** Entry Point, App Shell Layout, Welcome Screen, Exit Modal, Tests - dies sind strukturierte Aufgaben mit klaren Patterns und Vorlagen die KI 5x schneller umsetzen kann
- **Medium (Faktor 0.40):** Screen Switching & Navigation - erfordert etwas mehr Ueberlegung bei Focus-Management und Keyboard-Routing, KI hilft 2.5x schneller

---

## Annahmen & Hinweise

- Schaetzungen basieren auf der Komplexitaetsbewertung des Architects
- KI-Faktoren setzen aktive Nutzung von AI-Tools voraus (Claude Code, Cursor, etc.)
- Qualitaetssicherung und Code-Review bleiben unveraendert wichtig
- Unvorhergesehene Probleme koennen Aufwand erhoehen (+20-30% Puffer empfohlen)
- Alle Stories sind Frontend-only mit bekannten charm.land Patterns

---

## Empfehlung

**Geplanter Aufwand:** 6.6h (0.8d / 0.2w)
**Mit Puffer (+25%):** 8.3h (1.0d / 0.2w)

---

*Erstellt mit Agent OS /create-spec v3.0*
