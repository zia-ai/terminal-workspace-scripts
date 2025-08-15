# Go-Agent Scripts

A collection of AppleScript and Bash utilities for managing multiple Terminal-based development workspaces with Claude AI agents.

## Quick Start

```bash
# Start all workspaces
./go

# Stop all workspaces
./stop
```

## What It Does

This toolkit spawns and manages multiple development workspaces, each with:
- A Claude AI terminal for development assistance
- A Vite dev server running on a unique port
- Automatic git branch management
- HTML title synchronization with workspace names

Default workspaces:
- **bugs** - Port 8081
- **medium** - Port 8082  
- **quick** - Port 8083

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

### `go` - Start Everything
Cleans up any existing workspaces and spawns fresh ones:
```bash
./go          # Uses default 'greg' branch
./go main     # Uses 'main' branch
```

### `stop` - Stop Everything
Kills all servers and closes all workspace windows:
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

- **`project_workspace_single`** - Setup workspace with git + server + Claude
- **`project_workspace_tabs`** - Same but with tabs (requires accessibility)
- **`project_workspace_dual`** - Same but with separate windows
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

### Daily Development
```bash
# Morning - start fresh
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
├── go                      # Main start script
├── stop                    # Main stop script
├── spawn_staff_workspaces  # Spawn all workspaces
├── cleanup_workspaces      # Full cleanup
├── project_workspace_*     # Workspace creation variants
├── open_terminal          # Terminal window creation
├── list_terminals         # List windows
├── send_to_claude         # Message Claude sessions
└── *.md                   # Documentation files
```

## Contributing

Feel free to customize these scripts for your workflow. Key files to modify:
- `spawn_staff_workspaces` - Add/remove workspaces and ports
- `go` - Customize startup sequence
- Port assignments and branch defaults

## License

MIT