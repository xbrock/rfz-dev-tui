"use client"

import React from "react"

import { cn } from "@/lib/utils"

interface TuiProgressProps {
  value: number // 0-100
  width?: number // character width
  showPercentage?: boolean
  className?: string
  label?: string
  variant?: "default" | "gradient" | "success" | "danger"
}

// Gradient colors for progress bar (lipgloss-inspired)
const GRADIENT_COLORS = [
  "#5A56E0", "#7B5AE4", "#9C5EE8", "#BD62EC", "#DE66F0", "#EE6FF8"
]

export function TuiProgress({
  value,
  width = 30,
  showPercentage = true,
  className,
  label,
  variant = "default"
}: TuiProgressProps) {
  const clampedValue = Math.max(0, Math.min(100, value))
  const filledWidth = Math.round((clampedValue / 100) * width)
  const emptyWidth = width - filledWidth
  
  const filledChar = "█"
  const emptyChar = "░"
  
  // Determine color class based on variant
  const getFilledColor = () => {
    if (variant === "success" || clampedValue === 100) return "text-tui-green"
    if (variant === "danger") return "text-destructive"
    return "text-tui-cyan"
  }
  
  // Render gradient progress
  const renderGradientFill = () => {
    if (variant !== "gradient" || filledWidth === 0) return null
    
    const segments: React.ReactNode[] = []
    const segmentSize = Math.ceil(filledWidth / GRADIENT_COLORS.length)
    
    for (let i = 0; i < filledWidth; i++) {
      const colorIndex = Math.min(Math.floor(i / segmentSize), GRADIENT_COLORS.length - 1)
      segments.push(
        <span key={i} style={{ color: GRADIENT_COLORS[colorIndex] }}>{filledChar}</span>
      )
    }
    return <>{segments}</>
  }
  
  return (
    <div className={cn("flex items-center gap-2 font-mono text-sm", className)}>
      {label && <span className="text-muted-foreground min-w-16">{label}</span>}
      {variant === "gradient" ? (
        <>
          {renderGradientFill()}
          <span className="text-muted-foreground/40">{emptyChar.repeat(emptyWidth)}</span>
        </>
      ) : (
        <>
          <span className={getFilledColor()}>{filledChar.repeat(filledWidth)}</span>
          <span className="text-muted-foreground/40">{emptyChar.repeat(emptyWidth)}</span>
        </>
      )}
      {showPercentage && (
        <span className={cn(
          "w-10 text-right tabular-nums",
          clampedValue === 100 ? "text-tui-green" : "text-muted-foreground"
        )}>{Math.round(clampedValue)}%</span>
      )}
    </div>
  )
}

// Braille-style progress bar (lipgloss-inspired)
interface TuiBrailleProgressProps {
  value: number
  width?: number
  className?: string
}

export function TuiBrailleProgress({ value, width = 20, className }: TuiBrailleProgressProps) {
  const clampedValue = Math.max(0, Math.min(100, value))
  const filledWidth = Math.round((clampedValue / 100) * width)
  const emptyWidth = width - filledWidth
  
  return (
    <span className={cn("font-mono text-sm", className)}>
      <span className={clampedValue === 100 ? "text-tui-green" : "text-tui-cyan"}>
        {"⣿".repeat(filledWidth)}
      </span>
      <span className="text-muted-foreground/40">{"⣀".repeat(emptyWidth)}</span>
    </span>
  )
}

interface TuiSpinnerProps {
  className?: string
  label?: string
  variant?: "dots" | "line" | "arc" | "bounce"
}

const SPINNER_FRAMES = {
  dots: ["⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"],
  line: ["|", "/", "-", "\\"],
  arc: ["◴", "◷", "◶", "◵"],
  bounce: ["▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂"]
}

export function TuiSpinner({ className, label, variant = "dots" }: TuiSpinnerProps) {
  return (
    <span className={cn("inline-flex items-center gap-2", className)}>
      <span className="tui-spinner text-tui-yellow" data-variant={variant} />
      {label && <span>{label}</span>}
    </span>
  )
}

interface TuiStatusProps {
  status: "pending" | "running" | "success" | "failed" | "error"
  label?: string
  className?: string
}

const STATUS_ICONS = {
  pending: "○",
  running: "◐",
  success: "✓",
  failed: "✗",
  error: "!"
}

const STATUS_COLORS = {
  pending: "text-muted-foreground",
  running: "text-tui-yellow",
  success: "text-tui-green",
  failed: "text-destructive",
  error: "text-destructive"
}

export function TuiStatus({ status, label, className }: TuiStatusProps) {
  return (
    <span className={cn("inline-flex items-center gap-2", STATUS_COLORS[status], className)}>
      {status === "running" ? (
        <TuiSpinner />
      ) : (
        <span>{STATUS_ICONS[status]}</span>
      )}
      {label && <span>{label}</span>}
    </span>
  )
}
