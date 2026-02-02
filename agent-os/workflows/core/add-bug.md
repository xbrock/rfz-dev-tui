---
description: Add bug to backlog with hypothesis-driven root-cause analysis
globs:
alwaysApply: false
version: 2.4
encoding: UTF-8
---

# Add Bug Workflow

## Overview

Add a bug to the backlog with structured root-cause analysis. Uses hypothesis-driven debugging to identify the actual cause before creating the fix story.

**Key Difference to /add-todo:**
- Includes systematic Root-Cause-Analyse (RCA)
- 3 Hypothesen mit Wahrscheinlichkeiten
- Zuständiger Agent prüft jede Hypothese
- Dokumentierter Analyseprozess
- **NEU: User Hypothesis Dialog** - Benutzer-Wissen VOR der RCA abfragen

**v2.4 Changes:**
- **NEW: User Hypothesis Dialog (Step 2.5)** - Interaktiver Dialog VOR der RCA
  - Benutzer kann eigene Vermutungen teilen
  - Bereits untersuchte Bereiche dokumentieren
  - Ausgeschlossene Ursachen markieren
  - Gemeinsame Diskussion möglich
- **ENHANCED: RCA berücksichtigt User-Input** - User-Hypothesen werden priorisiert
- **NEW: Quelle-Spalte in Hypothesen-Tabelle** - Zeigt ob Hypothese von User oder Agent

**v2.3 Changes:**
- Gherkin-Style Bug-Fix Stories - Akzeptanzkriterien als Given-When-Then Szenarien
- Bug-spezifische Szenarien - Korrektes Verhalten, Regression-Schutz, Edge-Cases
- Trennung zwischen fachlichen Gherkin-Szenarien und technischer Verifikation

<pre_flight_check>
  EXECUTE: agent-os/workflows/meta/pre-flight.md
</pre_flight_check>

<process_flow>

<step number="1" name="backlog_setup">

### Step 1: Backlog Setup

<mandatory_actions>
  1. CHECK: Does agent-os/backlog/ directory exist?
     ```bash
     ls -la agent-os/backlog/ 2>/dev/null
     ```

  2. IF NOT exists:
     CREATE: agent-os/backlog/ directory
     CREATE: agent-os/backlog/story-index.md (from template)

     <template_lookup>
       PATH: backlog-story-index-template.md

       LOOKUP STRATEGY (MUST TRY BOTH):
         1. READ: agent-os/templates/docs/backlog-story-index-template.md
         2. IF file not found OR read error:
            READ: ~/.agent-os/templates/docs/backlog-story-index-template.md
         3. IF both fail: Error - run setup-devteam-global.sh

       ⚠️ WICHTIG: Bei "Error reading file" IMMER den Fallback-Pfad versuchen!
     </template_lookup>

  3. USE: date-checker to get current date (YYYY-MM-DD)

  4. DETERMINE: Next bug index for today
     COUNT: Existing bugs with today's date prefix
     ```bash
     ls agent-os/backlog/ | grep "^bug-$(date +%Y-%m-%d)" | wc -l
     ```
     NEXT_INDEX = count + 1 (formatted as 3 digits: 001, 002, etc.)

  5. GENERATE: Bug ID = YYYY-MM-DD-[INDEX]
     Example: 2025-01-15-001, 2025-01-15-002
</mandatory_actions>

</step>

<step number="2" name="bug_description">

### Step 2: Bug Description (PO Phase)

Gather structured bug information from user.

<mandatory_actions>
  1. IF user provided bug description in command:
     EXTRACT: Bug description from input

  2. ASK structured questions:

     **Symptom:**
     - Was genau passiert? (Fehlermeldung, falsches Verhalten, etc.)

     **Reproduktion:**
     - Wie kann der Bug reproduziert werden?
     - Schritt-für-Schritt Anleitung

     **Expected vs. Actual:**
     - Was sollte passieren? (Expected)
     - Was passiert stattdessen? (Actual)

     **Kontext:**
     - Welche Komponente/Seite ist betroffen?
     - Wann tritt es auf? (immer, manchmal, nach bestimmter Aktion)
     - Gibt es Fehlermeldungen in Console/Logs?

  3. DETERMINE: Bug-Typ
     - Frontend (UI, JavaScript, Styling)
     - Backend (API, Logik, Database)
     - DevOps (Build, Deployment, Infrastructure)
     - Integration (Zusammenspiel mehrerer Komponenten)

  4. DETERMINE: Severity
     - Critical: System unbenutzbar
     - High: Wichtige Funktion kaputt
     - Medium: Funktion eingeschränkt
     - Low: Kosmetisch oder Workaround vorhanden
