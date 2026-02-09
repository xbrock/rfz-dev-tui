# Aufwandsschaetzung: Layout Matching Design

**Erstellt:** 2026-02-09
**Spec:** 2026-02-09-bugfix-layout-matching-design
**Anzahl Stories:** 8 (+ 3 System Stories)

---

## Zusammenfassung

| Metrik | Human-only | Human + KI Agent | Ersparnis |
|--------|------------|------------------|-----------|
| **Stunden** | 36h | 9h | 27h (75%) |
| **Arbeitstage** | 4.5d | 1.1d | 3.4d |
| **Arbeitswochen** | 0.9w | 0.2w | 0.7w |

### Was bedeutet das?

**Human-only:** So lange wuerde die Implementierung dauern, wenn ein Entwickler komplett manuell arbeitet (ohne KI-Unterstuetzung).

**Human + KI Agent:** So lange dauert es realistisch mit modernen KI-Werkzeugen (Claude Code, Cursor, GitHub Copilot, etc.). Der Entwickler bleibt verantwortlich fuer Architektur, Code-Review und Qualitaetssicherung.

---

## Schaetzung pro Story

| ID | Story | Komplexitaet | Human (h) | KI-Faktor | KI-Adjusted (h) | Ersparnis |
|----|-------|-------------|-----------|-----------|-----------------|-----------|
| LAYOUT-001 | Update Style Tokens | XS | 3h | high (0.20) | 0.6h | 2.4h |
| LAYOUT-002 | Fix Navigation Sidebar | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-003 | Fix Status Bar | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-004 | Fix Welcome Screen | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-005 | Fix Build Components | S | 6h | high (0.20) | 1.2h | 4.8h |
| LAYOUT-006 | Fix Config Modal | XS | 3h | high (0.20) | 0.6h | 2.4h |
| LAYOUT-007 | Fix Build Execution | S | 6h | medium (0.40) | 2.4h | 3.6h |
| LAYOUT-008 | Fix Border Overflow | S | 6h | medium (0.40) | 2.4h | 3.6h |
| **TOTAL** | | | **42h** | | **10.8h** | **31.2h** |

*Note: System stories (997, 998, 999) are automated and not counted in effort.*

---

## KI-Beschleunigung nach Kategorie

| Kategorie | Stories | Human (h) | KI-Adjusted (h) | Reduktion |
|-----------|---------|-----------|-----------------|-----------|
| **High** (80% schneller) | 6 | 30h | 6.0h | -80% |
| **Medium** (60% schneller) | 2 | 12h | 4.8h | -60% |
| **Low** (30% schneller) | 0 | 0h | 0h | - |
| **None** (keine Beschleunigung) | 0 | 0h | 0h | - |

### Erklaerung der Kategorien

- **High (Faktor 0.20):** Die meisten Stories sind reine Styling/Layout-Aenderungen - Lip Gloss Styles, Farb-Tokens, Rendering-Anpassungen. KI kann diese sehr effizient umsetzen da die Patterns klar definiert sind.
- **Medium (Faktor 0.40):** Build Execution und Border Overflow erfordern etwas mehr Analyse der Width-Berechnungen und Layout-Logik.

---

## Annahmen & Hinweise

- Schaetzungen basieren auf der Komplexitaetsbewertung des Architects
- KI-Faktoren setzen aktive Nutzung von AI-Tools voraus (Claude Code, Cursor, etc.)
- Qualitaetssicherung und Code-Review bleiben unveraendert wichtig
- Unvorhergesehene Probleme koennen Aufwand erhoehen (+20-30% Puffer empfohlen)
- Golden File Tests muessen nach den Aenderungen aktualisiert werden (nicht in Schaetzung enthalten)

---

## Empfehlung

**Geplanter Aufwand:** 10.8h (1.4d / 0.3w)
**Mit Puffer (+25%):** 13.5h (1.7d / 0.3w)

---

*Erstellt mit Agent OS /create-spec v3.0*
