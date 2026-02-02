"use client"

import { cn } from "@/lib/utils"
import type { ReactNode } from "react"
import { TuiBox } from "./tui-box"

interface TuiModalProps {
  children: ReactNode
  title: string
  isOpen: boolean
  className?: string
  width?: string
}

export function TuiModal({ 
  children, 
  title, 
  isOpen,
  className,
  width = "max-w-2xl"
}: TuiModalProps) {
  if (!isOpen) return null
  
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* Backdrop */}
      <div className="absolute inset-0 bg-background/80" />
      
      {/* Modal */}
      <div className={cn("relative z-10 w-full mx-4", width, className)}>
        <TuiBox 
          title={title} 
          borderStyle="double"
          focused
        >
          <div className="p-2">
            {children}
          </div>
        </TuiBox>
      </div>
    </div>
  )
}

interface TuiModalActionsProps {
  children: ReactNode
  className?: string
}

export function TuiModalActions({ children, className }: TuiModalActionsProps) {
  return (
    <div className={cn(
      "flex items-center justify-end gap-4 mt-4 pt-2 border-t border-border",
      className
    )}>
      {children}
    </div>
  )
}

interface TuiButtonProps {
  label: string
  shortcut?: string
  focused?: boolean
  variant?: "default" | "primary" | "danger"
  onClick?: () => void
  className?: string
}

export function TuiButton({ 
  label, 
  shortcut, 
  focused = false,
  variant = "default",
  onClick,
  className 
}: TuiButtonProps) {
  const variantStyles = {
    default: focused ? "bg-tui-cyan text-background" : "text-foreground",
    primary: focused ? "bg-tui-green text-background" : "text-tui-green",
    danger: focused ? "bg-destructive text-background" : "text-destructive"
  }
  
  return (
    <button
      type="button"
      onClick={onClick}
      className={cn(
        "inline-flex items-center gap-1 px-2 py-0.5 cursor-pointer transition-colors",
        "hover:bg-secondary",
        variantStyles[variant],
        className
      )}
    >
      {focused && <span>{">"}</span>}
      <span>[{label}]</span>
      {shortcut && (
        <span className={cn(
          "text-xs",
          focused ? "opacity-70" : "text-muted-foreground"
        )}>
          ({shortcut})
        </span>
      )}
    </button>
  )
}