</mandatory_actions>

</step>

<step number="2.5" name="user_hypothesis_dialog">

### Step 2.5: User Hypothesis Dialog

⚠️ **QUALITÄTSBOOSTER:** Benutzer-Wissen VOR der automatischen RCA abfragen.
Der Benutzer kennt oft das System besser und hat möglicherweise bereits untersucht.

<mandatory_actions>
  1. ASK user via AskUserQuestion:

     ```
     Question: "Haben Sie bereits eigene Vermutungen zur Ursache des Bugs?"

     Options:
     1. Ja, ich habe Vermutungen
        → Ich habe eine Idee, wo der Fehler liegen könnte

     2. Ich habe bereits selbst gesucht
        → Ich habe schon untersucht und kann Erkenntnisse teilen

     3. Nein, keine Ahnung
        → Ich habe keine Vermutung, Agent soll analysieren

     4. Ich möchte diskutieren
        → Lass uns gemeinsam überlegen
     ```

  2. BASED ON user choice:

     **IF "Ja, ich habe Vermutungen":**

     ASK follow-up questions (iterativ):

     a) "In welchem Bereich vermuten Sie den Fehler?"
        - Frontend (UI, Komponenten, State)
        - Backend (API, Services, Logik)
        - Datenbank (Queries, Schema)
        - Integration (Zusammenspiel)
        - Konfiguration (Environment, Settings)

     b) "Welche Dateien oder Komponenten haben Sie im Verdacht?"
        - Konkrete Dateinamen/Pfade
        - Komponenten-Namen
        - Funktionen/Methoden

     c) "Was könnte die Ursache sein?"
        - Freie Beschreibung der Vermutung
        - Warum vermuten Sie das?

     DOCUMENT:
     ```
     User-Hypothese:
     - Vermuteter Bereich: [BEREICH]
     - Verdächtige Dateien: [DATEIEN]
     - Vermutete Ursache: [BESCHREIBUNG]
     - Begründung: [WARUM]
     ```

     **IF "Ich habe bereits selbst gesucht":**

     ASK follow-up questions:

     a) "Welche Stellen haben Sie bereits untersucht?"
        - Dateien die geprüft wurden
        - Logs die analysiert wurden
        - Tests die durchgeführt wurden

     b) "Was haben Sie dabei festgestellt?"
        - Auffälligkeiten
        - Fehlermeldungen
        - Unerwartetes Verhalten

     c) "Was können wir definitiv ausschließen?"
        - Bereiche die NICHT die Ursache sind
        - Komponenten die korrekt funktionieren

     DOCUMENT:
     ```
     User-Recherche:
     - Bereits untersucht: [STELLEN]
     - Erkenntnisse: [FESTSTELLUNGEN]
     - Ausgeschlossen: [BEREICHE]
     ```

     **IF "Nein, keine Ahnung":**

     ACKNOWLEDGE: "Kein Problem, der Agent wird systematisch analysieren."
     PROCEED: Direkt zu Step 3 ohne User-Hypothesen

     **IF "Ich möchte diskutieren":**

     ENGAGE in dialog:

     a) "Was wissen wir über das Problem?"
        - Zusammenfassung aus Step 2

     b) "Wo könnte man anfangen zu suchen?"
        - Gemeinsam Ideen sammeln

     c) "Gibt es ähnliche Probleme in der Vergangenheit?"
        - Bekannte Patterns
        - Wiederkehrende Issues

     d) "Was würden Sie als erstes prüfen?"
        - Priorisierung der Untersuchung

     DOCUMENT: Alle Diskussions-Ergebnisse

  3. COMPILE User-Input für Step 3:

     <user_input_summary>
       ## User-Input zur Bug-Analyse

       **Hat der User Vermutungen:** Ja/Nein

       **User-Hypothesen:**
       [Falls vorhanden - Vermutungen des Users]

       **Bereits untersucht:**
       [Falls vorhanden - was der User schon geprüft hat]

       **Ausgeschlossene Bereiche:**
       [Falls vorhanden - was definitiv NICHT die Ursache ist]

       **Diskussions-Erkenntnisse:**
       [Falls diskutiert - gemeinsame Überlegungen]

       **Priorisierte Untersuchungs-Bereiche:**
       [Falls vorhanden - wo zuerst suchen]
     </user_input_summary>

  4. PASS user_input_summary to Step 3
</mandatory_actions>

<instructions>
  ACTION: Benutzer-Wissen vor RCA abfragen
  FORMAT: AskUserQuestion mit Follow-up Dialog
  DOCUMENT: Alle User-Inputs strukturiert
  VALUE: Verbessert RCA-Qualität signifikant
  SKIP: Nur wenn User "Keine Ahnung" wählt
