"use client"

import { cn } from "@/lib/utils"
import { TuiBox, TuiDivider } from "../tui/tui-box"
import { TuiProgress, TuiBrailleProgress, TuiSpinner } from "../tui/tui-progress"
import { TuiButton } from "../tui/tui-modal"

export type BuildPhase = "Pending" | "Compiling" | "Testing" | "Packaging" | "Installing" | "Done"
export type BuildStatus = "pending" | "running" | "success" | "failed" | "error"

export interface ComponentBuildState {
  id: string
  name: string
  phase: BuildPhase
  status: BuildStatus
  elapsedTime: number // seconds
  logs?: string[]
}

interface BuildExecutionViewProps {
  builds: ComponentBuildState[]
  command: string
  selectedIndex: number
  focusArea: "list" | "actions"
  actionIndex: number
  isComplete: boolean
  onViewLogs?: (build: ComponentBuildState) => void
  onRebuildFailed?: () => void
  onBackToSelection?: () => void
  onCancelBuild?: () => void
  className?: string
}

const STATUS_SYMBOLS = {
  pending: "○",
  running: "◐",
  success: "✓",
  failed: "✗",
  error: "!"
}

const PHASE_COLORS = {
  Pending: "text-muted-foreground",
  Compiling: "text-tui-yellow",
  Testing: "text-tui-blue",
  Packaging: "text-tui-cyan",
  Installing: "text-tui-cyan",
  Done: "text-tui-green"
}

