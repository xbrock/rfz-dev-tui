"use client"

import React from "react"

import { cn } from "@/lib/utils"
import type { ReactNode } from "react"

interface TuiBoxProps {
  children: ReactNode
  title?: string
  className?: string
  borderStyle?: "single" | "double" | "rounded" | "heavy"
  focused?: boolean
}

// Box drawing characters
const BORDERS = {
  single: {
    tl: "┌", tr: "┐", bl: "└", br: "┘",
    h: "─", v: "│",
    lt: "├", rt: "┤", tt: "┬", bt: "┴",
  },
  double: {
    tl: "╔", tr: "╗", bl: "╚", br: "╝",
    h: "═", v: "║",
    lt: "╠", rt: "╣", tt: "╦", bt: "╩",
  },
  rounded: {
    tl: "╭", tr: "╮", bl: "╰", br: "╯",
    h: "─", v: "│",
    lt: "├", rt: "┤", tt: "┬", bt: "┴",
  },
  heavy: {
    tl: "┏", tr: "┓", bl: "┗", br: "┛",
    h: "━", v: "┃",
    lt: "┣", rt: "┫", tt: "┳", bt: "┻",
  },
}

export function TuiBox({ 
  children, 
  title, 
  className,
  borderStyle = "single",
  focused = false
}: TuiBoxProps) {
  const b = BORDERS[borderStyle]
  
  return (
    <div className={cn(
      "relative",
      className
    )}>
      {/* Top border with optional title */}
      <div className={cn(
        "flex items-center whitespace-pre",
        focused ? "text-tui-cyan" : "text-border"
      )}>
        <span>{b.tl}</span>
        {title && (
          <>
            <span>{b.h}</span>
            <span className={cn(
              "px-1",
              focused ? "text-tui-cyan font-bold" : "text-foreground"
            )}>
              {title}
            </span>
          </>
        )}
        <span className="flex-1 overflow-hidden">
          {b.h.repeat(200)}
        </span>
        <span>{b.tr}</span>
      </div>
      
      {/* Content with side borders */}
      <div className="flex">
        <span className={cn(
          "whitespace-pre",
          focused ? "text-tui-cyan" : "text-border"
        )}>{b.v}</span>
        <div className="flex-1 min-w-0">
          {children}
        </div>
        <span className={cn(
          "whitespace-pre",
          focused ? "text-tui-cyan" : "text-border"
        )}>{b.v}</span>
      </div>
      
      {/* Bottom border */}
      <div className={cn(
        "flex whitespace-pre",
        focused ? "text-tui-cyan" : "text-border"
      )}>
        <span>{b.bl}</span>
        <span className="flex-1 overflow-hidden">
          {b.h.repeat(200)}
        </span>
        <span>{b.br}</span>
      </div>
    </div>
  )
}

export function TuiDivider({ 
  style = "single",
  className,
  label
}: { 
  style?: "single" | "double" | "dashed"
  className?: string
  label?: string
}) {
  const char = style === "double" ? "═" : style === "dashed" ? "╌" : "─"
  
  if (label) {
    return (
      <div className={cn("flex items-center text-border whitespace-pre overflow-hidden", className)}>
        <span>{char.repeat(3)}</span>
        <span className="px-2 text-muted-foreground text-xs">{label}</span>
        <span className="flex-1">{char.repeat(200)}</span>
      </div>
    )
  }
  
  return (
    <div className={cn("text-border whitespace-pre overflow-hidden", className)}>
      {char.repeat(200)}
    </div>
  )
}

// Tree connector characters (lipgloss-style)
export const TREE_CHARS = {
  branch: "├── ",
  lastBranch: "└── ",
  pipe: "│   ",
  space: "    ",
}

interface TuiTreeItemProps {
  label: string
  isLast?: boolean
  depth?: number
  icon?: React.ReactNode
  className?: string
  children?: React.ReactNode
}

export function TuiTreeItem({ 
  label, 
  isLast = false, 
  depth = 0, 
  icon,
  className,
  children 
}: TuiTreeItemProps) {
  const prefix = depth === 0 ? "" : isLast ? TREE_CHARS.lastBranch : TREE_CHARS.branch
  const indent = TREE_CHARS.space.repeat(Math.max(0, depth - 1))
  
  return (
    <div className={cn("font-mono text-sm", className)}>
      <div className="flex items-center">
        <span className="text-muted-foreground whitespace-pre">{indent}{prefix}</span>
        {icon && <span className="mr-1">{icon}</span>}
        <span>{label}</span>
      </div>
      {children}
    </div>
  )
}

// Status bar segment (lipgloss-style colored segments)
interface TuiStatusSegmentProps {
  label: string
  value?: string
  color?: "brand" | "success" | "warning" | "danger" | "muted"
  className?: string
}

export function TuiStatusSegment({ label, value, color = "muted", className }: TuiStatusSegmentProps) {
  const colorClasses = {
    brand: "bg-brand text-white",
    success: "bg-tui-green text-background",
    warning: "bg-tui-yellow text-background",
    danger: "bg-destructive text-white",
    muted: "bg-secondary text-foreground"
  }
  
  return (
    <span className={cn("px-2 py-0.5 text-sm font-mono", colorClasses[color], className)}>
      {label}{value && `: ${value}`}
    </span>
  )
}
