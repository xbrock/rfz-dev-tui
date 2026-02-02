"use client"

import { useEffect, useState, useCallback, useRef } from "react"
import { TuiNavigation, NAV_ITEMS, type NavScreen } from "@/components/tui/tui-navigation"
import { TuiBox } from "@/components/tui/tui-box"
import { WelcomeScreen } from "@/components/screens/welcome-screen"
import { BuildScreen, RFZ_COMPONENTS } from "@/components/screens/build-screen"
import { BuildConfigModal, MAVEN_GOALS, MAVEN_PROFILES, PORT_OPTIONS, type BuildConfig } from "@/components/screens/build-config-modal"
import { BuildExecutionView, type ComponentBuildState, type BuildPhase, type BuildStatus } from "@/components/screens/build-execution-view"
import { DiscoverScreen } from "@/components/screens/discover-screen"
import { ConfigScreen } from "@/components/screens/config-screen"
import { LogsModal } from "@/components/screens/logs-modal"
import { LogViewerScreen, LOG_COMPONENTS } from "@/components/screens/log-viewer-screen"

type AppView = "main" | "building" | "config-modal" | "logs-modal"

export default function CLRfZCLI() {
  // Navigation state
  const [navIndex, setNavIndex] = useState(0)
  const [currentScreen, setCurrentScreen] = useState<NavScreen>("welcome")
  const [appView, setAppView] = useState<AppView>("main")
  
  // Build screen state
  const [buildListIndex, setBuildListIndex] = useState(0)
  const [selectedComponents, setSelectedComponents] = useState<Set<string>>(new Set())
  const [buildFocusArea, setBuildFocusArea] = useState<"list" | "actions">("list")
  const [buildActionIndex, setBuildActionIndex] = useState(0)
  
  // Build config modal state
  const [buildConfig, setBuildConfig] = useState<BuildConfig>({
    mavenGoal: "clean-install",
    profiles: new Set(["target_env_dev"]),
    port: "11090",
    skipTests: true
  })
  const [configFocusArea, setConfigFocusArea] = useState<"goals" | "profiles" | "port" | "skipTests" | "actions">("goals")
  const [configGoalIndex, setConfigGoalIndex] = useState(3) // clean-install
  const [configProfileIndex, setConfigProfileIndex] = useState(0)
  const [configPortIndex, setConfigPortIndex] = useState(0) // 11090 default
  const [configActionIndex, setConfigActionIndex] = useState(1) // Start Build
  
  // Build execution state
  const [builds, setBuilds] = useState<ComponentBuildState[]>([])
  const [execFocusArea, setExecFocusArea] = useState<"list" | "actions">("list")
  const [execListIndex, setExecListIndex] = useState(0)
  const [execActionIndex, setExecActionIndex] = useState(0)
  const [buildComplete, setBuildComplete] = useState(false)
  
  // Discover screen state
  const [discoverListIndex, setDiscoverListIndex] = useState(0)
  const [discoverFocusArea, setDiscoverFocusArea] = useState<"list" | "actions">("list")
  const [discoverActionIndex, setDiscoverActionIndex] = useState(0)
  
  // Log viewer screen state
  const [logViewerIndex, setLogViewerIndex] = useState(0)
  const [logViewerFocusArea, setLogViewerFocusArea] = useState<"components" | "logs" | "filters">("components")
  const [logFilterLevel, setLogFilterLevel] = useState<"ALL" | "INFO" | "WARN" | "ERROR" | "DEBUG">("ALL")
  const [logFollowMode, setLogFollowMode] = useState(true)
  const [logScrollPosition, setLogScrollPosition] = useState(0)
  const [logSearchQuery, setLogSearchQuery] = useState("")
  
  // Logs modal state
  const [logsScrollPosition, setLogsScrollPosition] = useState(0)
  const [selectedBuildForLogs, setSelectedBuildForLogs] = useState<ComponentBuildState | null>(null)
  const [logsErrorOnly, setLogsErrorOnly] = useState(false)
  
  // Configuration screen state
  const [configSectionIndex, setConfigSectionIndex] = useState(0)
  const [configScreenFocusArea, setConfigScreenFocusArea] = useState<"sections" | "details">("sections")
  const [configDetailIndex, setConfigDetailIndex] = useState(0)
  const [configEditMode, setConfigEditMode] = useState(false)
  const [configEditValue, setConfigEditValue] = useState("")
  const [configPathIndex, setConfigPathIndex] = useState(0)
  const [configBehaviorIndex, setConfigBehaviorIndex] = useState(0)
  const [configRegistryIndex, setConfigRegistryIndex] = useState(0)
  const [configDetectedIndex, setConfigDetectedIndex] = useState(0)
  
  // Focus management
  const [focusArea, setFocusArea] = useState<"nav" | "content">("nav")
  
  // Build simulation timer
  const buildTimerRef = useRef<NodeJS.Timeout | null>(null)
  
  // Handle navigation click
  const handleNavigate = useCallback((screen: NavScreen) => {
    if (screen === "exit") {
      // Could show exit confirmation
      return
    }
    setCurrentScreen(screen)
    // Find the index for this screen
    const idx = NAV_ITEMS.findIndex(item => item.id === screen)
    if (idx !== -1) {
      setNavIndex(idx)
    }
    if (screen !== "welcome") {
      setFocusArea("content")
    }
    // Reset log viewer state when navigating to it
    if (screen === "logs") {
      setLogViewerFocusArea("components")
    }
    // Reset config screen state when navigating to it
    if (screen === "config") {
      setConfigScreenFocusArea("sections")
    }
  }, [])
  
  // Simulate build progress
  const simulateBuild = useCallback(() => {
    if (buildTimerRef.current) {
      clearInterval(buildTimerRef.current)
    }
    
    buildTimerRef.current = setInterval(() => {
      setBuilds(prevBuilds => {
        const newBuilds = [...prevBuilds]
        let hasRunning = false
        let hasUpdate = false
        
        for (let i = 0; i < newBuilds.length; i++) {
          const build = newBuilds[i]
          
          if (build.status === "running") {
            hasRunning = true
            hasUpdate = true
            newBuilds[i] = {
              ...build,
              elapsedTime: build.elapsedTime + 1
            }
            
            // Progress through phases
            const phases: BuildPhase[] = ["Compiling", "Testing", "Packaging", "Installing", "Done"]
            const currentPhaseIndex = phases.indexOf(build.phase as BuildPhase)
            
            // Random chance to advance phase or complete
            if (Math.random() < 0.15 && currentPhaseIndex < phases.length - 1) {
              if (currentPhaseIndex === phases.length - 2) {
                // Random failure chance
                if (Math.random() < 0.1) {
                  newBuilds[i] = {
                    ...newBuilds[i],
                    phase: "Done",
                    status: "failed"
                  }
                } else {
                  newBuilds[i] = {
                    ...newBuilds[i],
                    phase: "Done",
                    status: "success"
                  }
                }
              } else {
                newBuilds[i] = {
                  ...newBuilds[i],
                  phase: phases[currentPhaseIndex + 1]
                }
              }
            }
          } else if (build.status === "pending") {
            // Check if previous builds are done
            const allPreviousDone = newBuilds
              .slice(0, i)
              .every(b => b.status === "success" || b.status === "failed" || b.status === "error")
            
            // Start next build with some probability
            if (allPreviousDone || (i > 0 && Math.random() < 0.3)) {
              hasUpdate = true
              newBuilds[i] = {
                ...build,
                status: "running",
                phase: "Compiling"
              }
            }
          }
        }
        
        // Check if build is complete
        const allDone = newBuilds.every(
          b => b.status === "success" || b.status === "failed" || b.status === "error"
        )
        
        if (allDone && !hasRunning) {
          setBuildComplete(true)
          if (buildTimerRef.current) {
            clearInterval(buildTimerRef.current)
            buildTimerRef.current = null
          }
        }
        
        return hasUpdate ? newBuilds : prevBuilds
      })
    }, 500)
  }, [])
  
  // Start build
  const startBuild = useCallback(() => {
    const componentsList = Array.from(selectedComponents)
    const initialBuilds: ComponentBuildState[] = componentsList.map((id, index) => ({
      id,
      name: id,
      phase: "Pending" as BuildPhase,
      status: index === 0 ? "running" : "pending" as BuildStatus,
      elapsedTime: 0,
      logs: []
    }))
    
    // Start first build
    if (initialBuilds.length > 0) {
      initialBuilds[0].phase = "Compiling"
    }
    
    setBuilds(initialBuilds)
    setBuildComplete(false)
    setAppView("building")
    setExecFocusArea("list")
    setExecListIndex(0)
    setExecActionIndex(0)
    
    simulateBuild()
  }, [selectedComponents, simulateBuild])
  
  // Cleanup timer on unmount
  useEffect(() => {
    return () => {
      if (buildTimerRef.current) {
        clearInterval(buildTimerRef.current)
      }
    }
  }, [])
  
  // Keyboard handler
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      const key = e.key.toLowerCase()
      
      // Global escape handling
      if (e.key === "Escape") {
        if (appView === "config-modal") {
          setAppView("main")
          return
        }
        if (appView === "logs-modal") {
          setAppView("building")
          return
        }
        if (appView === "building" && buildComplete) {
          setAppView("main")
          return
        }
        if (focusArea === "content") {
          setFocusArea("nav")
          return
        }
      }
      
      // Logs modal navigation
      if (appView === "logs-modal") {
        if (key === "j" || key === "arrowdown") {
          setLogsScrollPosition(prev => Math.min(prev + 1, 50))
        } else if (key === "k" || key === "arrowup") {
          setLogsScrollPosition(prev => Math.max(prev - 1, 0))
        } else if (key === "e") {
          setLogsErrorOnly(prev => !prev)
          setLogsScrollPosition(0)
        }
        return
      }
      
      // Config modal navigation
      if (appView === "config-modal") {
        if (key === "tab") {
          e.preventDefault()
          setConfigFocusArea(prev => {
            if (prev === "goals") return "profiles"
            if (prev === "profiles") return "port"
            if (prev === "port") return "skipTests"
            if (prev === "skipTests") return "actions"
            return "goals"
          })
        } else if (configFocusArea === "goals") {
          if (key === "h" || key === "arrowleft") {
            setConfigGoalIndex(prev => Math.max(prev - 1, 0))
          } else if (key === "l" || key === "arrowright") {
            setConfigGoalIndex(prev => Math.min(prev + 1, MAVEN_GOALS.length - 1))
          } else if (key === " " || key === "enter") {
            setBuildConfig(prev => ({
              ...prev,
              mavenGoal: MAVEN_GOALS[configGoalIndex].id
            }))
          }
        } else if (configFocusArea === "profiles") {
          if (key === "j" || key === "arrowdown") {
            setConfigProfileIndex(prev => Math.min(prev + 1, MAVEN_PROFILES.length - 1))
          } else if (key === "k" || key === "arrowup") {
            setConfigProfileIndex(prev => Math.max(prev - 1, 0))
          } else if (key === " ") {
            e.preventDefault()
            const profileId = MAVEN_PROFILES[configProfileIndex].id
            setBuildConfig(prev => {
              const newProfiles = new Set(prev.profiles)
              if (newProfiles.has(profileId)) {
                newProfiles.delete(profileId)
              } else {
                newProfiles.add(profileId)
              }
              return { ...prev, profiles: newProfiles }
            })
          }
        } else if (configFocusArea === "port") {
          if (key === "h" || key === "arrowleft") {
            setConfigPortIndex(prev => Math.max(prev - 1, 0))
          } else if (key === "l" || key === "arrowright") {
            setConfigPortIndex(prev => Math.min(prev + 1, PORT_OPTIONS.length - 1))
          } else if (key === " " || key === "enter") {
            setBuildConfig(prev => ({
              ...prev,
              port: PORT_OPTIONS[configPortIndex].id
            }))
          }
        } else if (configFocusArea === "skipTests") {
          if (key === " " || key === "enter") {
            e.preventDefault()
            setBuildConfig(prev => ({
              ...prev,
              skipTests: !prev.skipTests
            }))
          }
        } else if (configFocusArea === "actions") {
          if (key === "h" || key === "arrowleft") {
            setConfigActionIndex(prev => Math.max(prev - 1, 0))
          } else if (key === "l" || key === "arrowright") {
            setConfigActionIndex(prev => Math.min(prev + 1, 1))
          } else if (key === "enter") {
            if (configActionIndex === 0) {
              setAppView("main")
            } else {
              startBuild()
            }
          }
        }
        return
      }
      
      // Build execution navigation
      if (appView === "building") {
        if (key === "tab") {
          e.preventDefault()
          setExecFocusArea(prev => prev === "list" ? "actions" : "list")
        } else if (execFocusArea === "list") {
          if (key === "j" || key === "arrowdown") {
            setExecListIndex(prev => Math.min(prev + 1, builds.length - 1))
          } else if (key === "k" || key === "arrowup") {
            setExecListIndex(prev => Math.max(prev - 1, 0))
          }
        } else if (execFocusArea === "actions") {
          if (key === "h" || key === "arrowleft") {
            setExecActionIndex(prev => Math.max(prev - 1, 0))
          } else if (key === "arrowright") {
            setExecActionIndex(prev => Math.min(prev + 1, buildComplete ? 2 : 1))
          } else if (key === "l") {
            // "l" is shortcut for View Logs
            setSelectedBuildForLogs(builds[execListIndex])
            setLogsScrollPosition(0)
            setAppView("logs-modal")
          } else if (key === "enter") {
            if (execActionIndex === 0) {
              // View logs
              setSelectedBuildForLogs(builds[execListIndex])
              setLogsScrollPosition(0)
              setAppView("logs-modal")
            } else if (execActionIndex === 1 && buildComplete) {
              // Rebuild failed - go back to build screen
              setAppView("main")
            } else if (execActionIndex === 2 && buildComplete) {
              // Back to selection
              setAppView("main")
            } else if (execActionIndex === 1 && !buildComplete) {
              // Cancel build
              setAppView("main")
            }
          }
        }
        // Also allow "l" globally during build to view logs
        if (key === "l" && execFocusArea === "list") {
          setSelectedBuildForLogs(builds[execListIndex])
          setLogsScrollPosition(0)
          setAppView("logs-modal")
        }
        return
      }
      
      // Navigation focus
      if (focusArea === "nav") {
        if (key === "j" || key === "arrowdown") {
          setNavIndex(prev => Math.min(prev + 1, NAV_ITEMS.length - 1))
        } else if (key === "k" || key === "arrowup") {
          setNavIndex(prev => Math.max(prev - 1, 0))
        } else if (key === "enter" || key === "l" || key === "arrowright") {
          const screen = NAV_ITEMS[navIndex].id
          if (screen === "exit") {
            // Could show exit confirmation
          } else {
            setCurrentScreen(screen)
            if (screen !== "welcome") {
              setFocusArea("content")
            }
            if (screen === "config") {
              setConfigScreenFocusArea("sections")
            }
          }
        } else if (key === "1" || key === "2" || key === "3" || key === "4" || key === "5") {
          const idx = parseInt(key) - 1
          if (idx < NAV_ITEMS.length) {
            setNavIndex(idx)
            const screen = NAV_ITEMS[idx].id
            if (screen !== "exit") {
              setCurrentScreen(screen)
              if (screen !== "welcome") {
                setFocusArea("content")
              }
              if (screen === "config") {
                setConfigScreenFocusArea("sections")
              }
            }
          }
        } else if (key === "q") {
          setNavIndex(NAV_ITEMS.length - 1)
        }
        return
      }
      
      // Content area - Build screen
      if (currentScreen === "build" && focusArea === "content") {
        if (key === "tab") {
          e.preventDefault()
          setBuildFocusArea(prev => prev === "list" ? "actions" : "list")
        } else if (buildFocusArea === "list") {
          if (key === "j" || key === "arrowdown") {
            setBuildListIndex(prev => Math.min(prev + 1, RFZ_COMPONENTS.length - 1))
          } else if (key === "k" || key === "arrowup") {
            setBuildListIndex(prev => Math.max(prev - 1, 0))
          } else if (key === " ") {
            e.preventDefault()
            const componentId = RFZ_COMPONENTS[buildListIndex].id
            setSelectedComponents(prev => {
              const newSet = new Set(prev)
              if (newSet.has(componentId)) {
                newSet.delete(componentId)
              } else {
                newSet.add(componentId)
              }
              return newSet
            })
          } else if (key === "a") {
            setSelectedComponents(new Set(RFZ_COMPONENTS.map(c => c.id)))
          } else if (key === "n") {
            setSelectedComponents(new Set())
          } else if (key === "enter") {
            if (selectedComponents.size > 0) {
              setAppView("config-modal")
              setConfigFocusArea("goals")
            }
          }
        } else if (buildFocusArea === "actions") {
          if (key === "h" || key === "arrowleft") {
            setBuildActionIndex(prev => Math.max(prev - 1, 0))
          } else if (key === "l" || key === "arrowright") {
            setBuildActionIndex(prev => Math.min(prev + 1, 2))
          } else if (key === "enter") {
            if (buildActionIndex === 0 && selectedComponents.size > 0) {
              setAppView("config-modal")
              setConfigFocusArea("goals")
            } else if (buildActionIndex === 1) {
              setSelectedComponents(new Set(RFZ_COMPONENTS.map(c => c.id)))
            } else if (buildActionIndex === 2) {
              setSelectedComponents(new Set())
            }
          }
        }
      }
      
      // Content area - Discover screen
      if (currentScreen === "discover" && focusArea === "content") {
        if (key === "tab") {
          e.preventDefault()
          setDiscoverFocusArea(prev => prev === "list" ? "actions" : "list")
        } else if (discoverFocusArea === "list") {
          if (key === "j" || key === "arrowdown") {
            setDiscoverListIndex(prev => Math.min(prev + 1, 9))
          } else if (key === "k" || key === "arrowup") {
            setDiscoverListIndex(prev => Math.max(prev - 1, 0))
          }
        } else if (discoverFocusArea === "actions") {
          if (key === "h" || key === "arrowleft") {
            setDiscoverActionIndex(prev => Math.max(prev - 1, 0))
          } else if (key === "l" || key === "arrowright") {
            setDiscoverActionIndex(prev => Math.min(prev + 1, 3))
          }
        }
      }
      
      // Content area - Log Viewer screen
      if (currentScreen === "logs" && focusArea === "content") {
        if (key === "tab") {
          e.preventDefault()
          setLogViewerFocusArea(prev => {
            if (prev === "components") return "logs"
            if (prev === "logs") return "filters"
            return "components"
          })
        } else if (logViewerFocusArea === "components") {
          if (key === "j" || key === "arrowdown") {
            setLogViewerIndex(prev => Math.min(prev + 1, LOG_COMPONENTS.length - 1))
          } else if (key === "k" || key === "arrowup") {
            setLogViewerIndex(prev => Math.max(prev - 1, 0))
          }
        } else if (logViewerFocusArea === "logs") {
          if (key === "j" || key === "arrowdown") {
            setLogFollowMode(false)
            setLogScrollPosition(prev => Math.min(prev + 1, 50))
          } else if (key === "k" || key === "arrowup") {
            setLogFollowMode(false)
            setLogScrollPosition(prev => Math.max(prev - 1, 0))
          } else if (key === " ") {
            e.preventDefault()
            setLogFollowMode(prev => !prev)
          }
        } else if (logViewerFocusArea === "filters") {
          const levels: ("ALL" | "INFO" | "WARN" | "ERROR" | "DEBUG")[] = ["ALL", "INFO", "WARN", "ERROR", "DEBUG"]
          const currentIndex = levels.indexOf(logFilterLevel)
          if (key === "h" || key === "arrowleft") {
            setLogFilterLevel(levels[Math.max(0, currentIndex - 1)])
          } else if (key === "l" || key === "arrowright") {
            setLogFilterLevel(levels[Math.min(levels.length - 1, currentIndex + 1)])
          } else if (key === "f") {
            // Cycle through filter levels
            setLogFilterLevel(levels[(currentIndex + 1) % levels.length])
          }
        }
        // Global shortcuts for log viewer
        if (key === "f") {
          const levels: ("ALL" | "INFO" | "WARN" | "ERROR" | "DEBUG")[] = ["ALL", "INFO", "WARN", "ERROR", "DEBUG"]
          const currentIndex = levels.indexOf(logFilterLevel)
          setLogFilterLevel(levels[(currentIndex + 1) % levels.length])
        } else if (key === " " && logViewerFocusArea !== "filters") {
          e.preventDefault()
          setLogFollowMode(prev => !prev)
        }
      }
      
      // Content area - Configuration screen
      if (currentScreen === "config" && focusArea === "content") {
        // Handle edit mode
        if (configEditMode) {
          if (e.key === "Escape") {
            setConfigEditMode(false)
            setConfigEditValue("")
            return
          }
          if (e.key === "Enter") {
            // Save and exit edit mode
            setConfigEditMode(false)
            setConfigEditValue("")
            return
          }
          if (e.key === "Backspace") {
            setConfigEditValue(prev => prev.slice(0, -1))
            return
          }
          if (e.key.length === 1 && !e.ctrlKey && !e.altKey && !e.metaKey) {
            setConfigEditValue(prev => prev + e.key)
            return
          }
          return
        }
        
        if (key === "tab") {
          e.preventDefault()
          setConfigScreenFocusArea(prev => prev === "sections" ? "details" : "sections")
        } else if (configScreenFocusArea === "sections") {
          if (key === "j" || key === "arrowdown") {
            setConfigSectionIndex(prev => Math.min(prev + 1, 2))
            // Reset detail indices when changing sections
            setConfigDetailIndex(0)
            setConfigPathIndex(0)
            setConfigBehaviorIndex(0)
            setConfigRegistryIndex(0)
            setConfigDetectedIndex(0)
          } else if (key === "k" || key === "arrowup") {
            setConfigSectionIndex(prev => Math.max(prev - 1, 0))
            setConfigDetailIndex(0)
            setConfigPathIndex(0)
            setConfigBehaviorIndex(0)
            setConfigRegistryIndex(0)
            setConfigDetectedIndex(0)
          } else if (key === "enter" || key === "l" || key === "arrowright") {
            setConfigScreenFocusArea("details")
          }
        } else if (configScreenFocusArea === "details") {
          // Scan Configuration section
          if (configSectionIndex === 0) {
            if (configDetailIndex === 0) {
              // Scan paths
              if (key === "j" || key === "arrowdown") {
                setConfigPathIndex(prev => Math.min(prev + 1, 2))
              } else if (key === "k" || key === "arrowup") {
                setConfigPathIndex(prev => Math.max(prev - 1, 0))
              } else if (key === " ") {
                e.preventDefault()
                // Toggle path enabled state (would update state in real app)
              } else if (key === "enter") {
                // Enter edit mode for path
                setConfigEditMode(true)
                setConfigEditValue("/home/dev/rfz-workspace")
              } else if (key === "tab") {
                e.preventDefault()
                setConfigDetailIndex(1)
              }
            } else if (configDetailIndex === 1) {
              // Scan behavior
              if (key === "j" || key === "arrowdown") {
                setConfigBehaviorIndex(prev => Math.min(prev + 1, 2))
              } else if (key === "k" || key === "arrowup") {
                setConfigBehaviorIndex(prev => Math.max(prev - 1, 0))
              } else if (key === " ") {
                e.preventDefault()
                // Toggle behavior options
              } else if (key === "tab") {
                e.preventDefault()
                setConfigDetailIndex(0)
              }
            }
          }
          // Component Registry section (read-only)
          else if (configSectionIndex === 1) {
            if (key === "j" || key === "arrowdown") {
              setConfigRegistryIndex(prev => Math.min(prev + 1, 7))
            } else if (key === "k" || key === "arrowup") {
              setConfigRegistryIndex(prev => Math.max(prev - 1, 0))
            }
          }
          // Detected Components section (read-only)
          else if (configSectionIndex === 2) {
            if (key === "j" || key === "arrowdown") {
              setConfigDetectedIndex(prev => Math.min(prev + 1, 7))
            } else if (key === "k" || key === "arrowup") {
              setConfigDetectedIndex(prev => Math.max(prev - 1, 0))
            }
          }
        }
      }
    }
    
    window.addEventListener("keydown", handleKeyDown)
    return () => window.removeEventListener("keydown", handleKeyDown)
  }, [
    appView, focusArea, navIndex, currentScreen,
    buildFocusArea, buildListIndex, buildActionIndex, selectedComponents,
    configFocusArea, configGoalIndex, configProfileIndex, configPortIndex, configActionIndex,
    execFocusArea, execListIndex, execActionIndex, builds, buildComplete,
    discoverFocusArea, discoverListIndex, discoverActionIndex,
    logViewerFocusArea, logViewerIndex, logFilterLevel, logFollowMode, logScrollPosition,
    configSectionIndex, configScreenFocusArea, configDetailIndex, configEditMode, configEditValue,
    configPathIndex, configBehaviorIndex, configRegistryIndex, configDetectedIndex,
    startBuild
  ])
  
  // Generate command for display
  const generateCommand = () => {
    const profiles = Array.from(buildConfig.profiles)
    const portProfile = PORT_OPTIONS.find(p => p.id === buildConfig.port)?.profile
    if (portProfile) {
      profiles.push(portProfile)
    }
    const goal = MAVEN_GOALS.find(g => g.id === buildConfig.mavenGoal)?.label || buildConfig.mavenGoal
    const profileStr = profiles.length > 0 ? ` -P${profiles.join(",")}` : ""
    const skipTestsStr = buildConfig.skipTests ? " -DskipTests" : ""
    return `mvn ${goal}${profileStr}${skipTestsStr}`
  }
  
  return (
    <div className="h-screen w-screen flex flex-col bg-background text-foreground overflow-hidden">
      {/* Header */}
      <header className="flex-shrink-0 border-b border-border">
        <div className="h-1 bg-brand" /> {/* DB Red accent stripe */}
        <TuiBox borderStyle="heavy" title="RFZ-CLI v1.0.0">
          <div className="px-2 py-1 flex items-center justify-between">
            <span className="text-tui-cyan">Terminal Orchestration Tool</span>
            <span className="text-muted-foreground text-sm">
              {new Date().toLocaleTimeString()} | Deutsche Bahn Internal
            </span>
          </div>
        </TuiBox>
      </header>
      
      {/* Main content */}
      <main className="flex-1 flex min-h-0 p-2 gap-2">
        {/* Navigation panel */}
        <TuiNavigation 
          selectedIndex={focusArea === "nav" ? navIndex : -1}
          activeScreen={currentScreen}
          onNavigate={handleNavigate}
        />
        
        {/* Content area */}
        <div className="flex-1 min-w-0">
          <TuiBox 
            title={currentScreen === "welcome" ? "Welcome" : 
                   currentScreen === "build" ? "Build" :
                   currentScreen === "logs" ? "View Logs" :
                   currentScreen === "discover" ? "Discover" : "Configuration"}
            borderStyle="rounded"
            focused={focusArea === "content"}
            className="h-full"
          >
            <div className="p-2 h-full overflow-hidden">
              {currentScreen === "welcome" && <WelcomeScreen />}
              
              {currentScreen === "build" && appView === "main" && (
                <BuildScreen
                  selectedIndex={buildListIndex}
                  selectedComponents={selectedComponents}
                  focusArea={buildFocusArea}
                  actionIndex={buildActionIndex}
                  onOpenBuildConfig={() => setAppView("config-modal")}
                />
              )}
              
                {currentScreen === "build" && appView === "building" && (
                  <BuildExecutionView
                    builds={builds}
                    command={generateCommand()}
                    selectedIndex={execListIndex}
                    focusArea={execFocusArea}
                    actionIndex={execActionIndex}
                    isComplete={buildComplete}
                    onViewLogs={(build) => {
                      setSelectedBuildForLogs(build)
                      setLogsScrollPosition(0)
                      setAppView("logs-modal")
                    }}
                    onRebuildFailed={() => setAppView("main")}
                    onBackToSelection={() => setAppView("main")}
                    onCancelBuild={() => setAppView("main")}
                  />
                )}
              
              {currentScreen === "logs" && (
                <LogViewerScreen
                  selectedIndex={logViewerIndex}
                  focusArea={logViewerFocusArea}
                  filterLevel={logFilterLevel}
                  followMode={logFollowMode}
                  logScrollPosition={logScrollPosition}
                  searchQuery={logSearchQuery}
                />
              )}
              
              {currentScreen === "discover" && (
                <DiscoverScreen
                  selectedIndex={discoverListIndex}
                  focusArea={discoverFocusArea}
                  actionIndex={discoverActionIndex}
                />
              )}
              
              {currentScreen === "config" && (
                <ConfigScreen
                  sectionIndex={configSectionIndex}
                  focusArea={configScreenFocusArea}
                  detailIndex={configDetailIndex}
                  editMode={configEditMode}
                  editValue={configEditValue}
                  pathIndex={configPathIndex}
                  behaviorIndex={configBehaviorIndex}
                  registryIndex={configRegistryIndex}
                  detectedIndex={configDetectedIndex}
                />
              )}
            </div>
          </TuiBox>
        </div>
      </main>
      
      {/* Footer - Lipgloss-style status bar */}
      <footer className="flex-shrink-0">
        <div className="flex text-xs font-mono">
          {/* Screen indicator */}
          <span className="bg-brand text-white px-3 py-1">
            {currentScreen === "welcome" ? "HOME" : 
             currentScreen === "build" ? (appView === "building" ? "BUILD" : "SELECT") :
             currentScreen === "logs" ? "LOGS" :
             currentScreen === "discover" ? "DISCOVER" : "CONFIG"}
          </span>
          
          {/* Focus context */}
          <span className="bg-tui-cyan text-background px-3 py-1">
            {focusArea === "nav" 
              ? NAV_ITEMS[navIndex]?.label || ""
              : currentScreen === "build" && appView === "main"
                ? RFZ_COMPONENTS[buildListIndex]?.label || ""
                : currentScreen === "build" && appView === "building"
                  ? builds[execListIndex]?.name || ""
                  : currentScreen === "logs"
                    ? LOG_COMPONENTS[logViewerIndex]?.name || ""
                    : currentScreen === "discover"
                      ? "Components"
                      : currentScreen === "config"
                        ? "Settings"
                        : ""}
          </span>
          
          {/* Status indicators */}
          {currentScreen === "logs" && logFollowMode && (
            <span className="bg-tui-green text-background px-2 py-1">FOLLOW</span>
          )}
          {currentScreen === "build" && appView === "building" && !buildComplete && (
            <span className="bg-tui-yellow text-background px-2 py-1">RUNNING</span>
          )}
          {currentScreen === "build" && appView === "building" && buildComplete && (
            <span className="bg-tui-green text-background px-2 py-1">COMPLETE</span>
          )}
          
          {/* Spacer */}
          <span className="flex-1 bg-secondary px-3 py-1 text-muted-foreground">
            <span className="text-tui-cyan">Tab</span> Focus
            <span className="mx-2 text-border">│</span>
            <span className="text-tui-cyan">↑↓</span> Nav
            <span className="mx-2 text-border">│</span>
            <span className="text-tui-cyan">Enter</span> Select
            <span className="mx-2 text-border">│</span>
            <span className="text-tui-cyan">Esc</span> Back
          </span>
          
          {/* Quit hint */}
          <span className="bg-secondary border-l border-border px-3 py-1 text-muted-foreground">
            <span className="text-tui-cyan">q</span> Quit
          </span>
        </div>
      </footer>
      
      {/* Build Config Modal */}
      <BuildConfigModal
        isOpen={appView === "config-modal"}
        config={buildConfig}
        focusArea={configFocusArea}
        goalIndex={configGoalIndex}
        profileIndex={configProfileIndex}
        portIndex={configPortIndex}
        actionIndex={configActionIndex}
        selectedComponents={Array.from(selectedComponents)}
        onClose={() => setAppView("main")}
        onStartBuild={startBuild}
        onToggleSkipTests={() => setBuildConfig(prev => ({ ...prev, skipTests: !prev.skipTests }))}
      />
      
      {/* Logs Modal */}
      <LogsModal
        isOpen={appView === "logs-modal"}
        build={selectedBuildForLogs}
        scrollPosition={logsScrollPosition}
        errorOnly={logsErrorOnly}
        onToggleErrorOnly={() => {
          setLogsErrorOnly(prev => !prev)
          setLogsScrollPosition(0)
        }}
        onClose={() => setAppView("building")}
      />
    </div>
  )
}
