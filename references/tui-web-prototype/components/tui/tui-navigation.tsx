"use client"

import { cn } from "@/lib/utils"
import { TuiBox } from "./tui-box"
import { TuiNavItem } from "./tui-list"

export type NavScreen = "welcome" | "build" | "logs" | "discover" | "config" | "exit"

interface NavItem {
  id: NavScreen
  label: string
  shortcut?: string
}

const NAV_ITEMS: NavItem[] = [
  { id: "build", label: "Build Components", shortcut: "1" },
  { id: "logs", label: "View Logs", shortcut: "2" },
  { id: "discover", label: "Discover", shortcut: "3" },
  { id: "config", label: "Configuration", shortcut: "4" },
  { id: "exit", label: "Exit", shortcut: "q" },
]

interface TuiNavigationProps {
  selectedIndex: number
  activeScreen: NavScreen
  onNavigate: (screen: NavScreen) => void
  className?: string
}

export function TuiNavigation({ selectedIndex, activeScreen, onNavigate, className }: TuiNavigationProps) {
  return (
    <div className={cn("w-60 min-w-60 max-w-60 flex-shrink-0 overflow-hidden flex flex-col", className)}>
      <TuiBox 
        title="Navigation" 
        borderStyle="rounded"
        focused={false}
      >
        <div className="py-1">
          {NAV_ITEMS.map((item, index) => (
            <TuiNavItem
              key={item.id}
              label={item.label}
              index={index}
              isFocused={index === selectedIndex}
              isActive={item.id === activeScreen}
              shortcut={item.shortcut}
              onClick={() => onNavigate(item.id)}
            />
          ))}
        </div>
        
        {/* Help footer with tree-style hints */}
        <div className="mt-4 px-2 py-1 border-t border-border">
          <div className="text-xs text-muted-foreground font-mono">
            <div className="flex"><span className="text-border">├── </span><span className="text-tui-cyan">↑/k</span><span className="ml-2">Up</span></div>
            <div className="flex"><span className="text-border">├── </span><span className="text-tui-cyan">↓/j</span><span className="ml-2">Down</span></div>
            <div className="flex"><span className="text-border">├── </span><span className="text-tui-cyan">Enter</span><span className="ml-2">Select</span></div>
            <div className="flex"><span className="text-border">└── </span><span className="text-tui-cyan">1-5</span><span className="ml-2">Quick nav</span></div>
          </div>
        </div>
      </TuiBox>
    </div>
  )
}

export { NAV_ITEMS }
export type { NavItem }