function formatTime(seconds: number): string {
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins.toString().padStart(2, "0")}:${secs.toString().padStart(2, "0")}`
}

export function BuildExecutionView({
  builds,
  command,
  selectedIndex,
  focusArea,
  actionIndex,
  isComplete,
  onViewLogs,
  onRebuildFailed,
  onBackToSelection,
  onCancelBuild,
}: BuildExecutionViewProps) {
  const completedCount = builds.filter(b => b.status === "success").length
  const failedCount = builds.filter(b => b.status === "failed" || b.status === "error").length
  const runningCount = builds.filter(b => b.status === "running").length
  const pendingCount = builds.filter(b => b.status === "pending").length
  
  const totalProgress = builds.length > 0 
    ? ((completedCount + failedCount) / builds.length) * 100 
    : 0

  return (
    <div className="flex flex-col h-full">
      {/* Header with command - heavy border for emphasis */}
      <TuiBox title="Build Execution" borderStyle="heavy">
        <div className="p-2">
          <div className="flex items-center gap-2 text-sm font-mono">
            <span className="text-tui-green">$</span>
            <span className="text-tui-cyan">{command}</span>
            {!isComplete && <TuiSpinner />}
          </div>
        </div>
      </TuiBox>

      {/* Build status table with tree-style connectors */}
      <div className="flex-1 mt-2 overflow-hidden">
        <TuiBox 
          title="Components" 
          borderStyle="rounded" 
          focused={focusArea === "list"}
        >
          <div className="p-1">
            {/* Table header */}
            <div className="flex items-center gap-2 px-2 py-1 text-xs text-muted-foreground font-mono">
              <span className="w-6"></span>
              <span className="w-4">St</span>
              <span className="flex-1">Component</span>
              <span className="w-24 text-center">Phase</span>
              <span className="w-20">Progress</span>
              <span className="w-14 text-right">Time</span>
            </div>
            <TuiDivider style="dashed" className="mx-2" />

            {/* Table body with tree connectors */}
            <div className="max-h-64 overflow-y-auto font-mono">
              {builds.map((build, index) => {
                const isFocused = focusArea === "list" && index === selectedIndex
                const isLast = index === builds.length - 1
                const treeChar = isLast ? "└── " : "├── "
                
                // Calculate individual progress based on phase
                const phaseProgress = {
                  Pending: 0,
                  Compiling: 25,
                  Testing: 50,
                  Packaging: 75,
                  Installing: 90,
                  Done: 100
                }[build.phase]
                
                return (
                  <div
                    key={build.id}
                    className={cn(
                      "flex items-center gap-2 px-2 py-0.5 text-sm",
                      isFocused && "bg-tui-cyan text-background"
                    )}
                  >
                    {/* Tree connector */}
                    <span className={cn(
                      "w-6 whitespace-pre",
                      isFocused ? "text-background/60" : "text-border"
                    )}>
                      {isFocused ? "> " : treeChar.slice(0, 2)}
                    </span>
                    
                    {/* Status icon */}
                    <span className={cn(
                      "w-4",
                      isFocused ? "text-background" : {
                        "text-muted-foreground": build.status === "pending",
                        "text-tui-yellow": build.status === "running",
                        "text-tui-green": build.status === "success",
                        "text-destructive": build.status === "failed" || build.status === "error"
                      }[build.status]
                    )}>
                      {build.status === "running" ? (
                        <TuiSpinner />
                      ) : (
                        STATUS_SYMBOLS[build.status]
                      )}
                    </span>
                    
                    {/* Component name */}
                    <span className="flex-1">{build.name}</span>
                    
                    {/* Phase */}
                    <span className={cn(
                      "w-24 text-center text-xs",
                      isFocused ? "text-background" : PHASE_COLORS[build.phase]
                    )}>
                      {build.phase}
                    </span>
                    
                    {/* Progress bar */}
                    <span className="w-20">
                      <TuiBrailleProgress 
                        value={build.status === "success" ? 100 : build.status === "failed" ? phaseProgress : phaseProgress} 
                        width={8}
                      />
                    </span>
                    
                    {/* Time */}
                    <span className={cn(
                      "w-14 text-right tabular-nums text-xs",
                      isFocused ? "text-background/80" : "text-muted-foreground"
                    )}>
                      {formatTime(build.elapsedTime)}
                    </span>
                  </div>
                )
              })}
            </div>
          </div>
        </TuiBox>
      </div>

      {/* Progress footer */}
      <div className="mt-2">
        <TuiBox title="Progress" borderStyle="rounded">
          <div className="p-2 space-y-2">
            {/* Progress bar - cyan when running, green when complete */}
            <TuiProgress 
              value={totalProgress} 
              width={40}
              showPercentage
              label="Overall:"
              variant={isComplete ? "success" : "default"}
            />

            {/* Status counters in a row */}
            <div className="flex items-center gap-4 text-xs font-mono">
              <span className="flex items-center gap-1 px-2 py-0.5 bg-tui-yellow/10 rounded">
                <span className="text-tui-yellow">◐</span>
                <span className="text-muted-foreground">Running:</span>
                <span className="text-tui-yellow font-bold">{runningCount}</span>
              </span>
              <span className="flex items-center gap-1 px-2 py-0.5 bg-tui-green/10 rounded">
                <span className="text-tui-green">✓</span>
                <span className="text-muted-foreground">Success:</span>
                <span className="text-tui-green font-bold">{completedCount}</span>
              </span>
              <span className="flex items-center gap-1 px-2 py-0.5 bg-destructive/10 rounded">
                <span className="text-destructive">✗</span>
                <span className="text-muted-foreground">Failed:</span>
                <span className="text-destructive font-bold">{failedCount}</span>
              </span>
              <span className="flex items-center gap-1 px-2 py-0.5 bg-secondary rounded">
                <span className="text-muted-foreground">○</span>
                <span className="text-muted-foreground">Pending:</span>
                <span className="text-muted-foreground font-bold">{pendingCount}</span>
              </span>
            </div>
          </div>
        </TuiBox>
      </div>

      {/* Actions */}
      <div className="mt-2">
        <TuiBox title="Actions" borderStyle="single" focused={focusArea === "actions"}>
          <div className="p-2 flex items-center gap-4">
            <TuiButton
              label="View Logs"
              shortcut="l"
              focused={focusArea === "actions" && actionIndex === 0}
              onClick={() => onViewLogs?.(builds[selectedIndex])}
            />
            {isComplete && (
              <>
                <TuiButton
                  label="Rebuild Failed"
                  shortcut="r"
                  focused={focusArea === "actions" && actionIndex === 1}
                  variant={failedCount > 0 ? "danger" : "default"}
                  onClick={onRebuildFailed}
                />
                <TuiButton
                  label="Back to Selection"
                  shortcut="Esc"
                  focused={focusArea === "actions" && actionIndex === 2}
                  onClick={onBackToSelection}
                />
              </>
            )}
            {!isComplete && (
              <TuiButton
                label="Cancel Build"
                shortcut="Ctrl+C"
                focused={focusArea === "actions" && actionIndex === 1}
                variant="danger"
                onClick={onCancelBuild}
              />
            )}
          </div>
        </TuiBox>
      </div>
    </div>
  )
}
