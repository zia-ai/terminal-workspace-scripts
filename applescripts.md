# AppleScript Utilities

Collection of AppleScript utilities for macOS automation and workflow enhancement.

## Terminal Window Manager (`open_terminal`)

**Purpose:** Opens new Terminal windows with custom names and random color profiles for visual organization.

**Location:** `/Users/gregorywhiteside/Projects/open_terminal`

**Usage:**
```bash
# Basic usage - opens terminal with random profile
./open_terminal "window-name"

# With startup command
./open_terminal "window-name" "cd ~/Projects && ls"

# With specific profile
./open_terminal "window-name" "command" "Ocean"

# Quick workspace setup
./open_terminal "quick"
./open_terminal "bugs" 
./open_terminal "roadmap"
```

**Features:**
- Automatically assigns random Terminal profile colors when none specified
- Sets custom window titles for easy identification
- Optionally runs commands on window creation
- Executable from command line

**Use Cases:**
- Create dedicated terminal windows for different projects/tasks
- Visual separation of work contexts using color profiles
- Quick workspace bootstrapping with multiple named terminals
- Consistent terminal setup across work sessions

**Installation:**
```bash
chmod +x open_terminal
```

---

## Command Runner (`run_in_terminal`)

**Purpose:** Execute arbitrary commands in new Terminal windows with optional naming and profiles.

**Location:** `/Users/gregorywhiteside/Projects/run_in_terminal`

**Usage:**
```bash
# Basic - run command in new terminal
./run_in_terminal "npm run dev"

# With window title
./run_in_terminal "npm run dev" "Dev Server"

# With window title and specific profile
./run_in_terminal "npm test --watch" "Test Runner" "Ocean"

# Complex commands
./run_in_terminal "cd ~/Projects && git status" "Git Status"
./run_in_terminal "tail -f /var/log/system.log" "System Logs"
./run_in_terminal "python3 -m http.server 8000" "HTTP Server"
```

**Features:**
- Executes any shell command in a new Terminal window
- Optional custom window titles for organization
- Random or specific Terminal profile selection
- Supports complex commands with pipes, redirects, etc.
- Terminal window stays open after command completes

**Use Cases:**
- Launch long-running processes (servers, watchers, logs)
- Run commands that need separate terminal context
- Quick debugging sessions in isolated windows
- Parallel command execution across multiple terminals

---

## Window Command Runner (`run_in_window`)

**Purpose:** Execute commands in existing Terminal windows by title or index.

**Location:** `/Users/gregorywhiteside/Projects/run_in_window`

**Usage:**
```bash
# Run in current active window
./run_in_window "ls -la"

# Run in window by custom title
./run_in_window "git status" "quick"
./run_in_window "npm test" "bugs"

# Run in window by index number
./run_in_window "npm run dev" "2"

# Complex example workflow
./open_terminal "server"
./open_terminal "tests"
./run_in_window "npm run dev" "server"
./run_in_window "npm test --watch" "tests"
```

**Features:**
- Targets existing Terminal windows by custom title or index
- Falls back to frontmost window if no identifier provided
- Works with previously named windows from `open_terminal`
- Enables command coordination across multiple terminals

**Use Cases:**
- Send commands to specific workspace terminals
- Update running processes in dedicated windows
- Build automation across multiple terminal contexts
- Avoid creating new windows for related commands

---

## Terminal Window Lister (`list_terminals`)

**Purpose:** List all open Terminal windows with titles and tab counts.

**Location:** `/Users/gregorywhiteside/Projects/list_terminals`

**Usage:**
```bash
# List all Terminal windows
./list_terminals

# Example output:
# Window 1: "✳ List Terminals" [1 tab]
# Window 2: "Quick" [1 tab]
# Window 3: "Terminal" [1 tab]
# Window 4: "bugs" [2 tabs]
```

**Features:**
- Shows window index numbers for use with `run_in_window`
- Displays custom titles when set
- Shows tab count per window
- Identifies untitled windows