</instructions>

</step>

<step number="3" name="hypothesis_driven_rca">

### Step 3: Hypothesis-Driven Root-Cause-Analyse

⚠️ **KERNSTÜCK:** Systematische Fehleranalyse statt blindes Suchen.

<determine_agent>
  BASED ON bug_type (from Step 2):

  IF bug_type = "Frontend":
    AGENT = dev-team__frontend-developer-*
  ELSE IF bug_type = "Backend":
    AGENT = dev-team__backend-developer-*
  ELSE IF bug_type = "DevOps":
    AGENT = dev-team__devops-specialist
  ELSE IF bug_type = "Integration":
    AGENT = dev-team__architect

  FALLBACK: If specific agent not available, use dev-team__architect
</determine_agent>

<delegation>
  DELEGATE to [AGENT] via Task tool:

  PROMPT:
  "Führe eine Hypothesis-Driven Root-Cause-Analyse durch.

  **Bug-Symptom:**
  [Bug description from Step 2]

  **Reproduktionsschritte:**
  [Steps from Step 2]

  **Expected:** [Expected behavior]
  **Actual:** [Actual behavior]

  **Betroffene Komponente:** [Component]

  ---

  ## User-Input aus Vorgespräch (Step 2.5)

  ⚠️ **WICHTIG:** Der Benutzer hat folgende Informationen geteilt.
  Diese MÜSSEN in deine Hypothesen einfließen!

  [USER_INPUT_SUMMARY from Step 2.5]

  **Anweisungen zur Nutzung des User-Inputs:**

  - **User-Hypothesen:** Wenn der User eine Vermutung hat, mache diese zur
    Hypothese #1 oder #2 (hohe Priorität). Der User kennt das System!

  - **Bereits untersucht:** Bereiche die der User schon geprüft hat, können
    mit niedrigerer Priorität behandelt werden (aber nicht ausschließen).

  - **Ausgeschlossene Bereiche:** Diese kannst du als "unwahrscheinlich"
    markieren, aber prüfe sie trotzdem wenn andere Hypothesen scheitern.

  - **Verdächtige Dateien:** Beginne deine Analyse mit diesen Dateien!

  ---

  ## Deine Aufgabe: Root-Cause-Analyse

  ### Phase 1: Hypothesen aufstellen

  Basierend auf dem Symptom UND dem User-Input, stelle 3 wahrscheinliche Ursachen auf.
  Ordne jeder Hypothese eine Wahrscheinlichkeit zu (muss 100% ergeben).

  ⚠️ **User-Hypothesen priorisieren:** Wenn der User eine Vermutung geteilt hat,
  sollte diese als Hypothese #1 oder #2 erscheinen (es sei denn, sie ist
  offensichtlich falsch).

  FORMAT:
  | # | Hypothese | Wahrscheinlichkeit | Quelle | Prüfmethode |
  |---|-----------|-------------------|--------|-------------|
  | 1 | [Vermutung] | XX% | User/Agent | [Wie prüfen - konkret] |
  | 2 | [Vermutung] | XX% | User/Agent | [Wie prüfen - konkret] |
  | 3 | [Vermutung] | XX% | Agent | [Wie prüfen - konkret] |

  REGELN für Hypothesen:
  - **User-Input hat Vorrang** - User kennt das System oft besser
  - Beginne mit der wahrscheinlichsten Ursache (höchster %)
  - Hypothesen müssen prüfbar sein
  - Prüfmethode muss konkret sein (Datei lesen, Log prüfen, Code analysieren)
  - Keine vagen Vermutungen ('irgendwo im Code')
  - Markiere ob Hypothese vom User oder Agent stammt

  ### Phase 2: Hypothesen prüfen

  Prüfe jede Hypothese der Reihe nach (höchste Wahrscheinlichkeit zuerst).

  FORMAT für jede Prüfung:
  ```
  **Hypothese X prüfen:** [Hypothese]
  - Aktion: [Was du konkret geprüft hast]
  - Befund: [Was du gefunden hast - Code-Snippets, Logs, etc.]
  - Ergebnis: ❌ Ausgeschlossen / ✅ BESTÄTIGT
  - Begründung: [Warum ausgeschlossen oder bestätigt]
  ```

  REGELN für Prüfung:
  - Prüfe TATSÄCHLICH (lies Code, prüfe Logs, analysiere Daten)
  - Dokumentiere konkrete Befunde (Zeilen, Werte, Fehlermeldungen)
  - Stoppe wenn Root Cause gefunden (✅ BESTÄTIGT)
  - Wenn H1 ausgeschlossen → H2 prüfen → H3 prüfen

  ### Phase 3: Root Cause dokumentieren

  Wenn Root Cause gefunden:

  ```
  ## ROOT CAUSE

  **Ursache:** [Klare Beschreibung der Ursache]

  **Beweis:** [Konkreter Nachweis - Code, Logs, etc.]

  **Betroffene Dateien:**
  - [Datei 1]: [Was ist dort falsch]
  - [Datei 2]: [Was ist dort falsch]

  **Fix-Ansatz:** [Kurze Beschreibung wie zu beheben]
  ```

  ### Falls KEINE Hypothese bestätigt:

  Wenn alle 3 Hypothesen ausgeschlossen:
  1. Stelle 3 NEUE Hypothesen auf (andere Richtung)
  2. Wiederhole Prüfung
  3. Maximal 2 Runden, dann eskalieren an User

  ---

  WICHTIG:
  - Sei gründlich aber effizient
  - Dokumentiere jeden Schritt
  - Finde die ECHTE Ursache, nicht nur Symptome
  - Gib mir am Ende den vollständigen Analyse-Bericht zurück"

  WAIT for agent completion

  RECEIVE: Root-Cause-Analyse Bericht
