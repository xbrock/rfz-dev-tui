"use client"

import { cn } from "@/lib/utils"
import { TuiModal, TuiModalActions, TuiButton } from "../tui/tui-modal"
import { TuiBox } from "../tui/tui-box"

export interface BuildConfig {
  mavenGoal: string
  profiles: Set<string>
  port: string
  skipTests: boolean
}

const MAVEN_GOALS = [
  { id: "clean", label: "clean" },
  { id: "install", label: "install" },
  { id: "package", label: "package" },
  { id: "clean-install", label: "clean install" },
]

const MAVEN_PROFILES = [
  { id: "target_env_dev", label: "target_env_dev" },
  { id: "generate_local_config_files", label: "generate_local_config_files" },
]

const PORT_OPTIONS = [
  { id: "11090", label: "Port 11090", profile: "use_traktion_11090" },
  { id: "11091", label: "Port 11091", profile: "use_traktion_11091" },
]

interface BuildConfigModalProps {
  isOpen: boolean
  config: BuildConfig
  focusArea: "goals" | "profiles" | "port" | "skipTests" | "actions"
  goalIndex: number
  profileIndex: number
  portIndex: number
  actionIndex: number
  selectedComponents: string[]
  onClose?: () => void
  onStartBuild?: () => void
  onToggleSkipTests?: () => void
  className?: string
}

