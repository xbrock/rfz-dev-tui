# Layout UI not matching requirements/prototype

## Description
The current implementation of the TUI ui is not matching the prototype in a lot
of ways.

## Issues:

### Navigation

- Nav item active state should be light blue.
- Nav item shortcut should be aligned on the right side of the nav box.
- Active state should overwrite select state. So i an item is selected (j/k/arrows) and active. The item should have the active bg + the arrow. only active -> bg with no arrow.
- nav shortcut hints should be a tree list:
  - ↑/k
  - ↓/j Down
  - Enter Select
  - 1-5 Quick nav
- Nav container should not be full height of the screen but height based on the content if possible.

### Title bar (RFZ-CLI v1.0.0, Terminal Orchestration Tool etc.)

- The red line should be the top of the app and not below the title content.
- Terminal Orchestration Tool and the time + db branding should be on the same line.
- RFZ-CLI v1.0.0 is the main title above Terminal Orchestration Tool and the time

### status bar

- Should have 3 badges on the left: Navigation area (e.g. Build or Logs), current selection (e.g. " fisdv" when currently the cursor is on the fisdv component in the build area), OPTIONAL: state (e.g. if build has finished is says COMPLETE, while build are running it says RUNNING.)
- Then after the badegs also on the left side we have Nav hints: Tab focus, ↑↓/jk Nav, Enter Select, Esc Back
- on the right side we only have q Quit
- The whole bar should have a gray bg over full width

### Home screen

- RFZCLI Logo should look more like this:
  ██████╗ ███████╗███████╗       ██████╗ ██╗     ██╗
  ██╔══██╗██╔════╝╚══███╔╝      ██╔════╝██║     ██║
  ██████╔╝█████╗    ███╔╝ █████╗██║     ██║     ██║
  ██╔══██╗██╔══╝   ███╔╝  ╚════╝██║     ██║     ██║
  ██║  ██║██║     ███████╗      ╚██████╗███████╗██║
  ╚═╝  ╚═╝╚═╝     ╚══════╝       ╚═════╝╚══════╝╚═╝ but with the correct colors and i think some letters are broken because of copy pase
- Terminal Orchestration Tool should be in white color
- the line below the citat should not be a simple line but be more like this:
  ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
-   v1.0.0   Deutsche Bahn  │ Internal Tool should be 3 badges. version red bg, Deutsche Bahn in dark grey bg is light grey font. Internal Tool Blue light blue bg with dark font color.
- the navigation shortcut hints should be in a tree view.

### Build Components

- The None of the shortcut hin in the top row should not move into the second row.
- The checkbox in front of an component name should be  ○ and ◉ (green fill)
- Component category should be aligned right at the end of each row
- Select state should be blue bg with arrow in fron of the dot. active is just a filled dot.
- Actions: [Build Selected] should be green font size by default. active state is green bg and dark font size (light for the command badge (e.g.Enter))
- other actions have light font color by default and active state is blue bg and light font for both label and command batch
- The legend below the actions should display the new select icon

#### Config Modal

- the shortcut should the in blue and the label for it grey. e.g. ←→ or h/l to select everything but or and to select should be blue.- The navigation hints in the bottom should be seperated by | as well.

#### Build started

- component list should be full width
- st and component name alignet left and the rest to the right
- The active/select state should match the list view when selecting components to build
- component list should have the tree icon before the st. so like this 
  - ├─ ○ fisdv
  - ├─ <loading spinner> traktion
  - ├─ ✓ fisdv
- Progress per component can be much smaller (less width) and shold simple be ⣿⣿⣿⣿⣿⣿. Blue while running, red if error and green when sucessfully build
- Overall progress should have more spacing and also should have ░░░░░░░░ and ███████████ (length is just an example and should be adapted)
- we dont need running badge in the progress window.
- Pending should disapier when builds are finished in the progress window.


### General 

- borders lines of boxes are going out of the screen on the right side currently.

## things to consider

- existing comopnent library (change/enhance it if requried)
- existing design protoype screens and webapp as reference (./references)
