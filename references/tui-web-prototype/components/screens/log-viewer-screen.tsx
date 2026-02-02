"use client"

import { useState, useEffect, useRef } from "react"
import { cn } from "@/lib/utils"
import { TuiBox } from "../tui/tui-box"

export interface LogComponent {
  id: string
  name: string
  category: "Core" | "Standalone" | "Simulation"
  isRunning: boolean
  logPath: string
}

export const LOG_COMPONENTS: LogComponent[] = [
  { id: "rfz-core", name: "rfz-core", category: "Core", isRunning: true, logPath: "log/rfz/rfz-core.log" },
  { id: "rfz-api", name: "rfz-api", category: "Core", isRunning: true, logPath: "log/rfz/rfz-api.log" },
  { id: "rfz-dispatcher", name: "rfz-dispatcher", category: "Core", isRunning: true, logPath: "log/rfz/rfz-dispatcher.log" },
  { id: "rfz-traktion", name: "rfz-traktion", category: "Core", isRunning: false, logPath: "log/rfz/rfz-traktion.log" },
  { id: "rfz-fahrzeug", name: "rfz-fahrzeug", category: "Standalone", isRunning: true, logPath: "log/rfz/rfz-fahrzeug.log" },
  { id: "rfz-simulation", name: "rfz-simulation", category: "Simulation", isRunning: false, logPath: "log/rfz/rfz-simulation.log" },
  { id: "rfz-mock-server", name: "rfz-mock-server", category: "Simulation", isRunning: true, logPath: "log/rfz/rfz-mock-server.log" },
  { id: "rfz-db-connector", name: "rfz-db-connector", category: "Core", isRunning: true, logPath: "log/rfz/rfz-db-connector.log" },
]

type LogLevel = "INFO" | "WARN" | "ERROR" | "DEBUG"

interface LogEntry {
  timestamp: string
  level: LogLevel
  message: string
}

// Sample log entries for simulation
const generateSampleLogs = (): LogEntry[] => {
  const messages: { level: LogLevel; message: string }[] = [
    { level: "INFO", message: "Service started successfully" },
    { level: "INFO", message: "Listening on port 11090" },
    { level: "DEBUG", message: "Processing incoming request..." },
    { level: "INFO", message: "Connection established to database" },
    { level: "WARN", message: "High memory usage detected (85%)" },
    { level: "INFO", message: "Request processed in 45ms" },
    { level: "DEBUG", message: "Cache hit for key: session_abc123" },
    { level: "INFO", message: "New client connected: 192.168.1.42" },
    { level: "ERROR", message: "Failed to connect to remote service" },
    { level: "INFO", message: "Retrying connection in 5 seconds..." },
    { level: "INFO", message: "Connection restored" },
    { level: "WARN", message: "Deprecated API endpoint called: /v1/status" },
    { level: "DEBUG", message: "Serializing response payload..." },
    { level: "INFO", message: "Request completed successfully" },
    { level: "INFO", message: "Heartbeat sent to coordinator" },
    { level: "DEBUG", message: "Thread pool size: 8/16 active" },
    { level: "INFO", message: "Configuration reloaded" },
    { level: "WARN", message: "Connection timeout, retrying..." },
    { level: "INFO", message: "Checkpoint saved to disk" },
    { level: "ERROR", message: "Invalid message format received" },
    { level: "INFO", message: "Message discarded, awaiting retry" },
    { level: "DEBUG", message: "GC completed in 12ms" },
    { level: "INFO", message: "Metrics published to collector" },
  ]
  
  const logs: LogEntry[] = []
  const now = new Date()
  
  for (let i = 0; i < 50; i++) {
    const entry = messages[i % messages.length]
    const time = new Date(now.getTime() - (50 - i) * 1000)
    logs.push({
      timestamp: time.toISOString().substring(11, 23),
      level: entry.level,
      message: entry.message
    })
  }
  
  return logs
}

interface LogViewerScreenProps {
  selectedIndex: number
  focusArea: "components" | "logs" | "filters"
  filterLevel: LogLevel | "ALL"
  followMode: boolean
  logScrollPosition: number
  searchQuery: string
  className?: string
}

