"use client"

import { cn } from "@/lib/utils"
import { TuiBox } from "../tui/tui-box"

// Configuration sections
export const CONFIG_SECTIONS = [
  { id: "scan", label: "Scan Configuration", description: "Configure how the CLI searches for components" },
  { id: "registry", label: "Component Registry", description: "View registered component metadata (read-only)" },
  { id: "detected", label: "Detected Components", description: "View runtime component discovery results" },
] as const

export type ConfigSectionId = typeof CONFIG_SECTIONS[number]["id"]

// Scan configuration data
export interface ScanPath {
  id: string
  path: string
  enabled: boolean
}

export const DEFAULT_SCAN_PATHS: ScanPath[] = [
  { id: "1", path: "/home/dev/rfz-workspace", enabled: true },
  { id: "2", path: "/opt/rfz/components", enabled: true },
  { id: "3", path: "./", enabled: false },
]

export interface ScanBehavior {
  recursiveScan: boolean
  maxDepth: number
  excludePatterns: string[]
}

export const DEFAULT_SCAN_BEHAVIOR: ScanBehavior = {
  recursiveScan: true,
  maxDepth: 5,
  excludePatterns: ["*/target/*", "*/.git/*", "*/node_modules/*"],
}

// Component registry data (read-only)
export interface RegisteredComponent {
  name: string
  artifactId: string
  category: "Core" | "Simulator" | "Standalone"
  defaultProfiles: string[]
  supportedPorts: string[]
  buildOrder: number
}

export const COMPONENT_REGISTRY: RegisteredComponent[] = [
  { name: "rfz-api", artifactId: "rfz-api", category: "Core", defaultProfiles: ["target_env_dev"], supportedPorts: ["11090", "11091"], buildOrder: 1 },
  { name: "rfz-core", artifactId: "rfz-core", category: "Core", defaultProfiles: ["target_env_dev"], supportedPorts: ["11090", "11091"], buildOrder: 2 },
  { name: "rfz-domain", artifactId: "rfz-domain", category: "Core", defaultProfiles: ["target_env_dev"], supportedPorts: ["11090", "11091"], buildOrder: 3 },
  { name: "rfz-persistence", artifactId: "rfz-persistence", category: "Core", defaultProfiles: ["target_env_dev"], supportedPorts: ["11090", "11091"], buildOrder: 4 },
  { name: "rfz-service", artifactId: "rfz-service", category: "Core", defaultProfiles: ["target_env_dev"], supportedPorts: ["11090", "11091"], buildOrder: 5 },
  { name: "rfz-web", artifactId: "rfz-web", category: "Core", defaultProfiles: ["target_env_dev"], supportedPorts: ["11090", "11091"], buildOrder: 6 },
  { name: "rfz-simulator", artifactId: "rfz-simulator", category: "Simulator", defaultProfiles: ["target_env_dev", "simulator"], supportedPorts: ["11090"], buildOrder: 10 },
  { name: "rfz-batch", artifactId: "rfz-batch", category: "Standalone", defaultProfiles: ["target_env_dev"], supportedPorts: [], buildOrder: 20 },
]

// Detected components (runtime)
export interface DetectedComponent {
  name: string
  category: "Core" | "Simulator" | "Standalone"
  found: boolean
  path: string | null
  hasConfigModule: boolean
}

export const DETECTED_COMPONENTS: DetectedComponent[] = [
  { name: "rfz-api", category: "Core", found: true, path: "/home/dev/rfz-workspace/rfz-api", hasConfigModule: true },
  { name: "rfz-core", category: "Core", found: true, path: "/home/dev/rfz-workspace/rfz-core", hasConfigModule: true },
  { name: "rfz-domain", category: "Core", found: true, path: "/home/dev/rfz-workspace/rfz-domain", hasConfigModule: false },
  { name: "rfz-persistence", category: "Core", found: true, path: "/home/dev/rfz-workspace/rfz-persistence", hasConfigModule: true },
  { name: "rfz-service", category: "Core", found: true, path: "/home/dev/rfz-workspace/rfz-service", hasConfigModule: true },
  { name: "rfz-web", category: "Core", found: true, path: "/home/dev/rfz-workspace/rfz-web", hasConfigModule: true },
  { name: "rfz-simulator", category: "Simulator", found: false, path: null, hasConfigModule: false },
  { name: "rfz-batch", category: "Standalone", found: false, path: null, hasConfigModule: false },
]

