"use client"

import { cn } from "@/lib/utils"
import { TuiBox } from "../tui/tui-box"
import { TuiButton } from "../tui/tui-modal"

export interface ComponentInfo {
  id: string
  name: string
  category: string
  branch: string
  status: "clean" | "dirty" | "unknown"
  lastCommit?: string
  lastCommitDate?: string
}

const SAMPLE_COMPONENTS: ComponentInfo[] = [
  { id: "boss", name: "boss", category: "Core", branch: "develop", status: "clean", lastCommit: "a3f2b1c", lastCommitDate: "2025-01-28" },
  { id: "fistiv", name: "fistiv", category: "Core", branch: "develop", status: "dirty", lastCommit: "b4c3d2e", lastCommitDate: "2025-01-27" },
  { id: "audiocon", name: "audiocon", category: "Standalone", branch: "feature/audio-fix", status: "clean", lastCommit: "c5d4e3f", lastCommitDate: "2025-01-26" },
  { id: "traktion", name: "traktion", category: "Core", branch: "develop", status: "clean", lastCommit: "d6e5f4g", lastCommitDate: "2025-01-28" },
  { id: "signalsteuerung", name: "signalsteuerung", category: "Core", branch: "develop", status: "clean", lastCommit: "e7f6g5h", lastCommitDate: "2025-01-25" },
  { id: "weichensteuerung", name: "weichensteuerung", category: "Core", branch: "bugfix/weiche-123", status: "dirty", lastCommit: "f8g7h6i", lastCommitDate: "2025-01-24" },
  { id: "simkern", name: "simkern", category: "Simulation", branch: "develop", status: "clean", lastCommit: "g9h8i7j", lastCommitDate: "2025-01-28" },
  { id: "fahrdynamik", name: "fahrdynamik", category: "Simulation", branch: "develop", status: "clean", lastCommit: "h0i9j8k", lastCommitDate: "2025-01-27" },
  { id: "energierechnung", name: "energierechnung", category: "Simulation", branch: "develop", status: "unknown", lastCommit: "i1j0k9l", lastCommitDate: "2025-01-23" },
  { id: "zuglauf", name: "zuglauf", category: "Simulation", branch: "develop", status: "clean", lastCommit: "j2k1l0m", lastCommitDate: "2025-01-28" },
]

interface DiscoverScreenProps {
  selectedIndex: number
  focusArea: "list" | "actions"
  actionIndex: number
  className?: string
}

