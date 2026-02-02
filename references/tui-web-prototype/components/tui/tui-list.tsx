"use client"

import { cn } from "@/lib/utils"

interface TuiListItem {
  id: string
  label: string
  description?: string
  category?: string
  disabled?: boolean
}

interface TuiListProps {
  items: TuiListItem[]
  selectedIndex: number
  selectedItems?: Set<string>
  multiSelect?: boolean
  showCategories?: boolean
  showTreeConnectors?: boolean
  className?: string
}

// Category colors (lipgloss-inspired)
const CATEGORY_COLORS: Record<string, string> = {
  Core: "text-tui-cyan",
  Simulation: "text-tui-yellow",
  Standalone: "text-tui-green",
}

export function TuiList({
  items,
  selectedIndex,
  selectedItems = new Set(),
  multiSelect = false,
  showCategories = false,
  showTreeConnectors = false,
  className
}: TuiListProps) {
  return (
    <div className={cn("flex flex-col font-mono", className)}>
      {items.map((item, index) => {
        const isFocused = index === selectedIndex
        const isSelected = selectedItems.has(item.id)
        const isLast = index === items.length - 1
        const treePrefix = showTreeConnectors ? (isLast ? "└── " : "├── ") : ""
        
        return (
          <div
            key={item.id}
            className={cn(
              "flex items-center gap-2 px-2 py-0.5 whitespace-nowrap text-sm",
              isFocused && "bg-tui-cyan text-background",
              item.disabled && "opacity-50"
            )}
          >
            {/* Tree connector or focus indicator */}
            {showTreeConnectors ? (
              <span className={cn(
                "w-6 whitespace-pre",
                isFocused ? "text-background/60" : "text-border"
              )}>
                {isFocused ? "> " : treePrefix.slice(0, 2)}
              </span>
            ) : (
              <span className={cn(
                "w-2",
                isFocused ? "text-background" : "text-tui-cyan"
              )}>
                {isFocused ? ">" : " "}
              </span>
            )}
            
            {/* Checkbox for multi-select */}
            {multiSelect && (
              <span className={cn(
                "w-6",
                isFocused ? "text-background" : isSelected ? "text-tui-green" : "text-muted-foreground"
              )}>
                {isSelected ? "[x]" : "[ ]"}
              </span>
            )}
            
            {/* Item label */}
            <span className="flex-1 min-w-0">{item.label}</span>
            
            {/* Category badge with color */}
            {showCategories && item.category && (
              <span className={cn(
                "text-xs px-1.5 py-0.5 rounded",
                isFocused 
                  ? "text-background/80 bg-background/20" 
                  : cn(CATEGORY_COLORS[item.category] || "text-muted-foreground", "bg-secondary/50")
              )}>
                {item.category}
              </span>
            )}
          </div>
        )
      })}
    </div>
  )
}

interface TuiNavItemProps {
  label: string
  index: number
  isFocused: boolean
  isActive?: boolean
  shortcut?: string
  onClick?: () => void
}

export function TuiNavItem({ label, index, isFocused, isActive, shortcut, onClick }: TuiNavItemProps) {
  return (
    <button
      type="button"
      onClick={onClick}
      className={cn(
        "w-full flex items-center gap-2 px-2 py-1 whitespace-nowrap text-left transition-colors",
        "hover:bg-secondary cursor-pointer",
        isFocused && "bg-tui-cyan text-background hover:bg-tui-cyan",
        !isFocused && isActive && "bg-secondary text-tui-cyan"
      )}
    >
      <span className={cn(
        "w-2",
        isFocused ? "text-background" : isActive ? "text-tui-cyan" : "text-tui-cyan"
      )}>
        {isFocused ? ">" : isActive ? ">" : " "}
      </span>
      <span className={cn(
        "w-4 text-center",
        isFocused ? "text-background" : isActive ? "text-tui-cyan" : "text-muted-foreground"
      )}>
        {index + 1}.
      </span>
      <span className="flex-1">{label}</span>
      {shortcut && (
        <span className={cn(
          "text-xs",
          isFocused ? "text-background/70" : "text-muted-foreground"
        )}>
          {shortcut}
        </span>
      )}
    </button>
  )
}
