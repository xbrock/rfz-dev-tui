# Aufwandsschätzung: Core Components

**Erstellt:** 2026-02-02
**Spec:** Core Components (Phase 1, Week 1)
**Anzahl Stories:** 7

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
| CORE-001 | Styles Package | S | 6h | high (0.20) | 1.2h | 4.8h |
| CORE-002 | TuiBox Component | S | 4h | high (0.20) | 0.8h | 3.2h |
| CORE-003 | TuiDivider Component | XS | 3h | high (0.20) | 0.6h | 2.4h |
| CORE-004 | TuiButton Component | S | 4h | high (0.20) | 0.8h | 3.2h |
| CORE-005 | TuiStatus Component | S | 4h | high (0.20) | 0.8h | 3.2h |
| CORE-006 | teatest Infrastructure | S | 6h | medium (0.40) | 2.4h | 3.6h |
| CORE-007 | Component Gallery | S | 5h | medium (0.40) | 2.0h | 3.0h |
| **TOTAL** | | | **32h** | | **8.6h** | **23.4h** |

---

## KI-Beschleunigung nach Kategorie

| Kategorie | Stories | Human (h) | KI-Adjusted (h) | Reduktion |
|-----------|---------|-----------|-----------------|-----------|
| **High** (80% schneller) | 5 | 21h | 4.2h | -80% |
| **Medium** (60% schneller) | 2 | 11h | 4.4h | -60% |
| **Low** (30% schneller) | 0 | 0h | 0h | - |
| **None** (keine Beschleunigung) | 0 | 0h | 0h | - |

### Erklärung der Kategorien

- **High (Faktor 0.20):** Styles Package, Component-Implementierungen - klare Patterns, Boilerplate-heavy, KI kann 5x schneller helfen
- **Medium (Faktor 0.40):** Test-Infrastruktur, Gallery - erfordert mehr Konfiguration und Integration
- **Low (Faktor 0.70):** Nicht vorhanden in dieser Spec
- **None (Faktor 1.00):** Nicht vorhanden in dieser Spec

---

## Annahmen & Hinweise

- Schätzungen basieren auf der Komplexitätsbewertung des Architects
- KI-Faktoren setzen aktive Nutzung von AI-Tools voraus (Claude Code, Cursor, etc.)
- Qualitätssicherung und Code-Review bleiben unverändert wichtig
- Unvorhergesehene Probleme können Aufwand erhöhen (+20-30% Puffer empfohlen)
- Go-Entwickler mit Bubble Tea/Lip Gloss Erfahrung angenommen

---

## Empfehlung

**Geplanter Aufwand:** 10h (1.25 Arbeitstage)
**Mit Puffer (+25%):** 12.5h (1.5 Arbeitstage)

### Optimaler Ablauf mit KI

1. **Tag 1 Vormittag:** CORE-001 Styles Package (~1.5h)
2. **Tag 1 Mittag:** CORE-002 bis CORE-005 parallel (~3h mit 4 parallel tasks)
3. **Tag 1 Nachmittag:** CORE-006 teatest Infrastructure (~2.5h)
4. **Tag 2 Vormittag:** CORE-007 Component Gallery + Final Review (~2.5h)

---

*Erstellt mit Agent OS /create-spec v3.0*
