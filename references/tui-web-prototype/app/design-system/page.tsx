"use client"

import React from "react"

import { useState } from "react"
import { TuiBox, TuiDivider } from "@/components/tui/tui-box"
import { TuiList, TuiNavItem } from "@/components/tui/tui-list"
import { TuiModal, TuiModalActions, TuiButton } from "@/components/tui/tui-modal"
import { TuiProgress, TuiSpinner, TuiStatus } from "@/components/tui/tui-progress"
import { cn } from "@/lib/utils"

// Component Section wrapper
function ComponentSection({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <div className="mb-8">
      <h2 className="text-lg text-tui-cyan mb-3 font-bold">{title}</h2>
      <div className="space-y-4">
        {children}
      </div>
    </div>
  )
}

// Component Card
function ComponentCard({ 
  name, 
  description, 
  children 
}: { 
  name: string
  description?: string
  children: React.ReactNode 
}) {
  return (
    <TuiBox title={name} borderStyle="single">
      <div className="p-3">
        {description && (
          <p className="text-xs text-muted-foreground mb-3">{description}</p>
        )}
        <div className="bg-background/50 p-2">
          {children}
        </div>
      </div>
    </TuiBox>
  )
}

export default function DesignSystemPage() {
  const [showModal, setShowModal] = useState(false)
  const [focusedNav, setFocusedNav] = useState(1)
  const [focusedList, setFocusedList] = useState(0)

  const sampleListItems = [
    { id: "1", label: "rfz-core", category: "Core" },
    { id: "2", label: "rfz-simulation", category: "Simu" },
    { id: "3", label: "rfz-standalone", category: "Stan" },
  ]

  return (
    <div className="fixed inset-0 bg-background text-foreground p-6 overflow-y-auto">
      {/* Header */}
      <div className="mb-8">
        <div className="h-1 bg-brand mb-2" />
        <h1 className="text-2xl font-bold text-foreground mb-1">RFZ-CLI Design System</h1>
        <p className="text-sm text-muted-foreground">TUI Component Reference for Go/Bubble Tea Implementation</p>
      </div>

      {/* Color Tokens */}
      <ComponentSection title="Color Tokens">
        <div className="flex flex-wrap gap-4">
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-background border border-border" />
            <span className="text-xs">background</span>
          </div>
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-card border border-border" />
            <span className="text-xs">card</span>
          </div>
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-secondary" />
            <span className="text-xs">secondary</span>
          </div>
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-border" />
            <span className="text-xs">border</span>
          </div>
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-tui-cyan" />
            <span className="text-xs">tui-cyan (focus)</span>
          </div>
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-tui-green" />
            <span className="text-xs">tui-green (success)</span>
          </div>
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-tui-yellow" />
            <span className="text-xs">tui-yellow (warning)</span>
          </div>
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-destructive" />
            <span className="text-xs">destructive (error)</span>
          </div>
          <div className="flex items-center gap-2">
            <div className="w-6 h-6 bg-brand" />
            <span className="text-xs">brand (DB Red)</span>
          </div>
        </div>
      </ComponentSection>

      {/* Layout Components */}
      <ComponentSection title="Layout Components">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          {/* TuiBox variants */}
          <ComponentCard name="TuiBox" description="Container with box-drawing borders. Supports: single, double, rounded, heavy">
            <div className="space-y-3">
              <TuiBox title="Single Border" borderStyle="single">
                <div className="p-1 text-xs">Content area</div>
              </TuiBox>
              <TuiBox title="Double Border" borderStyle="double">
                <div className="p-1 text-xs">Content area</div>
              </TuiBox>
              <TuiBox title="Heavy Border" borderStyle="heavy">
                <div className="p-1 text-xs">Content area</div>
              </TuiBox>
              <TuiBox title="Focused State" borderStyle="single" focused>
                <div className="p-1 text-xs">Focused = cyan border</div>
              </TuiBox>
            </div>
          </ComponentCard>

          {/* TuiDivider */}
          <ComponentCard name="TuiDivider" description="Horizontal separator line">
            <div className="space-y-2">
              <p className="text-xs text-muted-foreground">Single:</p>
              <TuiDivider style="single" />
              <p className="text-xs text-muted-foreground">Double:</p>
              <TuiDivider style="double" />
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Navigation Components */}
      <ComponentSection title="Navigation Components">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          {/* NavItem */}
          <ComponentCard name="TuiNavItem" description="Navigation menu item with index, shortcut, focus/active states">
            <div className="space-y-0">
              <TuiNavItem label="Build Components" index={0} isFocused={false} isActive={false} shortcut="1" />
              <TuiNavItem label="View Logs" index={1} isFocused={true} isActive={false} shortcut="2" />
              <TuiNavItem label="Discover" index={2} isFocused={false} isActive={true} shortcut="3" />
              <TuiNavItem label="Configuration" index={3} isFocused={false} isActive={false} shortcut="4" />
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              Row 2: focused | Row 3: active (current screen)
            </p>
          </ComponentCard>

          {/* Full Navigation Panel */}
          <ComponentCard name="Navigation Panel" description="Complete sidebar navigation with help footer">
            <TuiBox title="Navigation" borderStyle="single">
              <div className="py-1">
                {["Build", "Logs", "Discover", "Config"].map((label, i) => (
                  <TuiNavItem
                    key={label}
                    label={label}
                    index={i}
                    isFocused={i === focusedNav}
                    isActive={i === 2}
                    shortcut={String(i + 1)}
                    onClick={() => setFocusedNav(i)}
                  />
                ))}
              </div>
              <div className="mt-2 px-2 py-1 border-t border-border">
                <div className="text-xs text-muted-foreground">
                  <span className="text-tui-cyan">↑↓</span> navigate
                </div>
              </div>
            </TuiBox>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* List Components */}
      <ComponentSection title="List Components">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          {/* Simple List */}
          <ComponentCard name="TuiList" description="Navigable list with focus indicator">
            <TuiList
              items={sampleListItems}
              selectedIndex={focusedList}
              showCategories
            />
            <div className="mt-2 flex gap-2">
              <button 
                onClick={() => setFocusedList(Math.max(0, focusedList - 1))}
                className="text-xs text-tui-cyan hover:underline"
              >
                [k] Up
              </button>
              <button 
                onClick={() => setFocusedList(Math.min(2, focusedList + 1))}
                className="text-xs text-tui-cyan hover:underline"
              >
                [j] Down
              </button>
            </div>
          </ComponentCard>

          {/* Multi-select List */}
          <ComponentCard name="TuiList (Multi-Select)" description="List with checkbox selection">
            <TuiList
              items={sampleListItems}
              selectedIndex={1}
              selectedItems={new Set(["1", "3"])}
              multiSelect
              showCategories
            />
            <p className="text-xs text-muted-foreground mt-2">
              [x] = selected, [ ] = unselected, Space to toggle
            </p>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Status & Progress */}
      <ComponentSection title="Status & Progress Components">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          {/* Status indicators */}
          <ComponentCard name="TuiStatus" description="Build status indicators">
            <div className="space-y-2">
              <TuiStatus status="pending" label="Pending" />
              <TuiStatus status="running" label="Running..." />
              <TuiStatus status="success" label="Success" />
              <TuiStatus status="failed" label="Failed" />
              <TuiStatus status="error" label="Error" />
            </div>
          </ComponentCard>

          {/* Progress bar */}
          <ComponentCard name="TuiProgress" description="Progress bar with percentage">
            <div className="space-y-3">
              <TuiProgress value={0} width={20} label="0%" />
              <TuiProgress value={35} width={20} label="Building" />
              <TuiProgress value={75} width={20} />
              <TuiProgress value={100} width={20} label="Done" />
            </div>
          </ComponentCard>

          {/* Spinner */}
          <ComponentCard name="TuiSpinner" description="Animated loading spinner (Braille dots)">
            <div className="flex items-center gap-4">
              <TuiSpinner />
              <TuiSpinner label="Loading..." />
              <span className="text-xs text-muted-foreground">Frames: ⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏</span>
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Button Components */}
      <ComponentSection title="Button Components">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <ComponentCard name="TuiButton" description="Action buttons with keyboard shortcuts">
            <div className="flex flex-wrap gap-4">
              <TuiButton label="Cancel" shortcut="Esc" />
              <TuiButton label="Cancel" shortcut="Esc" focused />
              <TuiButton label="Confirm" shortcut="Enter" variant="primary" />
              <TuiButton label="Confirm" shortcut="Enter" variant="primary" focused />
              <TuiButton label="Delete" shortcut="d" variant="danger" />
              <TuiButton label="Delete" shortcut="d" variant="danger" focused />
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              Variants: default, primary (green), danger (red) | States: normal, focused
            </p>
          </ComponentCard>

          <ComponentCard name="Action Buttons (Inline)" description="Row of action buttons in footer style">
            <div className="p-2 border-t border-border">
              <div className="flex items-center gap-4">
                <TuiButton label="View Logs" shortcut="l" />
                <TuiButton label="Rebuild Failed" shortcut="r" variant="danger" />
                <TuiButton label="Back" shortcut="Esc" focused />
              </div>
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              Pattern: action bar at bottom of panels
            </p>
          </ComponentCard>

          <ComponentCard name="Toggle Buttons" description="On/Off state buttons">
            <div className="space-y-2">
              <div className="flex items-center gap-4">
                <span className="text-muted-foreground">Follow Mode:</span>
                <button className="px-2 py-0.5 bg-tui-green text-background text-sm">[ON]</button>
                <button className="px-2 py-0.5 bg-secondary text-muted-foreground text-sm">[OFF]</button>
              </div>
              <div className="flex items-center gap-4">
                <span className="text-muted-foreground">Errors Only:</span>
                <button className="px-2 py-0.5 bg-destructive/20 text-destructive text-sm">[x] Enabled</button>
                <button className="px-2 py-0.5 bg-secondary text-muted-foreground text-sm">[ ] Disabled</button>
              </div>
            </div>
          </ComponentCard>

          <ComponentCard name="Button Group" description="Grouped related actions">
            <div className="flex">
              <button className="px-3 py-1 bg-tui-cyan text-background text-sm border-r border-background/20">ALL</button>
              <button className="px-3 py-1 bg-secondary text-muted-foreground text-sm border-r border-border">INFO</button>
              <button className="px-3 py-1 bg-secondary text-muted-foreground text-sm border-r border-border">WARN</button>
              <button className="px-3 py-1 bg-secondary text-muted-foreground text-sm border-r border-border">ERROR</button>
              <button className="px-3 py-1 bg-secondary text-muted-foreground text-sm">DEBUG</button>
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              Pattern: filter/tab selection bar
            </p>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Select Components */}
      <ComponentSection title="Select Components">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <ComponentCard name="Single Select (Radio)" description="One option from a list">
            <div className="space-y-1">
              <div className="flex items-center gap-2 px-1">
                <span className="text-muted-foreground">( )</span>
                <span>clean</span>
              </div>
              <div className="flex items-center gap-2 px-1">
                <span className="text-muted-foreground">( )</span>
                <span>install</span>
              </div>
              <div className="flex items-center gap-2 px-1 bg-tui-cyan text-background">
                <span>{">"}</span>
                <span>(●)</span>
                <span>clean install</span>
              </div>
              <div className="flex items-center gap-2 px-1">
                <span className="text-muted-foreground">( )</span>
                <span>package</span>
              </div>
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              (●) = selected, ( ) = unselected, Enter/Space to select
            </p>
          </ComponentCard>

          <ComponentCard name="Multi Select (Checkbox)" description="Multiple options from a list">
            <div className="space-y-1">
              <div className="flex items-center gap-2 px-1">
                <span className="w-2 text-tui-cyan">{">"}</span>
                <span className="text-tui-green">[x]</span>
                <span>target_env_dev</span>
                <span className="ml-auto text-xs text-muted-foreground">[Environment]</span>
              </div>
              <div className="flex items-center gap-2 px-1 bg-tui-cyan text-background">
                <span className="w-2">{">"}</span>
                <span>[x]</span>
                <span>generate_local_config</span>
                <span className="ml-auto text-xs opacity-70">[Environment]</span>
              </div>
              <div className="flex items-center gap-2 px-1">
                <span className="w-2"> </span>
                <span className="text-muted-foreground">[ ]</span>
                <span>skip_integration_tests</span>
                <span className="ml-auto text-xs text-muted-foreground">[Testing]</span>
              </div>
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              [x] = selected, [ ] = unselected, Space to toggle
            </p>
          </ComponentCard>

          <ComponentCard name="Horizontal Select" description="Inline single-select options">
            <div className="space-y-3">
              <div>
                <p className="text-xs text-muted-foreground mb-1">Port Selection:</p>
                <div className="flex gap-4">
                  <span className="px-2 py-0.5 bg-tui-cyan text-background">{">"} (●) Port 11090</span>
                  <span className="text-muted-foreground">( ) Port 11091</span>
                </div>
              </div>
              <div>
                <p className="text-xs text-muted-foreground mb-1">Log Level:</p>
                <div className="flex gap-2">
                  <span className="px-2 py-0.5 bg-tui-cyan text-background">ALL</span>
                  <span className="px-2 py-0.5 text-muted-foreground">INFO</span>
                  <span className="px-2 py-0.5 text-muted-foreground">WARN</span>
                  <span className="px-2 py-0.5 text-muted-foreground">ERROR</span>
                </div>
              </div>
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              Pattern: h/l or arrow keys to navigate, Enter/Space to select
            </p>
          </ComponentCard>

          <ComponentCard name="Dropdown Select (Closed)" description="Collapsed select showing current value">
            <div className="space-y-2">
              <div className="flex items-center gap-2 bg-secondary px-2 py-1">
                <span>Maven Goal:</span>
                <span className="text-tui-cyan">clean install</span>
                <span className="ml-auto text-muted-foreground">[Enter to change]</span>
              </div>
              <div className="flex items-center gap-2 bg-secondary px-2 py-1">
                <span>Environment:</span>
                <span className="text-tui-cyan">development</span>
                <span className="ml-auto text-muted-foreground">[Enter to change]</span>
              </div>
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              Pattern: collapsed state showing current selection
            </p>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Table Components */}
      <ComponentSection title="Table Components">
        <div className="grid grid-cols-1 gap-4">
          <ComponentCard name="Data Table" description="Tabular data with selectable rows">
            <div className="overflow-x-auto">
              <div className="flex items-center gap-4 px-2 py-1 text-xs text-muted-foreground border-b border-border">
                <span className="w-4"></span>
                <span className="w-32">Component</span>
                <span className="w-20">Category</span>
                <span className="w-24">Status</span>
                <span className="flex-1">Branch</span>
              </div>
              <div className="flex items-center gap-4 px-2 py-0.5 text-sm">
                <span className="w-4 text-tui-cyan"> </span>
                <span className="w-32">rfz-core</span>
                <span className="w-20 text-xs text-tui-cyan">Core</span>
                <span className="w-24 text-tui-green">clean</span>
                <span className="flex-1 text-tui-green">develop</span>
              </div>
              <div className="flex items-center gap-4 px-2 py-0.5 text-sm bg-tui-cyan text-background">
                <span className="w-4">{">"}</span>
                <span className="w-32">rfz-simulation</span>
                <span className="w-20 text-xs opacity-80">Simu</span>
                <span className="w-24">dirty</span>
                <span className="flex-1">feature/xyz</span>
              </div>
              <div className="flex items-center gap-4 px-2 py-0.5 text-sm">
                <span className="w-4 text-tui-cyan"> </span>
                <span className="w-32">rfz-standalone</span>
                <span className="w-20 text-xs text-tui-yellow">Stan</span>
                <span className="w-24 text-tui-green">clean</span>
                <span className="flex-1 text-tui-green">develop</span>
              </div>
            </div>
          </ComponentCard>

          <ComponentCard name="Build Status Table" description="Component build progress">
            <div className="overflow-x-auto">
              <div className="flex items-center gap-4 px-2 py-1 text-xs text-muted-foreground border-b border-border">
                <span className="w-4"></span>
                <span className="w-32">Component</span>
                <span className="w-20">Status</span>
                <span className="w-16">Phase</span>
                <span className="flex-1">Progress</span>
              </div>
              <div className="flex items-center gap-4 px-2 py-0.5 text-sm">
                <span className="w-4 text-tui-green">✓</span>
                <span className="w-32">rfz-core</span>
                <span className="w-20 text-tui-green">success</span>
                <span className="w-16 text-xs">done</span>
                <span className="flex-1"><TuiProgress value={100} width={15} /></span>
              </div>
              <div className="flex items-center gap-4 px-2 py-0.5 text-sm bg-tui-cyan text-background">
                <span className="w-4"><TuiSpinner /></span>
                <span className="w-32">rfz-simulation</span>
                <span className="w-20">running</span>
                <span className="w-16 text-xs">compile</span>
                <span className="flex-1"><TuiProgress value={65} width={15} /></span>
              </div>
              <div className="flex items-center gap-4 px-2 py-0.5 text-sm">
                <span className="w-4 text-muted-foreground">○</span>
                <span className="w-32 text-muted-foreground">rfz-standalone</span>
                <span className="w-20 text-muted-foreground">pending</span>
                <span className="w-16 text-xs text-muted-foreground">-</span>
                <span className="flex-1"><TuiProgress value={0} width={15} /></span>
              </div>
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Modal */}
      <ComponentSection title="Modal Component">
        <ComponentCard name="TuiModal" description="Overlay dialog with double border">
          <div>
            <TuiButton 
              label="Open Modal" 
              shortcut="Enter"
              onClick={() => setShowModal(true)}
            />
            <p className="text-xs text-muted-foreground mt-2">
              Click to preview modal overlay
            </p>
          </div>
        </ComponentCard>
      </ComponentSection>

      {/* Form Controls */}
      <ComponentSection title="Form Controls">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          {/* Checkbox */}
          <ComponentCard name="Checkbox" description="Toggle selection">
            <div className="space-y-1">
              <div className="flex items-center gap-2">
                <span className="text-tui-green">[x]</span>
                <span>Selected option</span>
              </div>
              <div className="flex items-center gap-2">
                <span className="text-muted-foreground">[ ]</span>
                <span>Unselected option</span>
              </div>
              <div className="flex items-center gap-2 bg-tui-cyan text-background px-1">
                <span>{">"}</span>
                <span>[x]</span>
                <span>Focused + selected</span>
              </div>
            </div>
          </ComponentCard>

          {/* Radio */}
          <ComponentCard name="Radio" description="Single selection">
            <div className="space-y-1">
              <div className="flex items-center gap-2">
                <span className="text-tui-green">(●)</span>
                <span>Selected</span>
              </div>
              <div className="flex items-center gap-2">
                <span className="text-muted-foreground">( )</span>
                <span>Unselected</span>
              </div>
            </div>
          </ComponentCard>

          {/* Text Input */}
          <ComponentCard name="Text Input" description="Editable text field with cursor">
            <div className="space-y-2">
              <div className="flex items-center bg-secondary px-2 py-1">
                <span>/home/dev/workspace</span>
                <span className="cursor-blink text-tui-cyan">█</span>
              </div>
              <p className="text-xs text-muted-foreground">
                Edit mode: blinking cursor, Esc to cancel, Enter to save
              </p>
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Border Styles - Lipgloss Inspired */}
      <ComponentSection title="Border Styles (Lipgloss)">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
          <ComponentCard name="Normal Border" description="Standard box-drawing characters">
            <pre className="font-mono text-xs text-tui-cyan">{`┌──────────────┐
│ Content      │
│ Area         │
└──────────────┘`}</pre>
          </ComponentCard>

          <ComponentCard name="Rounded Border" description="Rounded corners">
            <pre className="font-mono text-xs text-tui-cyan">{`╭──────────────╮
│ Content      │
│ Area         │
╰──────────────╯`}</pre>
          </ComponentCard>

          <ComponentCard name="Double Border" description="Double-line box">
            <pre className="font-mono text-xs text-tui-cyan">{`╔══════════════╗
║ Content      ║
║ Area         ║
╚══════════════╝`}</pre>
          </ComponentCard>

          <ComponentCard name="Thick/Heavy Border" description="Bold lines">
            <pre className="font-mono text-xs text-tui-cyan">{`┏━━━━━━━━━━━━━━┓
┃ Content      ┃
┃ Area         ┃
┗━━━━━━━━━━━━━━┛`}</pre>
          </ComponentCard>

          <ComponentCard name="Block Border" description="Full block characters">
            <pre className="font-mono text-xs text-tui-cyan">{`████████████████
█              █
█   Content    █
█              █
████████████████`}</pre>
          </ComponentCard>

          <ComponentCard name="Half-Block Border" description="Half-block outer frame">
            <pre className="font-mono text-xs text-tui-cyan">{`▛▀▀▀▀▀▀▀▀▀▀▀▀▀▀▜
▌              ▐
▌   Content    ▐
▌              ▐
▙▄▄▄▄▄▄▄▄▄▄▄▄▄▄▟`}</pre>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Progress Bars - Lipgloss/Bubbles Inspired */}
      <ComponentSection title="Progress Bars (Bubbles)">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <ComponentCard name="Block Progress" description="Using full/empty blocks">
            <div className="space-y-2 font-mono text-sm">
              <div className="flex items-center gap-2">
                <span className="text-muted-foreground">0%</span>
                <span className="text-muted-foreground">░░░░░░░░░░░░░░░░░░░░</span>
              </div>
              <div className="flex items-center gap-2">
                <span className="text-muted-foreground">25%</span>
                <span><span className="text-tui-cyan">█████</span><span className="text-muted-foreground">░░░░░░░░░░░░░░░</span></span>
              </div>
              <div className="flex items-center gap-2">
                <span className="text-muted-foreground">50%</span>
                <span><span className="text-tui-cyan">██████████</span><span className="text-muted-foreground">░░░░░░░░░░</span></span>
              </div>
              <div className="flex items-center gap-2">
                <span className="text-muted-foreground">75%</span>
                <span><span className="text-tui-cyan">███████████████</span><span className="text-muted-foreground">░░░░░</span></span>
              </div>
              <div className="flex items-center gap-2">
                <span className="text-tui-green">100%</span>
                <span className="text-tui-green">████████████████████</span>
              </div>
            </div>
          </ComponentCard>

          <ComponentCard name="Gradient Progress" description="Color gradient fill">
            <div className="space-y-2 font-mono text-sm">
              <div className="flex items-center gap-2">
                <span>Build:</span>
                <span>
                  <span style={{color: '#5A56E0'}}>██</span>
                  <span style={{color: '#7B5AE4'}}>██</span>
                  <span style={{color: '#9C5EE8'}}>██</span>
                  <span style={{color: '#BD62EC'}}>██</span>
                  <span style={{color: '#DE66F0'}}>██</span>
                  <span style={{color: '#EE6FF8'}}>██</span>
                  <span className="text-muted-foreground">░░░░░░░░</span>
                </span>
                <span className="text-muted-foreground">60%</span>
              </div>
              <div className="flex items-center gap-2">
                <span>Test:</span>
                <span>
                  <span style={{color: '#F25D94'}}>███</span>
                  <span style={{color: '#F49D77'}}>███</span>
                  <span style={{color: '#EDDD5B'}}>███</span>
                  <span style={{color: '#EDFF82'}}>███</span>
                  <span className="text-muted-foreground">░░░░░░░░</span>
                </span>
                <span className="text-muted-foreground">60%</span>
              </div>
            </div>
          </ComponentCard>

          <ComponentCard name="Braille Progress" description="Using braille dot patterns">
            <div className="space-y-2 font-mono text-sm">
              <div className="flex items-center gap-2">
                <span className="text-muted-foreground">⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀</span>
                <span className="text-muted-foreground">0%</span>
              </div>
              <div className="flex items-center gap-2">
                <span><span className="text-tui-cyan">⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿</span><span className="text-muted-foreground">⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀</span></span>
                <span className="text-muted-foreground">50%</span>
              </div>
              <div className="flex items-center gap-2">
                <span className="text-tui-green">⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿</span>
                <span className="text-tui-green">100%</span>
              </div>
            </div>
          </ComponentCard>

          <ComponentCard name="ASCII Progress" description="Simple ASCII characters">
            <div className="space-y-2 font-mono text-sm">
              <div className="flex items-center gap-2">
                <span>[</span>
                <span><span className="text-tui-cyan">====</span><span className="text-muted-foreground">................</span></span>
                <span>]</span>
                <span className="text-muted-foreground">20%</span>
              </div>
              <div className="flex items-center gap-2">
                <span>[</span>
                <span><span className="text-tui-cyan">########</span><span className="text-muted-foreground">............</span></span>
                <span>]</span>
                <span className="text-muted-foreground">40%</span>
              </div>
              <div className="flex items-center gap-2">
                <span>[</span>
                <span><span className="text-tui-cyan">================</span><span className="text-muted-foreground">....</span></span>
                <span>]</span>
                <span className="text-muted-foreground">80%</span>
              </div>
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* List Enumerators - Lipgloss List */}
      <ComponentSection title="List Enumerators (Lipgloss)">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
          <ComponentCard name="Bullet (Default)" description="Simple bullet points">
            <div className="font-mono text-sm space-y-0.5">
              <div><span className="text-tui-cyan">•</span> First item</div>
              <div><span className="text-tui-cyan">•</span> Second item</div>
              <div><span className="text-tui-cyan">•</span> Third item</div>
            </div>
          </ComponentCard>

          <ComponentCard name="Dash" description="Dash enumerator">
            <div className="font-mono text-sm space-y-0.5">
              <div><span className="text-tui-cyan">-</span> First item</div>
              <div><span className="text-tui-cyan">-</span> Second item</div>
              <div><span className="text-tui-cyan">-</span> Third item</div>
            </div>
          </ComponentCard>

          <ComponentCard name="Arrow" description="Arrow enumerator">
            <div className="font-mono text-sm space-y-0.5">
              <div><span className="text-tui-cyan">→</span> First item</div>
              <div><span className="text-tui-cyan">→</span> Second item</div>
              <div><span className="text-tui-cyan">→</span> Third item</div>
            </div>
          </ComponentCard>

          <ComponentCard name="Arabic (Numbered)" description="Numbered list">
            <div className="font-mono text-sm space-y-0.5">
              <div><span className="text-tui-cyan">1.</span> First item</div>
              <div><span className="text-tui-cyan">2.</span> Second item</div>
              <div><span className="text-tui-cyan">3.</span> Third item</div>
            </div>
          </ComponentCard>

          <ComponentCard name="Alphabet" description="Alphabetic list">
            <div className="font-mono text-sm space-y-0.5">
              <div><span className="text-tui-cyan">A.</span> First item</div>
              <div><span className="text-tui-cyan">B.</span> Second item</div>
              <div><span className="text-tui-cyan">C.</span> Third item</div>
            </div>
          </ComponentCard>

          <ComponentCard name="Roman" description="Roman numerals">
            <div className="font-mono text-sm space-y-0.5">
              <div><span className="text-tui-cyan">I.</span>   First item</div>
              <div><span className="text-tui-cyan">II.</span>  Second item</div>
              <div><span className="text-tui-cyan">III.</span> Third item</div>
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Tree Component - Lipgloss Tree */}
      <ComponentSection title="Tree Component (Lipgloss)">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <ComponentCard name="File Tree" description="Tree with box-drawing connectors">
            <div className="font-mono text-sm">
              <div className="text-tui-cyan">rfz-workspace/</div>
              <div><span className="text-muted-foreground">├── </span><span className="text-tui-cyan">rfz-core/</span></div>
              <div><span className="text-muted-foreground">│   ├── </span>pom.xml</div>
              <div><span className="text-muted-foreground">│   └── </span><span className="text-tui-cyan">src/</span></div>
              <div><span className="text-muted-foreground">├── </span><span className="text-tui-cyan">rfz-simulation/</span></div>
              <div><span className="text-muted-foreground">│   ├── </span>pom.xml</div>
              <div><span className="text-muted-foreground">│   └── </span><span className="text-tui-cyan">src/</span></div>
              <div><span className="text-muted-foreground">└── </span><span className="text-tui-cyan">rfz-standalone/</span></div>
              <div><span className="text-muted-foreground">    ├── </span>pom.xml</div>
              <div><span className="text-muted-foreground">    └── </span><span className="text-tui-cyan">src/</span></div>
            </div>
          </ComponentCard>

          <ComponentCard name="Nested List" description="Indented sublists">
            <div className="font-mono text-sm">
              <div><span className="text-tui-green">✓</span> Citrus Fruits</div>
              <div><span className="text-muted-foreground">  │</span></div>
              <div><span className="text-muted-foreground">  ├── </span><span className="line-through text-muted-foreground">Grapefruit</span></div>
              <div><span className="text-muted-foreground">  ├── </span><span className="line-through text-muted-foreground">Yuzu</span></div>
              <div><span className="text-muted-foreground">  ├── </span>Citron</div>
              <div><span className="text-muted-foreground">  └── </span>Kumquat</div>
              <div className="mt-1"><span className="text-tui-cyan">•</span> Actual Items</div>
              <div><span className="text-muted-foreground">  │</span></div>
              <div><span className="text-muted-foreground">  ├── </span>Item A</div>
              <div><span className="text-muted-foreground">  └── </span>Item B</div>
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Tabs - Lipgloss Inspired */}
      <ComponentSection title="Tab Components (Lipgloss)">
        <div className="grid grid-cols-1 gap-4">
          <ComponentCard name="Tab Bar (ASCII)" description="Active tab with bottom opening">
            <pre className="font-mono text-xs">{`┌─────────┐┌──────┐┌───────────┐┌──────────┐
│ `}<span className="text-tui-cyan">Build</span>{`   ││ Logs ││ Discovery ││ Config   │
┘         └┴──────┴┴───────────┴┴──────────┴───`}</pre>
          </ComponentCard>

          <ComponentCard name="Simple Tabs" description="Underline style tabs">
            <div className="flex border-b border-border">
              <div className="px-3 py-1 border-b-2 border-tui-cyan text-tui-cyan">Build</div>
              <div className="px-3 py-1 text-muted-foreground">Logs</div>
              <div className="px-3 py-1 text-muted-foreground">Discovery</div>
              <div className="px-3 py-1 text-muted-foreground">Config</div>
            </div>
          </ComponentCard>

          <ComponentCard name="Pill Tabs" description="Background highlight style">
            <div className="flex gap-1 bg-secondary p-1 rounded">
              <div className="px-3 py-1 bg-tui-cyan text-background rounded text-sm">Build</div>
              <div className="px-3 py-1 text-muted-foreground text-sm">Logs</div>
              <div className="px-3 py-1 text-muted-foreground text-sm">Discovery</div>
              <div className="px-3 py-1 text-muted-foreground text-sm">Config</div>
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Status Bar - Lipgloss Style */}
      <ComponentSection title="Status Bar (Lipgloss)">
        <div className="grid grid-cols-1 gap-4">
          <ComponentCard name="Full Status Bar" description="Colored segments with info">
            <div className="font-mono text-sm flex">
              <div className="bg-brand text-white px-2 py-0.5">STATUS</div>
              <div className="bg-secondary text-foreground px-2 py-0.5 flex-1">Building rfz-core...</div>
              <div className="bg-purple-600 text-white px-2 py-0.5">UTF-8</div>
              <div className="bg-indigo-700 text-white px-2 py-0.5">LF</div>
            </div>
          </ComponentCard>

          <ComponentCard name="Minimal Status" description="Key-value pairs">
            <div className="font-mono text-xs flex gap-4 text-muted-foreground">
              <span><span className="text-tui-cyan">Screen:</span> Build</span>
              <span className="text-border">│</span>
              <span><span className="text-tui-cyan">Focus:</span> Content</span>
              <span className="text-border">│</span>
              <span><span className="text-tui-cyan">Component:</span> rfz-core</span>
              <span className="text-border">│</span>
              <span><span className="text-tui-green">Ready</span></span>
            </div>
          </ComponentCard>
        </div>
      </ComponentSection>

      {/* Spinner Variants */}
      <ComponentSection title="Spinner Variants">
        <ComponentCard name="Spinner Styles" description="Different animation patterns">
          <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 font-mono text-sm">
            <div className="flex items-center gap-2">
              <span className="text-tui-cyan">⠋</span>
              <span className="text-muted-foreground">Braille (dots)</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-tui-cyan">|</span>
              <span className="text-muted-foreground">Line (|/-\)</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-tui-cyan">◐</span>
              <span className="text-muted-foreground">Circle quarters</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-tui-cyan">▁</span>
              <span className="text-muted-foreground">Bounce (▁▂▃▄▅▆▇█)</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-tui-cyan">⣾</span>
              <span className="text-muted-foreground">Braille (spin)</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-tui-cyan">◴</span>
              <span className="text-muted-foreground">Arc (◴◷◶◵)</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-tui-cyan">⠁</span>
              <span className="text-muted-foreground">Dot pulse</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-tui-cyan">[=   ]</span>
              <span className="text-muted-foreground">Marquee</span>
            </div>
          </div>
        </ComponentCard>
      </ComponentSection>

      {/* Dialog Box - Lipgloss */}
      <ComponentSection title="Dialog Box (Lipgloss)">
        <ComponentCard name="Confirmation Dialog" description="Centered dialog with buttons">
          <div className="flex justify-center">
            <pre className="font-mono text-xs text-tui-cyan">{`╭──────────────────────────────────────╮
│                                      │
│  Are you sure you want to rebuild?   │
│                                      │
│         [ Yes ]     [ No ]           │
│                                      │
╰──────────────────────────────────────╯`}</pre>
          </div>
        </ComponentCard>

        <ComponentCard name="Error Dialog" description="Error message with icon">
          <div className="flex justify-center">
            <pre className="font-mono text-xs text-destructive">{`┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                      ┃
┃  ✗ Build failed for rfz-core         ┃
┃                                      ┃
┃  Exit code: 1                        ┃
┃  See logs for details                ┃
┃                                      ┃
┃              [ OK ]                  ┃
┃                                      ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛`}</pre>
          </div>
        </ComponentCard>
      </ComponentSection>

      {/* Keyboard Hints */}
      <ComponentSection title="Keyboard Hint Patterns">
        <ComponentCard name="Shortcut Display" description="Consistent key hint styling">
          <div className="space-y-2 text-sm">
            <p><span className="text-tui-cyan">↑/k</span> Up | <span className="text-tui-cyan">↓/j</span> Down</p>
            <p><span className="text-tui-cyan">Enter</span> Select | <span className="text-tui-cyan">Space</span> Toggle</p>
            <p><span className="text-tui-cyan">Tab</span> Switch panel | <span className="text-tui-cyan">Esc</span> Back/Cancel</p>
            <p><span className="text-tui-cyan">1-5</span> Quick nav | <span className="text-tui-cyan">q</span> Quit</p>
          </div>
        </ComponentCard>
      </ComponentSection>

      {/* User Flow Summary */}
      <ComponentSection title="User Flow (Brief)">
        <TuiBox title="Primary Flows" borderStyle="single">
          <div className="p-2 text-sm space-y-1">
            <p className="text-muted-foreground">1. Navigate sidebar → Select screen → Interact with content</p>
            <p className="text-muted-foreground">2. Build: Select components → Configure → Execute → View logs</p>
          </div>
        </TuiBox>
      </ComponentSection>

      {/* Modal Preview */}
      <TuiModal title="Sample Modal" isOpen={showModal}>
        <p className="text-sm mb-4">Modal content with double border and backdrop overlay.</p>
        <TuiList
          items={[
            { id: "opt1", label: "Option One" },
            { id: "opt2", label: "Option Two" },
          ]}
          selectedIndex={0}
        />
        <TuiModalActions>
          <TuiButton label="Cancel" shortcut="Esc" onClick={() => setShowModal(false)} />
          <TuiButton label="Confirm" shortcut="Enter" variant="primary" focused onClick={() => setShowModal(false)} />
        </TuiModalActions>
      </TuiModal>
    </div>
  )
}
