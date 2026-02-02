"use client"

import { cn } from "@/lib/utils"
import { TuiModal, TuiModalActions, TuiButton } from "../tui/tui-modal"
import type { ComponentBuildState } from "./build-execution-view"

interface LogsModalProps {
  isOpen: boolean
  build: ComponentBuildState | null
  scrollPosition: number
  errorOnly?: boolean
  onToggleErrorOnly?: () => void
  onClose?: () => void
  className?: string
}

const SAMPLE_BUILD_LOGS = [
  "[INFO] Scanning for projects...",
  "[INFO] ",
  "[INFO] ----------------------< de.db.rfz:boss >-----------------------",
  "[INFO] Building boss 2.1.0-SNAPSHOT",
  "[INFO] --------------------------------[ jar ]---------------------------------",
  "[INFO] ",
  "[INFO] --- maven-clean-plugin:3.2.0:clean (default-clean) @ boss ---",
  "[INFO] Deleting /home/dev/rfz-workspace/boss/target",
  "[INFO] ",
  "[INFO] --- maven-resources-plugin:3.3.0:resources (default-resources) @ boss ---",
  "[INFO] Copying 15 resources",
  "[INFO] ",
  "[INFO] --- maven-compiler-plugin:3.11.0:compile (default-compile) @ boss ---",
  "[INFO] Changes detected - recompiling the module!",
  "[INFO] Compiling 127 source files to /home/dev/rfz-workspace/boss/target/classes",
  "[WARNING] /home/dev/rfz-workspace/boss/src/main/java/de/db/rfz/boss/util/LegacyHelper.java:[23,5] [deprecation] OldMethod in OldClass has been deprecated",
  "[INFO] ",
  "[INFO] --- maven-resources-plugin:3.3.0:testResources (default-testResources) @ boss ---",
  "[INFO] Copying 8 resources",
  "[INFO] ",
  "[INFO] --- maven-compiler-plugin:3.11.0:testCompile (default-testCompile) @ boss ---",
  "[INFO] Changes detected - recompiling the module!",
  "[INFO] Compiling 45 test source files to /home/dev/rfz-workspace/boss/target/test-classes",
  "[INFO] ",
  "[INFO] --- maven-surefire-plugin:3.0.0:test (default-test) @ boss ---",
  "[INFO] Using auto detected provider org.apache.maven.surefire.junitplatform.JUnitPlatformProvider",
  "[INFO] ",
  "[INFO] -------------------------------------------------------",
  "[INFO]  T E S T S",
  "[INFO] -------------------------------------------------------",
  "[INFO] Running de.db.rfz.boss.core.CoreModuleTest",
  "[INFO] Tests run: 23, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 1.234 s - in de.db.rfz.boss.core.CoreModuleTest",
  "[INFO] Running de.db.rfz.boss.service.ServiceLayerTest",
  "[INFO] Tests run: 15, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.876 s - in de.db.rfz.boss.service.ServiceLayerTest",
  "[INFO] Running de.db.rfz.boss.integration.IntegrationTest",
  "[INFO] Tests run: 8, Failures: 0, Errors: 0, Skipped: 2, Time elapsed: 2.341 s - in de.db.rfz.boss.integration.IntegrationTest",
  "[INFO] ",
  "[INFO] Results:",
  "[INFO] ",
  "[INFO] Tests run: 46, Failures: 0, Errors: 0, Skipped: 2",
  "[INFO] ",
  "[INFO] --- maven-jar-plugin:3.3.0:jar (default-jar) @ boss ---",
  "[INFO] Building jar: /home/dev/rfz-workspace/boss/target/boss-2.1.0-SNAPSHOT.jar",
  "[INFO] ",
  "[INFO] --- maven-install-plugin:3.1.0:install (default-install) @ boss ---",
  "[INFO] Installing /home/dev/rfz-workspace/boss/target/boss-2.1.0-SNAPSHOT.jar to /home/dev/.m2/repository/de/db/rfz/boss/2.1.0-SNAPSHOT/boss-2.1.0-SNAPSHOT.jar",
  "[INFO] Installing /home/dev/rfz-workspace/boss/pom.xml to /home/dev/.m2/repository/de/db/rfz/boss/2.1.0-SNAPSHOT/boss-2.1.0-SNAPSHOT.pom",
  "[INFO] ------------------------------------------------------------------------",
  "[INFO] BUILD SUCCESS",
  "[INFO] ------------------------------------------------------------------------",
  "[INFO] Total time:  12.456 s",
  "[INFO] Finished at: 2025-01-28T14:32:15+01:00",
  "[INFO] ------------------------------------------------------------------------",
]