</delegation>

</step>

<step number="3.5" name="fix_impact_layer_analysis">

### Step 3.5: Fix-Impact Layer Analysis (NEU)

⚠️ **PFLICHT:** Basierend auf RCA analysieren, welche Layer vom Fix betroffen sind.

<mandatory_actions>
  1. EXTRACT from RCA (Step 3):
     - Root Cause (confirmed hypothesis)
     - Betroffene Dateien (from analysis)
     - Fix-Ansatz (proposed fix)

  2. ANALYZE fix impact across layers:
     ```
     Fix-Impact Layer Checklist:
     - [ ] Frontend (UI behavior, components, state)
     - [ ] Backend (API response, services, logic)
     - [ ] Database (data integrity, queries)
     - [ ] Integration (connections between layers)
     - [ ] Tests (affected test files)
     ```

  3. FOR EACH potentially affected layer:
     ASSESS:
     - Direct impact: Layer where bug originates
     - Indirect impact: Layers that depend on the fix
     - Test coverage: Tests that verify the fix

  4. IDENTIFY Integration Points:
     IF bug fix affects data flow between layers:
       DOCUMENT: Connection points that need verification
       Example: "Backend API response change → Frontend must handle new field"

  5. DETERMINE Fix Scope:
     - IF only 1 layer affected: "[Layer]-only fix"
     - IF 2+ layers affected: "Full-stack fix"
       ⚠️ WARNING: "Full-stack bug fix - ensure all layers are updated"

  6. GENERATE Fix-Impact Summary:
     ```
     Fix Type: [Backend-only / Frontend-only / Full-stack]
     Affected Layers:
       - [Layer 1]: [Direct/Indirect] - [Impact description]
       - [Layer 2]: [Direct/Indirect] - [Impact description]
     Critical Integration Points:
       - [Point 1]: [Source] → [Target] - [Needs verification]
     Required Tests:
       - [Test scope per layer]
     ```

  7. PASS Fix-Impact Summary to:
     - Step 4 (Bug Story File creation)
     - Step 5 (Architect Refinement)
</mandatory_actions>

<output>
  Fix-Impact Summary:
  - Fix Type (scope)
  - Affected Layers with direct/indirect impact
  - Critical Integration Points
  - Required test coverage per layer
</output>

</step>

<step number="4" name="create_bug_story">

### Step 4: Create Bug Story File