export function LogViewerScreen({
  selectedIndex,
  focusArea,
  filterLevel,
  followMode,
  logScrollPosition,
  searchQuery,
}: LogViewerScreenProps) {
  const [logs, setLogs] = useState<LogEntry[]>(generateSampleLogs())
  const [liveIndicator, setLiveIndicator] = useState(true)
  const logContainerRef = useRef<HTMLDivElement>(null)
  
  // Simulate live log updates
  useEffect(() => {
    const interval = setInterval(() => {
      setLiveIndicator(prev => !prev)
      
      // Add new log entry occasionally
      if (Math.random() < 0.3) {
        const levels: LogLevel[] = ["INFO", "INFO", "INFO", "DEBUG", "WARN", "ERROR"]
        const messages = [
          "Processing request...",
          "Health check passed",
          "Data synchronized",
          "Cache invalidated",
          "Connection pool refreshed",
        ]
        
        setLogs(prev => {
          const newLogs = [...prev]
          if (newLogs.length > 100) {
            newLogs.shift()
          }
          newLogs.push({
            timestamp: new Date().toISOString().substring(11, 23),
            level: levels[Math.floor(Math.random() * levels.length)],
            message: messages[Math.floor(Math.random() * messages.length)]
          })
          return newLogs
        })
      }
    }, 1000)
    
    return () => clearInterval(interval)
  }, [])
  
  const selectedComponent = LOG_COMPONENTS[selectedIndex]
  
  // Filter logs
  const filteredLogs = logs.filter(log => {
    if (filterLevel !== "ALL" && log.level !== filterLevel) return false
    if (searchQuery && !log.message.toLowerCase().includes(searchQuery.toLowerCase())) return false
    return true
  })
  
  // Calculate visible logs based on scroll position
  const visibleLogsCount = 20
  const startIndex = Math.max(0, followMode ? filteredLogs.length - visibleLogsCount : logScrollPosition)
  const visibleLogs = filteredLogs.slice(startIndex, startIndex + visibleLogsCount)
  
  const getLevelColor = (level: LogLevel) => {
    switch (level) {
      case "INFO": return "text-tui-cyan"
      case "WARN": return "text-tui-yellow"
      case "ERROR": return "text-destructive"
      case "DEBUG": return "text-muted-foreground"
    }
  }
  
  return (
    <div className="h-full flex gap-2">
      {/* Left Panel - Component List */}
      <div className="w-64 flex-shrink-0 flex flex-col gap-2">
        <TuiBox 
          title="Components" 
          borderStyle="single" 
          focused={focusArea === "components"}
          className="flex-1"
        >
          <div className="py-1">
            {LOG_COMPONENTS.map((component, index) => {
              const isFocused = focusArea === "components" && index === selectedIndex
              
              return (
                <div
                  key={component.id}
                  className={cn(
                    "flex items-center gap-2 px-2 py-0.5 text-sm",
                    isFocused && "bg-tui-cyan text-background"
                  )}
                >
                  {/* Focus indicator */}
                  <span className={cn(
                    "w-2 flex-shrink-0",
                    isFocused ? "text-background" : "text-tui-cyan"
                  )}>
                    {isFocused ? ">" : " "}
                  </span>
                  
                  {/* Running indicator */}
                  <span className={cn(
                    "w-2 flex-shrink-0",
                    isFocused ? "text-background" : component.isRunning ? "text-tui-green" : "text-muted-foreground"
                  )}>
                    {component.isRunning ? "●" : "○"}
                  </span>
                  
                  {/* Component name */}
                  <span className="flex-1 truncate">{component.name}</span>
                  
                  {/* Category badge */}
                  <span className={cn(
                    "text-xs flex-shrink-0",
                    isFocused ? "text-background/70" : "text-muted-foreground"
                  )}>
                    {component.category.substring(0, 4)}
                  </span>
                </div>
              )
            })}
          </div>
          
          {/* Legend */}
          <div className="mt-2 px-2 py-1 border-t border-border text-xs text-muted-foreground">
            <span className="text-tui-green">●</span> Running{" "}
            <span className="text-muted-foreground ml-2">○</span> Stopped
          </div>
        </TuiBox>
        
        {/* Status */}
        <TuiBox title="Status" borderStyle="single">
          <div className="p-2 text-xs space-y-1">
            <p>
              <span className="text-muted-foreground">Component:</span>{" "}
              <span className="text-foreground">{selectedComponent?.name}</span>
            </p>
            <p>
              <span className="text-muted-foreground">Log file:</span>{" "}
              <span className="text-foreground truncate block">{selectedComponent?.logPath}</span>
            </p>
            <p>
              <span className="text-muted-foreground">Follow:</span>{" "}
              <span className={followMode ? "text-tui-green" : "text-muted-foreground"}>
                {followMode ? "ON" : "OFF"}
              </span>
            </p>
          </div>
        </TuiBox>
      </div>
      
      {/* Right Panel - Log View */}
      <div className="flex-1 flex flex-col gap-2 min-w-0">
        {/* Live Log View */}
        <TuiBox 
          title={
            <span className="flex items-center gap-2">
              Live Logs - {selectedComponent?.name}
              <span className={cn(
                "text-xs px-1",
                liveIndicator ? "bg-tui-green text-background" : "bg-tui-green/50 text-background"
              )}>
                LIVE
              </span>
            </span>
          }
          borderStyle="single" 
          focused={focusArea === "logs"}
          className="flex-1"
        >
          <div ref={logContainerRef} className="p-2 h-full overflow-hidden font-mono text-xs">
            {/* Log entries */}
            <div className="space-y-0.5">
              {visibleLogs.map((log, index) => (
                <div key={`${log.timestamp}-${index}`} className="flex gap-2">
                  <span className="text-muted-foreground flex-shrink-0">{log.timestamp}</span>
                  <span className={cn("w-12 flex-shrink-0", getLevelColor(log.level))}>
                    [{log.level.padEnd(5)}]
                  </span>
                  <span className="text-foreground truncate">{log.message}</span>
                </div>
              ))}
            </div>
            
            {/* Streaming indicator */}
            {followMode && (
              <div className="mt-2 flex items-center gap-2 text-muted-foreground">
                <span className="tui-spinner"></span>
                <span>Waiting for new log entries...</span>
              </div>
            )}
          </div>
        </TuiBox>
        
        {/* Filter Bar */}
        <TuiBox 
          title="Filters & Navigation" 
          borderStyle="single"
          focused={focusArea === "filters"}
        >
          <div className="p-2 flex flex-wrap items-center gap-4 text-xs">
            {/* Log Level Filters */}
            <div className="flex items-center gap-2">
              <span className="text-muted-foreground">Level:</span>
              {(["ALL", "INFO", "WARN", "ERROR", "DEBUG"] as const).map(level => (
                <span
                  key={level}
                  className={cn(
                    "px-1",
                    filterLevel === level 
                      ? "bg-tui-cyan text-background" 
                      : level === "ERROR" 
                        ? "text-destructive"
                        : level === "WARN"
                          ? "text-tui-yellow"
                          : "text-muted-foreground"
                  )}
                >
                  [{level}]
                </span>
              ))}
            </div>
            
            {/* Search */}
            <div className="flex items-center gap-2">
              <span className="text-muted-foreground">/</span>
              <span className={cn(
                "px-1 min-w-20",
                searchQuery ? "text-foreground" : "text-muted-foreground"
              )}>
                {searchQuery || "Search..."}
              </span>
            </div>
            
            {/* Time navigation */}
            <div className="flex items-center gap-2 text-muted-foreground">
              <span>Time:</span>
              <span className="text-tui-cyan">[{"<"}-1m]</span>
              <span className="text-tui-cyan">[{"<"}-5m]</span>
              <span className="text-tui-cyan">[+1m-{">"}]</span>
            </div>
            
            {/* Follow toggle */}
            <span className={cn(
              "px-1",
              followMode ? "bg-tui-green text-background" : "text-muted-foreground"
            )}>
              [Follow: {followMode ? "ON" : "OFF"}]
            </span>
          </div>
          
          {/* Keyboard hints */}
          <div className="px-2 pb-2 flex flex-wrap gap-4 text-xs text-muted-foreground border-t border-border pt-2">
            <span><span className="text-tui-cyan">↑↓</span> scroll</span>
            <span><span className="text-tui-cyan">/</span> search</span>
            <span><span className="text-tui-cyan">f</span> filter</span>
            <span><span className="text-tui-cyan">←→</span> time jump</span>
            <span><span className="text-tui-cyan">Space</span> follow</span>
            <span><span className="text-tui-cyan">Tab</span> switch panel</span>
          </div>
        </TuiBox>
      </div>
    </div>
  )
}

export { generateSampleLogs }
