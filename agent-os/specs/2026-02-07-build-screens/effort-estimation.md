# Aufwandsschaetzung: Build Screens (Sprint 2.2)

**Erstellt:** 2026-02-07
**Spec:** Build Screens (Sprint 2.2)
**Anzahl Stories:** 5 (+ 3 System Stories)

---

## Zusammenfassung

| Metrik | Human-only | Human + KI Agent | Ersparnis |
|--------|------------|------------------|-----------|
| **Stunden** | 36h | 12h | 24h (67%) |
| **Arbeitstage** | 4.5d | 1.5d | 3d |
| **Arbeitswochen** | ~1w | ~0.4w | ~0.6w |

### Was bedeutet das?

**Human-only:** So lange wuerde die Implementierung dauern, wenn ein Entwickler komplett manuell arbeitet (ohne KI-Unterstuetzung).

**Human + KI Agent:** So lange dauert es realistisch mit modernen KI-Werkzeugen (Claude Code, Cursor, GitHub Copilot, etc.). Der Entwickler bleibt verantwortlich fuer Architektur, Code-Review und Qualitaetssicherung.

---

## Schaetzung pro Story

| ID | Story | Komplexitaet | Human (h) | KI-Faktor | KI-Adjusted (h) | Ersparnis |
|----|-------|-------------|-----------|-----------|-----------------|-----------|
| BUILD-001 | Domain Model & Mock Data | S | 6h | high (0.20) | 1.2h | 4.8h |
| BUILD-002 | Build Component Selection | S | 6h | high (0.20) | 1.2h | 4.8h |
| BUILD-003 | Build Configuration Modal | S | 6h | medium (0.40) | 2.4h | 3.6h |
| BUILD-004 | Build Execution View | M | 12h | medium (0.40) | 4.8h | 7.2h |
| BUILD-005 | App Integration & Tests | S | 6h | high (0.20) | 1.2h | 4.8h |
| **TOTAL** | | | **36h** | | **10.8h** | **25.2h** |

---

## KI-Beschleunigung nach Kategorie

| Kategorie | Stories | Human (h) | KI-Adjusted (h) | Reduktion |
|-----------|---------|-----------|-----------------|-----------|
| **High** (80% schneller) | 3 (001, 002, 005) | 18h | 3.6h | -80% |
| **Medium** (60% schneller) | 2 (003, 004) | 18h | 7.2h | -60% |
| **Low** (30% schneller) | 0 | 0h | 0h | - |
| **None** (keine Beschleunigung) | 0 | 0h | 0h | - |

### Erklaerung der Kategorien

- **High (Faktor 0.20):** BUILD-001 (CRUD-artige Domain-Typen), BUILD-002 (Standard TUI-List-Pattern), BUILD-005 (Wiring/Integration) - KI kann 5x schneller helfen
- **Medium (Faktor 0.40):** BUILD-003 (komplexere Modal-Komposition), BUILD-004 (Simulation-Logik, Custom-Table) - KI hilft 2.5x schneller
- **Low/None:** Keine Stories in diesen Kategorien (alle sind gut automatisierbar)

---

## Annahmen & Hinweise

- Schaetzungen basieren auf der Komplexitaetsbewertung des Architects
- KI-Faktoren setzen aktive Nutzung von AI-Tools voraus (Claude Code, Cursor, etc.)
- Qualitaetssicherung und Code-Review bleiben unveraendert wichtig
- Unvorhergesehene Probleme koennen Aufwand erhoehen (+20-30% Puffer empfohlen)
- BUILD-004 ist als M (Medium) am Limit - koennte mehr Zeit benoetigen

---

## Empfehlung

**Geplanter Aufwand:** 10.8h (~1.4d)
**Mit Puffer (+25%):** 13.5h (~1.7d)

---

*Erstellt mit Agent OS /create-spec v3.0*