<mandatory_actions>
  1. GENERATE: File name
     FORMAT: bug-[YYYY-MM-DD]-[INDEX]-[slug].md
     Example: bug-2025-01-15-001-login-after-reset.md

  2. CREATE bug story file with RCA included:

     <bug_story_template>
       # 🐛 [BUG_TITLE]

       > Bug ID: [BUG_ID]
       > Created: [DATE]
       > Severity: [SEVERITY]
       > Status: Ready

       **Priority**: [PRIORITY]
       **Type**: Bug - [Frontend/Backend/DevOps]
       **Affected Component**: [COMPONENT]

       ---

       ## Bug Description

       ### Symptom
       [Bug symptom description]

       ### Reproduktion
       1. [Step 1]
       2. [Step 2]
       3. [Step 3]

       ### Expected vs. Actual
       - **Expected:** [What should happen]
       - **Actual:** [What happens instead]

       ---

       ## User-Input (aus Step 2.5)

       > Dokumentation des Benutzer-Wissens vor der RCA

       **Hat User Vermutungen geteilt:** [Ja/Nein]

       ### User-Hypothesen
       [Falls vorhanden - Vermutungen des Users]
       - Vermuteter Bereich: [BEREICH]
       - Verdächtige Dateien: [DATEIEN]
       - Vermutete Ursache: [BESCHREIBUNG]

       ### Bereits vom User untersucht
       [Falls vorhanden - was der User schon geprüft hat]

       ### Ausgeschlossene Bereiche
       [Falls vorhanden - was definitiv NICHT die Ursache ist]

       ---

       ## Root-Cause-Analyse

       ### Hypothesen (vor Analyse)

       | # | Hypothese | Wahrscheinlichkeit | Quelle | Prüfmethode |
       |---|-----------|-------------------|--------|-------------|
       | 1 | [H1] | XX% | User/Agent | [Method] |
       | 2 | [H2] | XX% | User/Agent | [Method] |
       | 3 | [H3] | XX% | Agent | [Method] |

       ### Prüfung

       **Hypothese 1 prüfen:** [H1]
       - Aktion: [What was checked]
       - Befund: [What was found]
       - Ergebnis: [❌/✅]
       - Begründung: [Why]

       [... weitere Hypothesen ...]

       ### Root Cause

       **Ursache:** [Root cause description]

       **Beweis:** [Evidence]

       **Betroffene Dateien:**
       - [File 1]
       - [File 2]

       ---

       ## Feature (Bug-Fix)

       ```gherkin
       Feature: [BUG_TITLE] beheben
         Als [USER_ROLE]
         möchte ich dass [BUG_DESCRIPTION] behoben wird,
         damit [BENEFIT/EXPECTED_BEHAVIOR].
       ```

       ---

       ## Akzeptanzkriterien (Gherkin-Szenarien)

       > **Bug-Fix Szenarien:** Beschreiben das KORREKTE Verhalten nach dem Fix

       ### Szenario 1: Korrektes Verhalten (was vorher fehlschlug)

       ```gherkin
       Scenario: [ORIGINAL_BUG_SCENARIO] funktioniert korrekt
         Given [AUSGANGSSITUATION die vorher zum Bug führte]
         When [AKTION die vorher den Bug auslöste]
         Then [KORREKTES_ERWARTETES_VERHALTEN]
         And [KEINE_FEHLERMELDUNG_ODER_FALSCHES_VERHALTEN]
       ```

       ### Szenario 2: Regression-Schutz

       ```gherkin
       Scenario: Verwandte Funktionalität bleibt intakt
         Given [SETUP für verwandte Funktion]
         When [VERWANDTE_AKTION]
         Then [ERWARTETES_VERHALTEN bleibt unverändert]
       ```

       ### Edge-Case nach Fix

       ```gherkin
       Scenario: Edge-Case wird korrekt behandelt
         Given [EDGE_CASE_SITUATION]
         When [EDGE_CASE_AKTION]
         Then [KORREKTE_EDGE_CASE_BEHANDLUNG]
       ```

       **Beispiel für Bug "Login nach Passwort-Reset fehlschlägt":**
       ```gherkin
       Scenario: Login nach Passwort-Reset funktioniert
         Given ich habe mein Passwort auf "NeuesPasswort123" zurückgesetzt
         And ich habe die Bestätigungs-Email erhalten
         When ich mich mit meiner Email und "NeuesPasswort123" anmelde
         Then bin ich erfolgreich eingeloggt
         And ich sehe mein Dashboard

       Scenario: Normaler Login bleibt funktionsfähig
         Given ich bin ein Benutzer ohne Passwort-Reset
         When ich mich mit meinen ursprünglichen Zugangsdaten anmelde
         Then bin ich erfolgreich eingeloggt

       Scenario: Falsches neues Passwort wird abgelehnt
         Given ich habe mein Passwort zurückgesetzt
         When ich mich mit dem alten Passwort anmelde
         Then sehe ich "Ungültige Zugangsdaten"
       ```

       ---

       ## Technische Verifikation

       - [ ] BUG_FIXED: [Description of fix verification]
       - [ ] TEST_PASS: Regression test added and passing
       - [ ] LINT_PASS: No linting errors
       - [ ] MANUAL: Bug no longer reproducible with original steps

       ---

       ## Technisches Refinement (vom Architect)

       > **⚠️ WICHTIG:** Dieser Abschnitt wird vom Architect ausgefüllt

       ### DoR (Definition of Ready) - Vom Architect

       #### Bug-Analyse
       - [x] Bug reproduzierbar
       - [x] Root Cause identifiziert
       - [x] Betroffene Dateien bekannt

       #### Technische Vorbereitung
       - [ ] Fix-Ansatz definiert (WAS/WIE/WO)
       - [ ] Abhängigkeiten identifiziert
       - [ ] Risiken bewertet

       **Bug ist READY wenn alle Checkboxen angehakt sind.**

       ---

       ### DoD (Definition of Done) - Vom Architect

       - [ ] Bug behoben gemäß Root Cause
       - [ ] Regression Test hinzugefügt
       - [ ] Keine neuen Bugs eingeführt
       - [ ] Code Review durchgeführt
       - [ ] Original Reproduktionsschritte führen nicht mehr zum Bug

       **Bug ist DONE wenn alle Checkboxen angehakt sind.**

       ---

       ### Betroffene Layer & Komponenten (Fix-Impact)

       > **PFLICHT:** Basierend auf Fix-Impact Analysis (Step 3.5)

       **Fix Type:** [Backend-only / Frontend-only / Full-stack]

       **Betroffene Komponenten:**

       | Layer | Komponenten | Impact | Änderung |
       |-------|-------------|--------|----------|
       | [Layer] | [components] | Direct/Indirect | [Fix description] |

       **Kritische Integration Points:**
       - [Point]: [Source] → [Target] - [Verification needed]

       ---

       ### Technical Details

       **WAS:** [What needs to be fixed - based on Root Cause]

       **WIE (Architektur-Guidance ONLY):**
       - [Fix approach based on RCA]
       - [Constraints to respect]

       **WO:** [Files to modify - MUST cover ALL layers from Fix-Impact Analysis!]

       **WER:** [Agent based on bug type]

       **Abhängigkeiten:** None

       **Geschätzte Komplexität:** [XS/S/M based on RCA]

       ---

       ### Completion Check

       ```bash
       # Verify bug is fixed
       [VERIFY_COMMAND based on bug type]
       ```

       **Bug ist DONE wenn:**
       1. Original Reproduktionsschritte funktionieren korrekt
       2. Regression Test besteht
       3. Keine verwandten Fehler auftreten
     </bug_story_template>

  3. FILL in all fields from:
     - Step 2 (Bug Description)
     - Step 3 (RCA - vollständig übernehmen)

  4. LEAVE Architect sections partially empty (Step 5 fills them)

  5. WRITE: Bug file to agent-os/backlog/
