# Implementation Plan: Layout & Navigation Components

**Created:** 2026-02-06
**Status:** Draft
**Based on:** requirements-clarification.md (approved)

---

## Executive Summary

Implementierung von 8 Layout & Navigation TUI-Komponenten für den RFZ Developer CLI, die das Component Library Foundation vervollständigen. Diese Komponenten nutzen das Bubbles-Wrapper-Pattern wo sinnvoll (Table) und folgen den etablierten Patterns aus Sprint 1.1/1.2.

---

## Architektur-Entscheidungen

### Pattern 1: Bubbles Wrapper (TuiTable)
- Nutzt `bubbles/table` als Basis
- RFZ-Styling über Lipgloss angewandt
- Eigene API mit vereinfachter Nutzung
- **Vorteil:** Bewährte Scroll-/Selection-Logik, weniger Code

### Pattern 2: Pure Lipgloss (Navigation, Modal, Tabs, StatusBar, Tree, KeyHints)
- Funktionale Render-Funktionen wie bestehende Komponenten
- Kein State-Management in der Komponente selbst
- State wird vom Parent (Bubble Tea Model) verwaltet
- **Vorteil:** Konsistent mit TuiBox, TuiList, etc.

### Pattern 3: Composite Component (TuiNavigation)
- Kombiniert TuiNavItem-Elemente
- Container mit optionalem Header/Footer
- **Vorteil:** Wiederverwendbare Sub-Komponenten

---

## Komponenten-Übersicht

### Neue Komponenten (8 total)

| Komponente | Datei | Pattern | Komplexität |
|------------|-------|---------|-------------|
| TuiNavigation | navigation.go | Pure Lipgloss + Composite | S |
| TuiNavItem | navigation.go | Pure Lipgloss | XS |
| TuiModal | modal.go | Pure Lipgloss | M |
| TuiTable | table.go | Bubbles Wrapper | S |
| TuiTabs | tabs.go | Pure Lipgloss | S |
| TuiStatusBar | statusbar.go | Pure Lipgloss | S |
| TuiTree | tree.go | Pure Lipgloss | M |
| TuiKeyHints | keyhints.go | Pure Lipgloss | XS |

### Komponenten-Verbindungen

| Source | Target | Verbindungsart | Zuständige Story |
|--------|--------|----------------|------------------|
| TuiNavigation | TuiNavItem | Composition (enthält N Items) | Story-001 |
| TuiNavigation | TuiKeyHints | Optional Footer | Story-007 Integration |
| TuiModal | TuiButton | Footer Actions | Story-002 |
| TuiStatusBar | TuiKeyHints | Enthält KeyHints | Story-005 |

### Styles-Erweiterungen benötigt

- Navigation: Bereits in styles.go vorhanden (StyleNavItem, StyleNavItemFocused, StyleNavItemActive)
- Modal: Neue Styles für Backdrop, Title Bar
- Tabs: Neue Styles für aktive/inaktive Tabs
- StatusBar: Footer-Styles bereits vorhanden, erweitern
- Tree: Neue Styles für Nodes, Expand/Collapse Icons
- KeyHints: Nutzt StyleKeyboard, erweitert

---

## Umsetzungsphasen

### Phase 1: Basis-Komponenten (Stories 1-3)
**Ziel:** Navigation, Modal, und KeyHints - die Grundbausteine für Screen-Navigation

1. **TuiNavigation + TuiNavItem** (Story-001)
   - TuiNavItem als einzelne Render-Funktion
   - TuiNavigation als Container mit Items-Array
   - Vertikale Ausrichtung, Cursor-Highlight
   - Optional: Header, Footer (für KeyHints)

2. **TuiModal** (Story-002)
   - Double Border, zentriert
   - Title Bar, Content Area, Footer mit Buttons
   - Focus Trapping: Escape schließt, Tab cycled Buttons
   - Backdrop (dimmed)

3. **TuiKeyHints** (Story-003)
   - Kompakte Darstellung: "Enter Select • Esc Cancel"
   - Adaptive Breite
   - Separator zwischen Items (•)

### Phase 2: Daten-Komponenten (Stories 4-5)
**Ziel:** Table und Tree für Datenvisualisierung

4. **TuiTable** (Story-004)
   - Wrapper um bubbles/table
   - RFZ-Styles angewandt
   - Selectable Rows, Column Headers

5. **TuiTree** (Story-005)
   - Einfache Expand/Collapse
   - Hierarchische Darstellung mit Indentation
   - Icons: ▶/▼ oder +/-
   - Keyboard: j/k Navigation

### Phase 3: UI-Komponenten (Stories 6-7)
**Ziel:** Tabs und StatusBar für Screen-Layout

6. **TuiTabs** (Story-006)
   - Horizontale Tab-Leiste
   - Numerische Shortcuts (1-9)
   - Aktiver Tab visuell hervorgehoben

7. **TuiStatusBar** (Story-007)
   - Full-width Bottom Bar
   - Left/Center/Right Sections
   - Integration mit TuiKeyHints

### Phase 4: Demo & Tests (Stories 8-9)
**Ziel:** Validierung und Dokumentation

8. **Layout Navigation Demo** (Story-008)
   - Separates Demo-Programm
   - Alle 8 Komponenten showcased
   - Interaktive Navigation möglich

9. **Visual Regression Tests** (Story-009)
   - teatest Golden Files
   - Alle States pro Komponente
   - Integration mit CI