export function DiscoverScreen({
  selectedIndex,
  focusArea,
  actionIndex,
}: DiscoverScreenProps) {
  const dirtyCount = SAMPLE_COMPONENTS.filter(c => c.status === "dirty").length
  const cleanCount = SAMPLE_COMPONENTS.filter(c => c.status === "clean").length
  const nonDevelopCount = SAMPLE_COMPONENTS.filter(c => c.branch !== "develop").length

  return (
    <div className="flex flex-col h-full">
      {/* Header */}
      <TuiBox title="Component List / Discover Components" borderStyle="double">
        <div className="p-2">
          <div className="flex items-center gap-6 text-sm">
            <span>
              <span className="text-muted-foreground">Total:</span>{" "}
              <span className="text-foreground">{SAMPLE_COMPONENTS.length}</span>
            </span>
            <span>
              <span className="text-muted-foreground">Clean:</span>{" "}
              <span className="text-tui-green">{cleanCount}</span>
            </span>
            <span>
              <span className="text-muted-foreground">Dirty:</span>{" "}
              <span className="text-tui-yellow">{dirtyCount}</span>
            </span>
            <span>
              <span className="text-muted-foreground">Not on develop:</span>{" "}
              <span className="text-tui-cyan">{nonDevelopCount}</span>
            </span>
          </div>
        </div>
      </TuiBox>

      {/* Component table */}
      <div className="flex-1 mt-2 overflow-hidden">
        <TuiBox 
          title="Components" 
          borderStyle="single" 
          focused={focusArea === "list"}
        >
          <div className="p-1">
            {/* Table header */}
            <div className="flex items-center gap-4 px-2 py-1 text-sm text-muted-foreground border-b border-border">
              <span className="w-4 flex-shrink-0"></span>
              <span className="w-44 flex-shrink-0">Component</span>
              <span className="w-20 flex-shrink-0">Category</span>
              <span className="w-36 flex-shrink-0">Branch</span>
              <span className="w-12 text-center flex-shrink-0">Status</span>
              <span className="w-16 flex-shrink-0">Commit</span>
              <span className="w-24 text-right flex-shrink-0">Date</span>
            </div>

            {/* Table body */}
            <div className="max-h-72 overflow-y-auto">
              {SAMPLE_COMPONENTS.map((component, index) => {
                const isFocused = focusArea === "list" && index === selectedIndex
                
                return (
                  <div
                    key={component.id}
                    className={cn(
                      "flex items-center gap-4 px-2 py-0.5 text-sm",
                      isFocused && "bg-tui-cyan text-background"
                    )}
                  >
                    {/* Focus indicator */}
                    <span className={cn(
                      "w-4 flex-shrink-0",
                      isFocused ? "text-background" : "text-tui-cyan"
                    )}>
                      {isFocused ? ">" : " "}
                    </span>
                    
                    {/* Component name */}
                    <span className="w-44 flex-shrink-0 truncate">{component.name}</span>
                    
                    {/* Category */}
                    <span className={cn(
                      "w-20 flex-shrink-0 text-xs truncate",
                      isFocused ? "text-background/80" : "text-muted-foreground"
                    )}>
                      {component.category}
                    </span>
                    
                    {/* Branch */}
                    <span className={cn(
                      "w-36 flex-shrink-0 truncate",
                      isFocused ? "text-background" : component.branch === "develop" ? "text-tui-green" : "text-tui-cyan"
                    )}>
                      {component.branch}
                    </span>
                    
                    {/* Status */}
                    <span className={cn(
                      "w-12 text-center flex-shrink-0",
                      isFocused ? "text-background" : {
                        "text-tui-green": component.status === "clean",
                        "text-tui-yellow": component.status === "dirty",
                        "text-muted-foreground": component.status === "unknown"
                      }[component.status]
                    )}>
                      {component.status === "clean" && "✓"}
                      {component.status === "dirty" && "●"}
                      {component.status === "unknown" && "?"}
                    </span>
                    
                    {/* Last commit */}
                    <span className={cn(
                      "w-16 flex-shrink-0 font-mono text-xs",
                      isFocused ? "text-background/80" : "text-muted-foreground"
                    )}>
                      {component.lastCommit}
                    </span>
                    
                    {/* Date */}
                    <span className={cn(
                      "w-24 text-right flex-shrink-0 text-xs",
                      isFocused ? "text-background/80" : "text-muted-foreground"
                    )}>
                      {component.lastCommitDate}
                    </span>
                  </div>
                )
              })}
            </div>
          </div>
        </TuiBox>
      </div>

      {/* Legend */}
      <div className="mt-2 px-2 text-xs text-muted-foreground flex gap-6">
        <span><span className="text-tui-green">✓</span> Clean (no uncommitted changes)</span>
        <span><span className="text-tui-yellow">●</span> Dirty (uncommitted changes)</span>
        <span><span className="text-muted-foreground">?</span> Unknown status</span>
      </div>

      {/* Actions */}
      <div className="mt-2">
        <TuiBox title="Git Actions" borderStyle="single" focused={focusArea === "actions"}>
          <div className="p-2 flex items-center gap-4">
            <TuiButton
              label="Fetch All Develop"
              shortcut="f"
              focused={focusArea === "actions" && actionIndex === 0}
            />
            <TuiButton
              label="Checkout Develop"
              shortcut="c"
              focused={focusArea === "actions" && actionIndex === 1}
            />
            <TuiButton
              label="Update All"
              shortcut="u"
              focused={focusArea === "actions" && actionIndex === 2}
              variant="primary"
            />
            <TuiButton
              label="Git Status"
              shortcut="s"
              focused={focusArea === "actions" && actionIndex === 3}
            />
          </div>
        </TuiBox>
      </div>
    </div>
  )
}

export { SAMPLE_COMPONENTS }