</mandatory_actions>

</step>

<step number="5" name="architect_refinement">

### Step 5: Architect Phase - Technical Refinement (v3.0)

Main agent does technical refinement guided by architect-refinement skill.

<refinement_process>
  LOAD skill: .claude/skills/architect-refinement/SKILL.md
  (This skill provides guidance for technical refinement)

  **Bug Context:**
  - Bug File: agent-os/backlog/bug-[YYYY-MM-DD]-[INDEX]-[slug].md
  - Fix-Impact Summary (from Step 3.5): [FIX_IMPACT_SUMMARY]
  - Root Cause: Already identified in bug story
  - Tech Stack: agent-os/product/tech-stack.md
  - Architecture: Try both locations:
    1. agent-os/product/architecture-decision.md
    2. agent-os/product/architecture/platform-architecture.md
  - DoR/DoD: agent-os/team/dor.md and dod.md (if exist)

  **Tasks (guided by architect-refinement skill):**
  1. READ the bug story file (Root Cause section)
  2. REVIEW Fix-Impact Summary - ensure ALL layers addressed
  3. LOAD project quality definitions
  4. FILL technical sections:

     **Betroffene Layer & Komponenten (PFLICHT):**
     Based on Fix-Impact Summary:
     - Fix Type: [Backend-only / Frontend-only / Full-stack]
     - Betroffene Komponenten Table with Direct/Indirect impact
     - Kritische Integration Points (if Full-stack fix)

     **DoR vervollständigen:**
     - Apply relevant DoR criteria
     - Mark ALL checkboxes as [x] when complete

     **DoD:**
     - Define completion criteria (unchecked [ ])

     **Technical Details:**
     - WAS: What needs to be fixed
     - WIE: Fix approach (patterns, constraints)
     - WO: Files to modify (ALL layers!)
     - Domain: Optional domain area reference
     - Abhängigkeiten: None
     - Geschätzte Komplexität: XS/S/M

     **Completion Check:**
     - Add bash verify commands

  5. VALIDATE: Bug not too complex for backlog

  **IMPORTANT (v3.0):**
  - NO "WER" field (main agent implements directly)
  - Skills auto-load during implementation
  - Follow architect-refinement skill guidance
  - Keep lightweight
  - Mark ALL DoR checkboxes as [x] when ready
</refinement_process>

</step>

<step number="5.5" name="bug_size_validation">

### Step 5.5: Bug Size Validation