---

## Abhängigkeiten

### Interne Abhängigkeiten
```
Story-001 (Navigation)  → Story-007 (StatusBar) integriert KeyHints
Story-003 (KeyHints)    → Story-005 (StatusBar) verwendet KeyHints
Story-002 (Modal)       → keine (nutzt TuiButton aus Sprint 1.1)
Story-004 (Table)       → keine
Story-005 (Tree)        → keine
Story-006 (Tabs)        → keine
Story-007 (StatusBar)   → Story-003 (KeyHints)
Story-008 (Demo)        → Story-001 bis Story-007
Story-009 (Tests)       → Story-001 bis Story-007
```

### Execution Order
1. **Parallel möglich:** Stories 1-6 (alle Basis-Komponenten)
2. **Sequentiell danach:** Story-007 (benötigt KeyHints)
3. **Parallel danach:** Stories 8-9 (benötigen alle Komponenten)

---

## Risiken & Mitigationen

| Risiko | Wahrscheinlichkeit | Auswirkung | Mitigation |
|--------|-------------------|------------|------------|
| Modal Backdrop Performance | Niedrig | Mittel | Backdrop nur bei Änderung rendern |
| bubbles/table API Inkompatibilität | Niedrig | Mittel | Wrapper abstrahiert bubbles API vollständig |
| Tree Rekursion bei tiefen Hierarchien | Niedrig | Hoch | Max-Depth Limit (10 Levels) |
| Focus Trapping in Modal komplex | Mittel | Mittel | Einfaches Tab-Cycling, kein Full Focus Management |

---

## Self-Review Ergebnisse

### Vollständigkeit
- [x] Alle 8 Komponenten aus Requirements abgedeckt
- [x] Demo-Programm geplant (separate Demo, nicht Gallery-Erweiterung)
- [x] Visual Tests geplant

### Konsistenz
- [x] Pattern-Entscheidungen konsistent mit Sprint 1.1/1.2
- [x] Bubbles-Wrapper nur wo sinnvoll (Table)
- [x] Styles-Erweiterungen dokumentiert

### Risiken
- [x] Modal Focus-Trapping als mittleres Risiko identifiziert
- [x] Tree-Rekursion mit Max-Depth mitigiert

### Alternativen
- **Alternative für Modal:** Eigenes Focus-Management vs. Parent-gesteuert
  - **Gewählt:** Eigenes Focus-Management (Escape + Tab) - bessere UX

- **Alternative für Tree:** bubbles/list Wrapper vs. Custom
  - **Gewählt:** Custom - bubbles/list ist für flache Listen, nicht Hierarchien

### Komponenten-Verbindungen
- [x] TuiNavigation → TuiNavItem Verbindung in Story-001
- [x] TuiNavigation → TuiKeyHints Verbindung in Story-007
- [x] TuiModal → TuiButton Verbindung in Story-002
- [x] TuiStatusBar → TuiKeyHints Verbindung in Story-005

---

## Minimalinvasiv-Optimierungen

### Wiederverwendung bestehender Patterns

1. **styles.go bereits enthält:**
   - Navigation Styles (StyleNavItem, StyleNavItemFocused, StyleNavItemActive)
   - Footer Styles (StyleFooter)
   - Button Styles (für Modal Footer)
   - Box Styles (für Modal Container)

2. **helpers.go Truncate-Funktion:**
   - Wird für lange Labels in NavItem, Tabs, Table wiederverwendet

3. **TuiButton aus Sprint 1.1:**
   - Wird für Modal Footer Actions wiederverwendet

4. **TuiDivider aus Sprint 1.1:**
   - Wird für Modal Title/Content Trennung wiederverwendet

### Änderungsumfang minimiert

| Was | Änderung | Statt |
|-----|----------|-------|
| Neue Styles | Nur 4 neue Style-Blöcke | 8 separate Style-Dateien |
| Modal | Nutzt TuiBox als Basis | Komplett eigene Box-Logik |
| StatusBar | Erweitert StyleFooter | Neuer Footer-Stack |
| Tree | Eigene Render-Logik | bubbles/tree (existiert nicht) |

### Feature-Preservation Checkliste
- [x] Alle Requirements aus Clarification sind abgedeckt
- [x] Kein Feature wurde geopfert
- [x] Alle Akzeptanzkriterien bleiben erfüllbar

---

## Dateien zu erstellen

### Neue Dateien (7 total)
```
internal/ui/components/
├── navigation.go          # TuiNavigation + TuiNavItem
├── navigation_test.go     # Tests
├── modal.go               # TuiModal
├── modal_test.go          # Tests
├── table.go               # TuiTable (bubbles wrapper)
├── table_test.go          # Tests
├── tabs.go                # TuiTabs
├── tabs_test.go           # Tests
├── statusbar.go           # TuiStatusBar
├── statusbar_test.go      # Tests
├── tree.go                # TuiTree
├── tree_test.go           # Tests
├── keyhints.go            # TuiKeyHints
├── keyhints_test.go       # Tests
└── demo/
    └── layout_gallery.go  # Separate Demo
```

### Modifizierte Dateien (1 total)
```
internal/ui/components/
└── styles.go              # Neue Style-Definitionen hinzufügen
```

---

**Status:** Bereit zur Genehmigung

---

*Erstellt mit Agent OS /create-spec v3.0*
