# Go-Agent Scripts

A collection of AppleScript and Bash utilities for managing multiple Terminal-based development workspaces with Claude AI agents.

## Quick Start

### Two-Window Mode (Recommended)
```bash
# Start with 2 windows (servers in one, Claude agents in another)
./go2

# Stop two-window workspace
./stop2
```

### Multi-Window Mode (Legacy)
```bash
# Start with 6 separate windows
./go

# Stop all workspaces
./stop
```

## What It Does

This toolkit spawns and manages multiple development workspaces with two organization modes:

### Two-Window Mode (`go2/stop2`) - **Recommended**
- **Window 1 "Servers"**: Three tabs for dev servers (BUGS, MEDIUM, QUICK)
- **Window 2 "Claude Agents"**: Three tabs for Claude AI instances
- Clean separation of concerns
- Easy tab navigation with Cmd+1, Cmd+2, Cmd+3

### Multi-Window Mode (`go/stop`) - Legacy
- 6 separate Terminal windows (3 servers + 3 Claude agents)
- Each workspace gets its own server and Claude window

Both modes provide:
- Automatic git branch management (stash, checkout, pull)
- HTML title synchronization with workspace names
- Unique port assignments (8081, 8082, 8083)
- One-command startup and cleanup

## Installation

1. Clone this repository:
```bash
git clone <repo-url> ~/.scripts
cd ~/.scripts
```

2. Make scripts executable:
```bash
chmod +x *.sh *.scpt go stop
```

3. Grant Terminal accessibility permissions:
   - System Settings → Privacy & Security → Accessibility
   - Add Terminal.app and grant permission

## Main Commands

### Two-Window Mode (Recommended)

#### `go2` - Start Two-Window Workspace
Creates two Terminal windows with tabs:
```bash
./go2          # Uses default 'greg' branch
./go2 main     # Uses 'main' branch
```

**Window Layout:**
- **Window 1 "Servers"**: 
  - Tab 1: BUGS server (port 8081)
  - Tab 2: MEDIUM server (port 8082)
  - Tab 3: QUICK server (port 8083)
- **Window 2 "Claude Agents"**:
  - Tab 1: BUGS Claude
  - Tab 2: MEDIUM Claude
  - Tab 3: QUICK Claude

#### `stop2` - Stop Two-Window Workspace
```bash
./stop2
```

### Multi-Window Mode (Legacy)

#### `go` - Start Everything (6 Windows)
```bash
./go          # Uses default 'greg' branch
./go main     # Uses 'main' branch
```

#### `stop` - Stop Everything
```bash
./stop
```

## Individual Scripts

### Workspace Management

- **`spawn_staff_workspaces`** - Opens all workspaces from `.staff/` directory
- **`cleanup_workspaces`** - Complete cleanup (servers + windows)
- **`close_workspaces`** - Close Terminal windows only
- **`kill_workspace_servers`** - Kill server processes only

### Window Operations

- **`open_terminal`** - Open named Terminal with optional command
- **`run_in_terminal`** - Run command in new Terminal window
- **`run_in_window`** - Run command in existing Terminal window
- **`list_terminals`** - List all open Terminal windows
- **`send_to_claude`** - Send messages to Claude sessions

### Project Setup

- **`project_workspace_two_windows`** - Two-window setup with tabs (new)
- **`project_workspace_single`** - Single workspace with server then Claude
- **`project_workspace_tabs`** - Single window with tabs (requires accessibility)
- **`project_workspace_dual`** - Separate windows for server and Claude
- **`setup_workspaces`** - Configure git branch and HTML titles

## Configuration

### Workspace Directories
Edit `spawn_staff_workspaces` to change workspace locations:
```bash
STAFF_DIR="/Users/gregorywhiteside/Projects/.staff"
```

### Port Assignments
Edit port mappings in `spawn_staff_workspaces`:
```bash
PORTS[BUGS]=8081
PORTS[MEDIUM]=8082
PORTS[QUICK]=8083
```

### Default Branch
The default branch is `greg`. Change it by passing a parameter:
```bash
./go feature-branch
```

## Workflow Examples

### Daily Development (Two-Window Mode)
```bash
# Morning - start fresh with 2 windows
./go2

# Window 1: All servers running in tabs
# Window 2: All Claude agents ready in tabs
# Switch tabs with Cmd+1, Cmd+2, Cmd+3

# Evening - clean shutdown
./stop2
```

### Daily Development (Multi-Window Mode)
```bash
# Morning - start fresh with 6 windows
./go

# Work with Claude agents across workspaces
# Each has its own port, no conflicts

# Evening - clean shutdown
./stop
```

### Task Distribution
```bash
# Send tasks to specific agents
./send_to_claude "Implement auth module" "bugs"
./send_to_claude "Write unit tests" "medium"
./send_to_claude "Update documentation" "quick"
```

### Workspace Monitoring
```bash
# See what's running
./list_terminals

# Check server ports
lsof -i :8081-8083
```

## Requirements

- macOS with Terminal.app
- Node.js and Yarn
- Git
- Claude CLI (`claude` command)
- Projects with `package.json` and `vite` setup

## Troubleshooting

### "Window not found" errors
- Check window names with `./list_terminals`
- Ensure Terminal has accessibility permissions

### Ports already in use
- Run `./stop` to clean up
- Or manually: `./kill_workspace_servers`

### Commands not executing in Terminal
- Grant accessibility permissions to Terminal
- Increase delays in AppleScript if needed

### Server not starting
- Check `package.json` has a valid start script
- Ensure `yarn install` has been run
- Check for port conflicts

## Project Structure

```
.scripts/
├── go2                           # Two-window start (recommended)
├── stop2                         # Two-window stop
├── go                            # Multi-window start (legacy)
├── stop                          # Multi-window stop
├── project_workspace_two_windows # Two-window workspace setup
├── spawn_staff_workspaces        # Spawn all workspaces
├── cleanup_workspaces            # Full cleanup
├── project_workspace_*           # Various workspace creation modes
├── open_terminal                 # Terminal window creation
├── list_terminals                # List windows
├── send_to_claude                # Message Claude sessions
└── *.md                          # Documentation files
```

## Contributing

Feel free to customize these scripts for your workflow. Key files to modify:
- `spawn_staff_workspaces` - Add/remove workspaces and ports
- `go` - Customize startup sequence
- Port assignments and branch defaults

## License

MIT