const SAMPLE_FAILED_LOGS = [
  "[INFO] Scanning for projects...",
  "[INFO] ",
  "[INFO] ---------------------< de.db.rfz:fistiv >----------------------",
  "[INFO] Building fistiv 2.1.0-SNAPSHOT",
  "[INFO] --------------------------------[ jar ]---------------------------------",
  "[INFO] ",
  "[INFO] --- maven-clean-plugin:3.2.0:clean (default-clean) @ fistiv ---",
  "[INFO] Deleting /home/dev/rfz-workspace/fistiv/target",
  "[INFO] ",
  "[INFO] --- maven-resources-plugin:3.3.0:resources (default-resources) @ fistiv ---",
  "[INFO] Copying 12 resources",
  "[INFO] ",
  "[INFO] --- maven-compiler-plugin:3.11.0:compile (default-compile) @ fistiv ---",
  "[INFO] Changes detected - recompiling the module!",
  "[ERROR] COMPILATION ERROR :",
  "[ERROR] /home/dev/rfz-workspace/fistiv/src/main/java/de/db/rfz/fistiv/handler/MessageHandler.java:[45,23] cannot find symbol",
  "[ERROR]   symbol:   method processAsync(Message)",
  "[ERROR]   location: class de.db.rfz.fistiv.service.MessageService",
  "[ERROR] /home/dev/rfz-workspace/fistiv/src/main/java/de/db/rfz/fistiv/handler/MessageHandler.java:[67,12] incompatible types: void cannot be converted to Response",
  "[INFO] 2 errors ",
  "[INFO] ------------------------------------------------------------------------",
  "[INFO] BUILD FAILURE",
  "[INFO] ------------------------------------------------------------------------",
  "[INFO] Total time:  3.234 s",
  "[INFO] Finished at: 2025-01-28T14:32:45+01:00",
  "[INFO] ------------------------------------------------------------------------",
  "[ERROR] Failed to execute goal org.apache.maven.plugins:maven-compiler-plugin:3.11.0:compile (default-compile) on project fistiv: Compilation failure",
  "[ERROR] -> [Help 1]",
]

function getLogLineClass(line: string): string {
  if (line.startsWith("[ERROR]")) return "text-destructive"
  if (line.startsWith("[WARNING]")) return "text-tui-yellow"
  if (line.includes("BUILD SUCCESS")) return "text-tui-green font-bold"
  if (line.includes("BUILD FAILURE")) return "text-destructive font-bold"
  if (line.includes("Tests run:")) return "text-tui-cyan"
  if (line.startsWith("[INFO] ---")) return "text-muted-foreground"
  return "text-foreground"
}

export function LogsModal({
  isOpen,
  build,
  scrollPosition,
  errorOnly = false,
  onToggleErrorOnly,
  onClose,
}: LogsModalProps) {
  if (!build) return null
  
  const allLogs = build.status === "failed" || build.status === "error" 
    ? SAMPLE_FAILED_LOGS 
    : SAMPLE_BUILD_LOGS
    
  const logs = errorOnly 
    ? allLogs.filter(line => line.startsWith("[ERROR]") || line.startsWith("[WARNING]"))
    : allLogs

  return (
    <TuiModal 
      isOpen={isOpen} 
      title={`Build Logs: ${build.name}`}
      width="max-w-4xl"
    >
      <div className="space-y-2">
        {/* Build info */}
        <div className="flex items-center justify-between text-sm">
          <div className="flex items-center gap-6">
            <span>
              <span className="text-muted-foreground">Status:</span>{" "}
              <span className={cn(
                build.status === "success" && "text-tui-green",
                build.status === "failed" && "text-destructive",
                build.status === "running" && "text-tui-yellow"
              )}>
                {build.status.toUpperCase()}
              </span>
            </span>
            <span>
              <span className="text-muted-foreground">Phase:</span>{" "}
              <span className="text-foreground">{build.phase}</span>
            </span>
            <span>
              <span className="text-muted-foreground">Duration:</span>{" "}
              <span className="text-foreground">{Math.floor(build.elapsedTime / 60)}m {build.elapsedTime % 60}s</span>
            </span>
          </div>
          
          {/* Error filter toggle */}
          <button
            type="button"
            onClick={onToggleErrorOnly}
            className={cn(
              "flex items-center gap-2 px-2 py-1 text-xs cursor-pointer transition-colors",
              errorOnly 
                ? "bg-destructive/20 text-destructive" 
                : "bg-secondary text-muted-foreground hover:text-foreground"
            )}
          >
            <span>{errorOnly ? "[x]" : "[ ]"}</span>
            <span>Errors/Warnings Only</span>
            <span className="text-muted-foreground">(e)</span>
          </button>
        </div>

        {/* Log content */}
        <div className="border border-border bg-background max-h-96 overflow-y-auto">
          <div className="p-2 text-xs leading-relaxed">
            {logs.map((line, index) => (
              <div 
                key={index}
                className={cn(
                  "whitespace-pre font-mono",
                  getLogLineClass(line),
                  index === scrollPosition && "bg-muted"
                )}
              >
                <span className="text-muted-foreground w-6 inline-block text-right mr-2">
                  {(index + 1).toString().padStart(3, " ")}
                </span>
                {line}
              </div>
            ))}
          </div>
        </div>

        {/* Scroll indicator */}
        <div className="text-xs text-muted-foreground text-center">
          {logs.length > 0 ? `Line ${scrollPosition + 1} of ${logs.length}` : "No matching logs"} | 
          <span className="text-tui-cyan"> ↑↓</span> or <span className="text-tui-cyan">j/k</span> to scroll | 
          <span className="text-tui-cyan"> e</span> toggle errors
        </div>

        <TuiModalActions>
          <TuiButton
            label="Toggle Errors"
            shortcut="e"
            focused={false}
            onClick={onToggleErrorOnly}
          />
          <TuiButton
            label="Close"
            shortcut="Esc"
            focused={true}
            onClick={onClose}
          />
        </TuiModalActions>
      </div>
    </TuiModal>
  )
}