export function BuildConfigModal({
  isOpen,
  config,
  focusArea,
  goalIndex,
  profileIndex,
  portIndex,
  actionIndex,
  selectedComponents,
  onClose,
  onStartBuild,
  onToggleSkipTests,
}: BuildConfigModalProps) {
  // Generate command preview
  const generateCommand = () => {
    const profiles = Array.from(config.profiles)
    const portProfile = PORT_OPTIONS.find(p => p.id === config.port)?.profile
    if (portProfile) {
      profiles.push(portProfile)
    }
    const profileStr = profiles.length > 0 ? ` -P${profiles.join(",")}` : ""
    const skipTestsStr = config.skipTests ? " -DskipTests" : ""
    return `mvn ${config.mavenGoal}${profileStr}${skipTestsStr}`
  }

  return (
    <TuiModal 
      isOpen={isOpen} 
      title="Build Configuration"
      width="max-w-3xl"
    >
      <div className="space-y-4">
        {/* Selected components summary */}
        <div className="text-sm text-muted-foreground">
          Building {selectedComponents.length} component{selectedComponents.length !== 1 ? "s" : ""}:{" "}
          <span className="text-foreground">
            {selectedComponents.slice(0, 5).join(", ")}
            {selectedComponents.length > 5 && ` +${selectedComponents.length - 5} more`}
          </span>
        </div>

        {/* Maven Goal Selection */}
        <TuiBox 
          title="Maven Goal" 
          borderStyle="single" 
          focused={focusArea === "goals"}
        >
          <div className="p-2">
            <div className="flex flex-wrap gap-2">
              {MAVEN_GOALS.map((goal, index) => {
                const isSelected = config.mavenGoal === goal.id
                const isFocused = focusArea === "goals" && index === goalIndex
                
                return (
                  <span
                    key={goal.id}
                    className={cn(
                      "px-2 py-0.5",
                      isFocused && "bg-tui-cyan text-background",
                      !isFocused && isSelected && "text-tui-green",
                      !isFocused && !isSelected && "text-muted-foreground"
                    )}
                  >
                    {isFocused && ">"} 
                    {isSelected ? "(●)" : "( )"} {goal.label}
                  </span>
                )
              })}
            </div>
            <div className="mt-2 text-xs text-muted-foreground">
              <span className="text-tui-cyan">←→</span> or <span className="text-tui-cyan">h/l</span> to select
            </div>
          </div>
        </TuiBox>

        {/* Maven Profiles Selection */}
        <TuiBox 
          title="Maven Profiles (multi-select)" 
          borderStyle="single" 
          focused={focusArea === "profiles"}
        >
          <div className="p-2">
            <div className="space-y-1">
              {MAVEN_PROFILES.map((profile, index) => {
                const isSelected = config.profiles.has(profile.id)
                const isFocused = focusArea === "profiles" && index === profileIndex
                
                return (
                  <div
                    key={profile.id}
                    className={cn(
                      "flex items-center gap-2 px-1",
                      isFocused && "bg-tui-cyan text-background"
                    )}
                  >
                    <span className={cn(
                      "w-2",
                      isFocused ? "text-background" : "text-tui-cyan"
                    )}>
                      {isFocused ? ">" : " "}
                    </span>
                    <span className={cn(
                      isFocused ? "text-background" : isSelected ? "text-tui-green" : "text-muted-foreground"
                    )}>
                      {isSelected ? "[x]" : "[ ]"}
                    </span>
                    <span>{profile.label}</span>
                  </div>
                )
              })}
            </div>
            <div className="mt-2 text-xs text-muted-foreground">
              <span className="text-tui-cyan">↑↓</span> navigate | <span className="text-tui-cyan">Space</span> toggle
            </div>
          </div>
        </TuiBox>

        {/* Port Selection */}
        <TuiBox 
          title="Traktion Port" 
          borderStyle="single" 
          focused={focusArea === "port"}
        >
          <div className="p-2">
            <div className="flex flex-wrap gap-4">
              {PORT_OPTIONS.map((port, index) => {
                const isSelected = config.port === port.id
                const isFocused = focusArea === "port" && index === portIndex
                
                return (
                  <span
                    key={port.id}
                    className={cn(
                      "px-2 py-0.5",
                      isFocused && "bg-tui-cyan text-background",
                      !isFocused && isSelected && "text-tui-green",
                      !isFocused && !isSelected && "text-muted-foreground"
                    )}
                  >
                    {isFocused && ">"} 
                    {isSelected ? "(●)" : "( )"} {port.label}
                  </span>
                )
              })}
            </div>
            <div className="mt-2 text-xs text-muted-foreground">
              <span className="text-tui-cyan">←→</span> or <span className="text-tui-cyan">h/l</span> to select | Appends <span className="text-tui-green">use_traktion_*</span> profile
            </div>
          </div>
        </TuiBox>

        {/* Skip Tests Toggle */}
        <TuiBox 
          title="Build Options" 
          borderStyle="single" 
          focused={focusArea === "skipTests"}
        >
          <div className="p-2">
            <button
              type="button"
              onClick={onToggleSkipTests}
              className={cn(
                "flex items-center gap-2 px-2 py-0.5 cursor-pointer",
                focusArea === "skipTests" && "bg-tui-cyan text-background"
              )}
            >
              <span className={cn(
                "w-2",
                focusArea === "skipTests" ? "text-background" : "text-tui-cyan"
              )}>
                {focusArea === "skipTests" ? ">" : " "}
              </span>
              <span className={cn(
                focusArea === "skipTests" 
                  ? "text-background" 
                  : config.skipTests 
                    ? "text-tui-green" 
                    : "text-muted-foreground"
              )}>
                {config.skipTests ? "[x]" : "[ ]"}
              </span>
              <span>Skip Tests</span>
              <span className={cn(
                "text-xs ml-2",
                focusArea === "skipTests" ? "text-background/70" : "text-muted-foreground"
              )}>
                (adds -DskipTests)
              </span>
            </button>
            <div className="mt-2 text-xs text-muted-foreground">
              <span className="text-tui-cyan">Space</span> or <span className="text-tui-cyan">Enter</span> to toggle
            </div>
          </div>
        </TuiBox>

        {/* Command Preview */}
        <TuiBox title="Command Preview" borderStyle="single">
          <div className="p-2">
            <pre className="text-tui-green text-sm">
              <span className="text-muted-foreground">$</span> {generateCommand()}
            </pre>
          </div>
        </TuiBox>

        {/* Actions */}
        <TuiModalActions>
          <TuiButton
            label="Cancel"
            shortcut="Esc"
            focused={focusArea === "actions" && actionIndex === 0}
            onClick={onClose}
          />
          <TuiButton
            label="Start Build"
            shortcut="Enter"
            focused={focusArea === "actions" && actionIndex === 1}
            variant="primary"
            onClick={onStartBuild}
          />
        </TuiModalActions>

        {/* Navigation hint */}
        <div className="text-xs text-muted-foreground text-center">
          <span className="text-tui-cyan">Tab</span> Switch sections | 
          <span className="text-tui-cyan"> Enter</span> Confirm | 
          <span className="text-tui-cyan"> Esc</span> Cancel
        </div>
      </div>
    </TuiModal>
  )
}

export { MAVEN_GOALS, MAVEN_PROFILES, PORT_OPTIONS }
