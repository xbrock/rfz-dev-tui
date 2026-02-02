"use client"

import { cn } from "@/lib/utils"
import { TuiBox } from "../tui/tui-box"

interface WelcomeScreenProps {
  className?: string
}

const ASCII_LOGO = `
██████╗ ███████╗███████╗       ██████╗██╗     ██╗
██╔══██╗██╔════╝╚══███╔╝      ██╔════╝██║     ██║
██████╔╝█████╗    ███╔╝ █████╗██║     ██║     ██║
██╔══██╗██╔══╝   ███╔╝  ╚════╝██║     ██║     ██║
██║  ██║██║     ███████╗      ╚██████╗███████╗██║
╚═╝  ╚═╝╚═╝     ╚══════╝       ╚═════╝╚══════╝╚═╝
`.trim()

export function WelcomeScreen({ className }: WelcomeScreenProps) {
  return (
    <div className={cn(
      "flex items-center justify-center h-full p-4",
      className
    )}>
      <TuiBox 
        borderStyle="double" 
        className="max-w-3xl w-full"
      >
        <div className="flex flex-col items-center py-6 px-4">
          {/* ASCII Logo with gradient effect */}
          <pre className="text-xs leading-tight mb-6 overflow-x-auto">
            <span style={{color: '#5A56E0'}}>██████╗ </span>
            <span style={{color: '#7B5AE4'}}>███████╗</span>
            <span style={{color: '#9C5EE8'}}>███████╗</span>
            <span style={{color: '#BD62EC'}}>       </span>
            <span style={{color: '#DE66F0'}}>██████╗</span>
            <span style={{color: '#EE6FF8'}}>██╗     ██╗</span>
            {"\n"}
            <span style={{color: '#5A56E0'}}>██╔══██╗</span>
            <span style={{color: '#7B5AE4'}}>██╔════╝</span>
            <span style={{color: '#9C5EE8'}}>╚══███╔╝</span>
            <span style={{color: '#BD62EC'}}>      </span>
            <span style={{color: '#DE66F0'}}>██╔════╝</span>
            <span style={{color: '#EE6FF8'}}>██║     ██║</span>
            {"\n"}
            <span style={{color: '#5A56E0'}}>██████╔╝</span>
            <span style={{color: '#7B5AE4'}}>█████╗  </span>
            <span style={{color: '#9C5EE8'}}>  ███╔╝ </span>
            <span style={{color: '#BD62EC'}}>█████╗</span>
            <span style={{color: '#DE66F0'}}>██║     </span>
            <span style={{color: '#EE6FF8'}}>██║     ██║</span>
            {"\n"}
            <span style={{color: '#5A56E0'}}>██╔══██╗</span>
            <span style={{color: '#7B5AE4'}}>██╔══╝  </span>
            <span style={{color: '#9C5EE8'}}> ███╔╝  </span>
            <span style={{color: '#BD62EC'}}>╚════╝</span>
            <span style={{color: '#DE66F0'}}>██║     </span>
            <span style={{color: '#EE6FF8'}}>██║     ██║</span>
            {"\n"}
            <span style={{color: '#5A56E0'}}>██║  ██║</span>
            <span style={{color: '#7B5AE4'}}>██║     </span>
            <span style={{color: '#9C5EE8'}}>███████╗</span>
            <span style={{color: '#BD62EC'}}>      </span>
            <span style={{color: '#DE66F0'}}>╚██████╗</span>
            <span style={{color: '#EE6FF8'}}>███████╗██║</span>
            {"\n"}
            <span style={{color: '#5A56E0'}}>╚═╝  ╚═╝</span>
            <span style={{color: '#7B5AE4'}}>╚═╝     </span>
            <span style={{color: '#9C5EE8'}}>╚══════╝</span>
            <span style={{color: '#BD62EC'}}>      </span>
            <span style={{color: '#DE66F0'}}> ╚═════╝</span>
            <span style={{color: '#EE6FF8'}}>╚══════╝╚═╝</span>
          </pre>
          
          {/* Subtitle */}
          <div className="text-center mb-8">
            <p className="text-lg text-foreground mb-2 font-mono">
              Terminal Orchestration Tool
            </p>
            <p className="text-sm text-muted-foreground italic">
              "First, solve the problem. Then, write the code."
            </p>
          </div>
          
          {/* Decorative separator with braille */}
          <div className="w-full max-w-md text-center text-muted-foreground/40 mb-6 font-mono">
            {"⣿".repeat(25)}
          </div>
          
          {/* Version info - status bar style */}
          <div className="flex items-center text-xs font-mono mb-8">
            <span className="bg-brand text-white px-2 py-0.5">v1.0.0</span>
            <span className="bg-secondary text-muted-foreground px-2 py-0.5">Deutsche Bahn</span>
            <span className="bg-tui-cyan text-background px-2 py-0.5">Internal Tool</span>
          </div>
          
          {/* Blinking cursor prompt */}
          <div className="flex items-center gap-2 text-foreground font-mono">
            <span className="text-tui-green">$</span>
            <span>rfz-cli ready</span>
            <span className="cursor-blink text-tui-cyan">█</span>
          </div>
          
          {/* Navigation hints with tree style */}
          <div className="mt-8 text-center font-mono">
            <p className="text-muted-foreground text-sm mb-3">
              Use navigation panel to get started
            </p>
            <div className="inline-flex flex-col text-xs text-muted-foreground text-left">
              <div><span className="text-border">├── </span><span className="text-tui-cyan">↑↓/jk</span> navigate</div>
              <div><span className="text-border">├── </span><span className="text-tui-cyan">Enter</span> select</div>
              <div><span className="text-border">└── </span><span className="text-tui-cyan">q</span> quit</div>
            </div>
          </div>
        </div>
      </TuiBox>
    </div>
  )
}
