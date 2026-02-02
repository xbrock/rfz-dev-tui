"use client"

import { cn } from "@/lib/utils"
import { TuiBox } from "../tui/tui-box"
import { TuiList } from "../tui/tui-list"
import { TuiButton } from "../tui/tui-modal"

export interface RfzComponent {
  id: string
  label: string
  category: "Core" | "Simulation" | "Standalone"
}

const RFZ_COMPONENTS: RfzComponent[] = [
  { id: "boss", label: "boss", category: "Core" },
  { id: "fistiv", label: "fistiv", category: "Core" },
  { id: "audiocon", label: "audiocon", category: "Standalone" },
  { id: "traktion", label: "traktion", category: "Core" },
  { id: "signalsteuerung", label: "signalsteuerung", category: "Core" },
  { id: "weichensteuerung", label: "weichensteuerung", category: "Core" },
  { id: "simkern", label: "simkern", category: "Simulation" },
  { id: "fahrdynamik", label: "fahrdynamik", category: "Simulation" },
  { id: "energierechnung", label: "energierechnung", category: "Simulation" },
  { id: "zuglauf", label: "zuglauf", category: "Simulation" },
  { id: "stellwerk", label: "stellwerk", category: "Core" },
  { id: "diagnose", label: "diagnose", category: "Standalone" },
  { id: "konfiguration", label: "konfiguration", category: "Standalone" },
]

interface BuildScreenProps {
  selectedIndex: number
  selectedComponents: Set<string>
  focusArea: "list" | "actions"
  actionIndex: number
  onOpenBuildConfig: () => void
  className?: string
}

export function BuildScreen({ 
  selectedIndex, 
  selectedComponents,
  focusArea,
  actionIndex,
  className 
}: BuildScreenProps) {
  const allSelected = selectedComponents.size === RFZ_COMPONENTS.length
  const noneSelected = selectedComponents.size === 0
  
  return (
    <div className={cn("flex flex-col h-full", className)}>
      {/* Header */}
      <TuiBox title="Build RFZ Components" borderStyle="rounded" focused={focusArea === "list"}>
        <div className="p-2">
          <div className="flex items-center justify-between mb-2">
            <span className="text-muted-foreground text-sm font-mono">
              Select components to build ({selectedComponents.size}/{RFZ_COMPONENTS.length} selected)
            </span>
            <div className="flex items-center gap-4 text-xs text-muted-foreground font-mono">
              <span><span className="text-tui-cyan">Space</span> Toggle</span>
              <span><span className="text-tui-cyan">a</span> All</span>
              <span><span className="text-tui-cyan">n</span> None</span>
            </div>
          </div>
          
          {/* Component list with subtle border */}
          <div className="border border-border/50 rounded">
            <TuiList
              items={RFZ_COMPONENTS.map(c => ({
                id: c.id,
                label: c.label,
                category: c.category
              }))}
              selectedIndex={focusArea === "list" ? selectedIndex : -1}
              selectedItems={selectedComponents}
              multiSelect
              showCategories
            />
          </div>
        </div>
      </TuiBox>
      
      {/* Actions bar */}
      <div className="mt-2">
        <TuiBox title="Actions" borderStyle="rounded" focused={focusArea === "actions"}>
          <div className="p-2 flex items-center gap-4">
            <TuiButton 
              label="Build Selected" 
              shortcut="Enter"
              focused={focusArea === "actions" && actionIndex === 0}
              variant="primary"
            />
            <TuiButton 
              label="Select All" 
              shortcut="a"
              focused={focusArea === "actions" && actionIndex === 1}
            />
            <TuiButton 
              label="Clear Selection" 
              shortcut="n"
              focused={focusArea === "actions" && actionIndex === 2}
            />
            
            <span className="ml-auto text-xs text-muted-foreground">
              <span className="text-tui-cyan">Tab</span> Switch focus
            </span>
          </div>
        </TuiBox>
      </div>
      
      {/* Quick legend with tree style */}
      <div className="mt-4 px-2 text-xs text-muted-foreground font-mono">
        <div className="flex gap-6">
          <span><span className="text-tui-green">[x]</span> Selected</span>
          <span><span className="text-muted-foreground">[ ]</span> Not selected</span>
          <span><span className="text-tui-cyan">{">"}</span> Current</span>
        </div>
      </div>
    </div>
  )
}

export { RFZ_COMPONENTS }