Validate that the bug fix complies with size guidelines for single-session execution.

<validation_process>
  READ: The bug file from agent-os/backlog/bug-[...].md

  <extract_metrics>
    ANALYZE: WO (Where) field
      COUNT: Number of file paths mentioned
      EXTRACT: File paths list

    ANALYZE: Geschätzte Komplexität field
      EXTRACT: Complexity rating (XS/S/M/L/XL)

    ANALYZE: Root Cause section
      ASSESS: Is this a localized bug or systemic issue?
      CHECK: Number of "Betroffene Dateien"

    ANALYZE: WAS (What) field
      ESTIMATE: Lines of code for fix
      HEURISTIC:
        - Simple fix (1-2 files) ~50-100 lines
        - Medium fix (3-4 files) ~150-250 lines
        - Complex fix (5+ files) ~300+ lines
  </extract_metrics>

  <check_thresholds>
    CHECK: Number of affected files
      IF files > 5:
        FLAG: Bug as "Too Large - Affects Too Many Files"
        SEVERITY: High

    CHECK: Complexity rating
      IF complexity in [L, XL]:
        FLAG: Bug as "Too Complex for /add-bug"
        SEVERITY: High
      ELSE IF complexity = M:
        FLAG: Bug as "Borderline Complexity"
        SEVERITY: Medium

    CHECK: Estimated LOC
      IF estimated_loc > 400:
        FLAG: Bug as "Too Large - Code Volume"
        SEVERITY: High
      ELSE IF estimated_loc > 250:
        FLAG: Bug as "Watch - Approaching Limit"
        SEVERITY: Low

    CHECK: Systemic issue detection
      IF Root Cause mentions "architectural", "design flaw", or "multiple components":
        FLAG: Bug as "Systemic Issue"
        SEVERITY: High
        SUGGEST: "Consider /create-spec for architectural fixes"

    CHECK: Full-Stack Fix Coverage (Enhanced)
      EXTRACT: "Betroffene Layer & Komponenten" section
      IF Fix Type = "Full-stack":
        CHECK: WO section covers ALL layers from "Betroffene Komponenten" table
        IF missing_layers detected:
          FLAG: Bug as "Incomplete Full-Stack Fix"
          SEVERITY: Critical
          LIST: "Missing file paths for layers: [missing_layers]"
          WARN: "Bug fix does not cover all affected layers - risk of partial fix!"
          SUGGEST: "Add ALL layer files to WO section OR split into multiple bugs"

        CHECK: Integration Points coverage
        IF Critical Integration Points defined:
          VERIFY: Each integration point has source AND target in WO
          IF missing_connections:
            FLAG: Bug as "Missing Integration Coverage"
            SEVERITY: High
            LIST: "Integration points not fully covered: [points]"
            WARN: "Fix may break integration between layers"
  </check_thresholds>
</validation_process>

<decision_tree>
  IF no flags raised OR only low severity:
    LOG: "✅ Bug passes size validation - appropriate for /add-bug"
    PROCEED: To Step 6 (Update Story Index)

  ELSE (bug flagged with Medium/High severity):
    GENERATE: Validation Report

    <validation_report_format>
      ⚠️ Bug Size Validation - Issues Detected

      **Bug:** 🐛 [Bug Title]
      **File:** [Bug file path]
      **Root Cause:** [Brief RC description]

      **Metrics:**
      - Affected Files: [count] (max recommended: 5) [✅/❌]
      - Complexity: [rating] (max recommended: S, tolerated: M) [✅/⚠️/❌]
      - Est. LOC for Fix: ~[count] (max recommended: 400) [✅/❌]
      - Systemic Issue: [Yes/No] [✅/❌]

      **Issue:** [Description of what exceeds guidelines]

      **Why this matters:**
      - /add-bug is designed for localized, contained bug fixes
      - Systemic issues need proper planning to avoid introducing new bugs
      - Complex fixes benefit from story splitting and integration testing

      **Recommendation:** Use /create-spec instead for:
      - Proper architectural analysis
      - Story splitting for safer implementation
      - Integration story to validate complete fix
      - Better dependency mapping
    </validation_report_format>

    PRESENT: Validation Report to user

    ASK user via AskUserQuestion:
    "This bug fix exceeds /add-bug size guidelines. How would you like to proceed?

    Options:
    1. Switch to /create-spec (Recommended)
       → Full specification with proper planning
       → Story splitting for safer implementation
       → Integration story for validation

    2. Edit bug to reduce scope
       → Focus on most critical part of the fix
       → Create follow-up bugs for remaining issues
       → Re-run validation after edits

    3. Proceed anyway
       → Accept higher context usage
       → Risk mid-execution context compaction
       → Continue with current bug fix"

    WAIT for user choice

    <user_choice_handling>
      IF choice = "Switch to /create-spec":
        INFORM: "Switching to /create-spec workflow.
                 The bug analysis and Root Cause will be preserved as context."

        PRESERVE: Root-Cause-Analyse for create-spec input

        INVOKE: /create-spec with bug description and RCA
        STOP: This workflow

      ELSE IF choice = "Edit bug to reduce scope":
        INFORM: "Please edit the bug file: agent-os/backlog/[bug-file].md"
        INFORM: "Reduce the scope by:
                 - Focus on the most critical affected file
                 - Create separate bugs for other affected areas
                 - Reduce WO section to essential files only"
        PAUSE: Wait for user to edit
        ASK: "Ready to re-validate? (yes/no)"
        IF yes:
          REPEAT: Step 5.5 (this validation step)
        ELSE:
          PROCEED: To Step 6 with warning flag

      ELSE IF choice = "Proceed anyway":
        WARN: "⚠️ Proceeding with oversized bug fix
               - Expect higher token costs
               - Mid-execution compaction possible
               - Consider splitting into multiple bugs next time"
        LOG: Validation bypassed by user
        PROCEED: To Step 6
    </user_choice_handling>