**Use Cases:**
- Discover available windows for command targeting
- Verify workspace setup before automation
- Debug window reference issues
- Monitor Terminal workspace organization

---

## Project Workspace Setup (`project_workspace`)

**Purpose:** Automated workflow to set up a development environment with server and Claude in separate tabs.

**Location:** `/Users/gregorywhiteside/Projects/project_workspace`

**Workflow:**
1. Opens new Terminal window and navigates to project directory
2. Creates a second tab in the same window
3. Runs yarn server command in second tab (start/preview/startProd)
4. Returns to first tab and launches Claude

**Usage:**
```bash
# Basic with default 'yarn run start'
./project_workspace /path/to/project

# With specific yarn command
./project_workspace /path/to/project preview
./project_workspace /path/to/project startProd

# With custom window title
./project_workspace ~/Projects/my-app start "MyApp Dev"

# Real examples
./project_workspace ~/Projects/frontend-app
./project_workspace ~/Projects/backend-api preview "API Preview"
```

**Features:**
- Single command to set up complete dev environment
- Automatically manages tab creation and switching
- Runs server in background tab while Claude is active
- Customizable yarn commands (start, preview, startProd)
- Random Terminal profile for visual variety

**Use Cases:**
- Quick project startup with consistent environment
- Development workflow automation
- Ensure server is running before starting Claude session
- Standardize team development setup

---

## Send to Claude (`send_to_claude`)

**Purpose:** Type and send messages to Claude running in Terminal windows.

**Location:** `/Users/gregorywhiteside/Projects/send_to_claude`

**Usage:**
```bash
# Send to currently active Terminal window
./send_to_claude "What files are in the current directory?"

# Send to specific window by title
./send_to_claude "Analyze the authentication flow" "Backend"
./send_to_claude "Check test coverage" "testing"

# Send to window by number
./send_to_claude "Review the latest changes" "2"

# Multi-line messages (use \n)
./send_to_claude "First line\nSecond line"

# Send task assignments
./send_to_claude "Task: Implement user login with OAuth2" "coding"
```

**Features:**
- Types message directly into Claude session
- Automatically presses Enter to send
- Targets specific windows by title or index
- Works with any active Claude session

**Requirements:**
- Terminal needs Accessibility permissions (System Settings > Privacy & Security > Accessibility)
- Claude must be running in the target window

**Use Cases:**
- Distribute tasks to multiple Claude agents
- Send follow-up questions to specific agents
- Automate task assignment across windows
- Coordinate multi-agent workflows

---

## Workspace Setup (`setup_workspaces`)

**Purpose:** Automatically configure all Claude workspace terminals with correct git branch and HTML titles.

**Location:** `/Users/gregorywhiteside/Projects/setup_workspaces`

**What it does:**
1. Iterates through all Terminal windows with custom titles
2. Switches to specified git branch (default: "greg") and pulls latest
3. Updates `index.html` title tag to match the Terminal window name
4. Skips windows labeled as "Server" or untitled windows

**Usage:**
```bash
# Use default branch (greg)
./setup_workspaces

# Specify different branch
./setup_workspaces "development"
./setup_workspaces "main"
```

**Requirements:**
- Terminal needs Accessibility permissions
- Git repositories in workspace directories
- index.html files to update (optional)

**Alternative: Batch Setup Script (`setup_workspaces_batch`)**

For environments without accessibility permissions:

```bash
# Generate setup commands
./setup_workspaces_batch > setup_all.sh

# Edit setup_all.sh to match your workspaces
nano setup_all.sh

# Run in each Terminal window manually
bash setup_all.sh
```

**Features:**
- Ensures consistent git branch across all workspaces
- Syncs HTML page titles with workspace names
- Handles missing branches (creates if needed)
- Safe git operations (won't overwrite local changes)

**Use Cases:**
- Initialize multiple project workspaces after opening terminals
- Ensure all agents work on same branch
- Maintain consistent naming between terminals and web pages
- Quick team environment standardization

---

## Future Scripts

Space for additional AppleScript utilities as they're created.