interface ConfigScreenProps {
  sectionIndex: number
  focusArea: "sections" | "details"
  detailIndex: number
  editMode: boolean
  editValue: string
  pathIndex: number
  behaviorIndex: number
  registryIndex: number
  detectedIndex: number
  className?: string
}

export function ConfigScreen({
  sectionIndex,
  focusArea,
  detailIndex,
  editMode,
  editValue,
  pathIndex,
  behaviorIndex,
  registryIndex,
  detectedIndex,
}: ConfigScreenProps) {
  const currentSection = CONFIG_SECTIONS[sectionIndex]

  return (
    <div className="flex h-full gap-2">
      {/* Left Panel - Sections */}
      <div className="w-64 flex-shrink-0">
        <TuiBox 
          title="Configuration" 
          borderStyle="double" 
          focused={focusArea === "sections"}
        >
          <div className="py-1">
            {CONFIG_SECTIONS.map((section, index) => {
              const isFocused = focusArea === "sections" && index === sectionIndex
              const isActive = index === sectionIndex
              
              return (
                <div
                  key={section.id}
                  className={cn(
                    "flex items-center gap-2 px-2 py-1",
                    isFocused && "bg-tui-cyan text-background",
                    !isFocused && isActive && "bg-secondary text-tui-cyan"
                  )}
                >
                  <span className={cn(
                    "w-2",
                    isFocused ? "text-background" : isActive ? "text-tui-cyan" : "text-tui-cyan"
                  )}>
                    {isFocused || isActive ? ">" : " "}
                  </span>
                  <span>{section.label}</span>
                </div>
              )
            })}
          </div>
          
          {/* Section description */}
          <div className="mt-2 px-2 py-1 border-t border-border">
            <p className="text-xs text-muted-foreground">
              {currentSection.description}
            </p>
          </div>
          
          {/* Help */}
          <div className="mt-2 px-2 py-1 border-t border-border">
            <div className="text-xs text-muted-foreground space-y-1">
              <p><span className="text-tui-cyan">↑/k</span> Up</p>
              <p><span className="text-tui-cyan">↓/j</span> Down</p>
              <p><span className="text-tui-cyan">Tab</span> Switch panel</p>
              <p><span className="text-tui-cyan">Enter</span> Edit/Select</p>
            </div>
          </div>
        </TuiBox>
      </div>

      {/* Right Panel - Details */}
      <div className="flex-1 min-w-0">
        <TuiBox 
          title={currentSection.label}
          borderStyle="single"
          focused={focusArea === "details"}
        >
          <div className="p-2 h-full overflow-y-auto">
            {/* Scan Configuration Section */}
            {currentSection.id === "scan" && (
              <div className="space-y-4">
                {/* Scan Paths */}
                <div>
                  <div className="text-sm font-bold text-tui-cyan mb-2">Scan Paths</div>
                  <div className="space-y-1">
                    {DEFAULT_SCAN_PATHS.map((scanPath, index) => {
                      const isFocused = focusArea === "details" && detailIndex === 0 && index === pathIndex
                      const isEditing = isFocused && editMode
                      
                      return (
                        <div
                          key={scanPath.id}
                          className={cn(
                            "flex items-center gap-2 px-1 py-0.5",
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
                            isFocused ? "text-background" : scanPath.enabled ? "text-tui-green" : "text-muted-foreground"
                          )}>
                            {scanPath.enabled ? "[x]" : "[ ]"}
                          </span>
                          {isEditing ? (
                            <span className="flex-1 bg-background text-foreground px-1">
                              {editValue}<span className="animate-pulse">_</span>
                            </span>
                          ) : (
                            <span className={cn(
                              "flex-1 font-mono text-sm",
                              isFocused ? "text-background" : "text-foreground"
                            )}>
                              {scanPath.path}
                            </span>
                          )}
                        </div>
                      )
                    })}
                  </div>
                  <div className="mt-1 text-xs text-muted-foreground">
                    <span className="text-tui-cyan">a</span> Add path | 
                    <span className="text-tui-cyan"> d</span> Delete | 
                    <span className="text-tui-cyan"> Space</span> Toggle | 
                    <span className="text-tui-cyan"> Enter</span> Edit
                  </div>
                </div>

                {/* Scan Behavior */}
                <div>
                  <div className="text-sm font-bold text-tui-cyan mb-2">Scan Behavior</div>
                  <div className="space-y-1">
                    {/* Recursive Scan */}
                    <div className={cn(
                      "flex items-center gap-2 px-1 py-0.5",
                      focusArea === "details" && detailIndex === 1 && behaviorIndex === 0 && "bg-tui-cyan text-background"
                    )}>
                      <span className={cn(
                        "w-2",
                        focusArea === "details" && detailIndex === 1 && behaviorIndex === 0 ? "text-background" : "text-tui-cyan"
                      )}>
                        {focusArea === "details" && detailIndex === 1 && behaviorIndex === 0 ? ">" : " "}
                      </span>
                      <span className={cn(
                        focusArea === "details" && detailIndex === 1 && behaviorIndex === 0 
                          ? "text-background" 
                          : DEFAULT_SCAN_BEHAVIOR.recursiveScan 
                            ? "text-tui-green" 
                            : "text-muted-foreground"
                      )}>
                        {DEFAULT_SCAN_BEHAVIOR.recursiveScan ? "[x]" : "[ ]"}
                      </span>
                      <span>Recursive scan</span>
                    </div>
                    
                    {/* Max Depth */}
                    <div className={cn(
                      "flex items-center gap-2 px-1 py-0.5",
                      focusArea === "details" && detailIndex === 1 && behaviorIndex === 1 && "bg-tui-cyan text-background"
                    )}>
                      <span className={cn(
                        "w-2",
                        focusArea === "details" && detailIndex === 1 && behaviorIndex === 1 ? "text-background" : "text-tui-cyan"
                      )}>
                        {focusArea === "details" && detailIndex === 1 && behaviorIndex === 1 ? ">" : " "}
                      </span>
                      <span className="w-32">Max depth:</span>
                      <span className={cn(
                        "font-mono",
                        focusArea === "details" && detailIndex === 1 && behaviorIndex === 1 ? "text-background" : "text-tui-green"
                      )}>
                        {DEFAULT_SCAN_BEHAVIOR.maxDepth}
                      </span>
                    </div>
                    
                    {/* Exclude Patterns */}
                    <div className={cn(
                      "flex items-start gap-2 px-1 py-0.5",
                      focusArea === "details" && detailIndex === 1 && behaviorIndex === 2 && "bg-tui-cyan text-background"
                    )}>
                      <span className={cn(
                        "w-2 flex-shrink-0",
                        focusArea === "details" && detailIndex === 1 && behaviorIndex === 2 ? "text-background" : "text-tui-cyan"
                      )}>
                        {focusArea === "details" && detailIndex === 1 && behaviorIndex === 2 ? ">" : " "}
                      </span>
                      <span className="w-32 flex-shrink-0">Exclude:</span>
                      <span className={cn(
                        "font-mono text-sm",
                        focusArea === "details" && detailIndex === 1 && behaviorIndex === 2 ? "text-background" : "text-muted-foreground"
                      )}>
                        {DEFAULT_SCAN_BEHAVIOR.excludePatterns.join(", ")}
                      </span>
                    </div>
                  </div>
                </div>

                {/* Persistence hint */}
                <div className="mt-4 px-2 py-2 border border-border">
                  <p className="text-xs text-muted-foreground">
                    Default: If no paths configured, CLI scans current directory (<span className="text-tui-cyan">./</span>)
                  </p>
                  <p className="text-xs text-muted-foreground mt-1">
                    Saved to <span className="text-tui-cyan">~/.rfz/config.yaml</span>
                  </p>
                </div>
              </div>
            )}

            {/* Component Registry Section (Read-Only) */}
            {currentSection.id === "registry" && (
              <div className="space-y-2">
                {/* Table header */}
                <div className="flex items-center gap-4 px-2 py-1 text-xs text-muted-foreground border-b border-border">
                  <span className="w-4"></span>
                  <span className="w-32">Name</span>
                  <span className="w-32">Artifact ID</span>
                  <span className="w-20">Category</span>
                  <span className="w-8 text-center">Order</span>
                  <span className="flex-1">Profiles</span>
                </div>
                
                {/* Component rows */}
                {COMPONENT_REGISTRY.map((component, index) => {
                  const isFocused = focusArea === "details" && index === registryIndex
                  
                  return (
                    <div
                      key={component.name}
                      className={cn(
                        "flex items-center gap-4 px-2 py-0.5 text-sm",
                        isFocused && "bg-tui-cyan text-background"
                      )}
                    >
                      <span className={cn(
                        "w-4",
                        isFocused ? "text-background" : "text-tui-cyan"
                      )}>
                        {isFocused ? ">" : " "}
                      </span>
                      <span className="w-32 truncate">{component.name}</span>
                      <span className={cn(
                        "w-32 truncate font-mono text-xs",
                        isFocused ? "text-background/80" : "text-muted-foreground"
                      )}>
                        {component.artifactId}
                      </span>
                      <span className={cn(
                        "w-20 text-xs",
                        isFocused ? "text-background" : 
                          component.category === "Core" ? "text-tui-green" :
                          component.category === "Simulator" ? "text-tui-yellow" : "text-tui-cyan"
                      )}>
                        {component.category}
                      </span>
                      <span className={cn(
                        "w-8 text-center font-mono",
                        isFocused ? "text-background" : "text-muted-foreground"
                      )}>
                        {component.buildOrder}
                      </span>
                      <span className={cn(
                        "flex-1 text-xs truncate",
                        isFocused ? "text-background/80" : "text-muted-foreground"
                      )}>
                        {component.defaultProfiles.join(", ")}
                      </span>
                    </div>
                  )
                })}
                
                {/* Read-only hint */}
                <div className="mt-4 px-2 py-2 border border-border">
                  <p className="text-xs text-muted-foreground">
                    Component metadata is defined by the CLI registry and not auto-discovered.
                  </p>
                  <p className="text-xs text-muted-foreground mt-1">
                    This view is <span className="text-tui-yellow">read-only</span>.
                  </p>
                </div>
              </div>
            )}

            {/* Detected Components Section (Runtime) */}
            {currentSection.id === "detected" && (
              <div className="space-y-2">
                {/* Table header */}
                <div className="flex items-center gap-4 px-2 py-1 text-xs text-muted-foreground border-b border-border">
                  <span className="w-4"></span>
                  <span className="w-32">Component</span>
                  <span className="w-20">Category</span>
                  <span className="w-16 text-center">Status</span>
                  <span className="w-16 text-center">Config</span>
                  <span className="flex-1">Path</span>
                </div>
                
                {/* Detected component rows */}
                {DETECTED_COMPONENTS.map((component, index) => {
                  const isFocused = focusArea === "details" && index === detectedIndex
                  
                  return (
                    <div
                      key={component.name}
                      className={cn(
                        "flex items-center gap-4 px-2 py-0.5 text-sm",
                        isFocused && "bg-tui-cyan text-background"
                      )}
                    >
                      <span className={cn(
                        "w-4",
                        isFocused ? "text-background" : "text-tui-cyan"
                      )}>
                        {isFocused ? ">" : " "}
                      </span>
                      <span className="w-32 truncate">{component.name}</span>
                      <span className={cn(
                        "w-20 text-xs",
                        isFocused ? "text-background" : 
                          component.category === "Core" ? "text-tui-green" :
                          component.category === "Simulator" ? "text-tui-yellow" : "text-tui-cyan"
                      )}>
                        {component.category}
                      </span>
                      <span className={cn(
                        "w-16 text-center",
                        isFocused ? "text-background" : component.found ? "text-tui-green" : "text-destructive"
                      )}>
                        {component.found ? "Found" : "Missing"}
                      </span>
                      <span className={cn(
                        "w-16 text-center",
                        isFocused ? "text-background" : component.hasConfigModule ? "text-tui-green" : "text-muted-foreground"
                      )}>
                        {component.hasConfigModule ? "Yes" : "No"}
                      </span>
                      <span className={cn(
                        "flex-1 font-mono text-xs truncate",
                        isFocused ? "text-background/80" : component.found ? "text-foreground" : "text-muted-foreground italic"
                      )}>
                        {component.path || "(not found)"}
                      </span>
                    </div>
                  )
                })}
                
                {/* Summary */}
                <div className="mt-4 px-2 py-2 border border-border">
                  <div className="flex items-center gap-4 text-sm">
                    <span>
                      <span className="text-tui-green">{DETECTED_COMPONENTS.filter(c => c.found).length}</span>
                      <span className="text-muted-foreground"> found</span>
                    </span>
                    <span className="text-border">|</span>
                    <span>
                      <span className="text-destructive">{DETECTED_COMPONENTS.filter(c => !c.found).length}</span>
                      <span className="text-muted-foreground"> missing</span>
                    </span>
                  </div>
                  <p className="text-xs text-muted-foreground mt-2">
                    Component locations are detected at runtime based on scan configuration.
                  </p>
                </div>
              </div>
            )}
          </div>
        </TuiBox>
      </div>
    </div>
  )
}