</decision_tree>

<instructions>
  ACTION: Validate bug against size guidelines
  CHECK: Affected files, complexity, estimated LOC, systemic issue detection
  THRESHOLD: Max 5 files, max M complexity (S preferred), max 400 LOC
  REPORT: Issues found with specific recommendations
  OFFER: Three options (switch to create-spec, edit scope, proceed)
  ENFORCE: Validation before adding to backlog
</instructions>

**Output:**
- Validation report (if issues found)
- User decision on how to proceed
- Bug either validated, edited, or escalated to /create-spec

</step>

<step number="6" name="update_story_index">

### Step 6: Update Backlog Story Index

<mandatory_actions>
  1. READ: agent-os/backlog/story-index.md

  2. ADD new bug to Story Summary table:
     | Bug ID | Title | Type | Priority | Dependencies | Status | Points |
     Note: Use 🐛 emoji prefix for bug entries

  3. UPDATE totals:
     - Total Stories: +1
     - Backlog Count: +1

  4. UPDATE: Last Updated date

  5. WRITE: Updated story-index.md
</mandatory_actions>

</step>

<step number="7" name="completion_summary">

### Step 7: Bug Added Confirmation

⚠️ **Note:** Only reached if bug passed size validation (Step 5.5)

<summary_template>
  ✅ Bug added to backlog with Root-Cause-Analyse!

  **Bug ID:** [YYYY-MM-DD-INDEX]
  **File:** agent-os/backlog/bug-[YYYY-MM-DD]-[INDEX]-[slug].md

  **Summary:**
  - Title: 🐛 [Bug Title]
  - Severity: [Critical/High/Medium/Low]
  - Root Cause: [Brief RC description]
  - Complexity: [XS/S/M]
  - Status: Ready

  **Root-Cause-Analyse:**
  - Hypothesen geprüft: [N]
  - Root Cause gefunden: ✅
  - Betroffene Dateien: [N]

  **Backlog Status:**
  - Total tasks: [N]
  - Bugs: [N]
  - Ready for execution: [N]

  **Next Steps:**
  1. Add more bugs: /add-bug "[description]"
  2. Add quick tasks: /add-todo "[description]"
  3. Execute backlog: /execute-tasks backlog
  4. View backlog: agent-os/backlog/story-index.md
</summary_template>

</step>

</process_flow>

## Final Checklist

<verify>
  - [ ] Backlog directory exists
  - [ ] Bug description gathered (symptom, repro, expected/actual)
  - [ ] Bug type determined (Frontend/Backend/DevOps)
  - [ ] **User Hypothesis Dialog completed (Step 2.5)**
  - [ ] **User-Input dokumentiert (falls vorhanden)**
  - [ ] Hypothesis-Driven RCA completed (mit User-Input)
  - [ ] Root Cause identified and documented
  - [ ] Bug story file created with correct naming
  - [ ] Technical refinement complete
  - [ ] All DoR checkboxes marked [x]
  - [ ] **Bug size validation passed (Step 5.5)**
  - [ ] Story-index.md updated
  - [ ] Ready for /execute-tasks backlog
</verify>

## When NOT to Use /add-bug

Suggest /create-spec instead when:
- Root Cause requires architectural changes
- Fix affects >5 files
- Multiple related bugs need coordinated fix
- Bug reveals larger design issue
- Estimated complexity > M
