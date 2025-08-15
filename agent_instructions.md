# Agent Spawning Instructions

Quick reference for spawning Claude agents across multiple Terminal windows and coordinating work between them.

## Basic Agent Setup

### 1. Single Agent in New Window
```bash
# Open named terminal with Claude
./open_terminal "agent-1" "claude"

# Or run Claude in new window with title
./run_in_terminal "claude" "Research Agent"
```

### 2. Project Agent with Server
```bash
# Spawn agent with running server (yarn start/preview/startProd)
./project_workspace_simple /path/to/project
./project_workspace_simple /Users/gregorywhiteside/Projects/BUGS preview
```

## Multi-Agent Workflows

### Spawn Multiple Specialized Agents
```bash
# Create workspace with named terminals
./open_terminal "research" "claude"
./open_terminal "coding" "claude" 
./open_terminal "testing" "claude"
./open_terminal "docs" "claude"
```

### Send Tasks to Specific Agents
```bash
# List available windows
./list_terminals

# Send commands to specific agent windows
./run_in_window "echo 'Task: Research authentication methods' | pbcopy" "research"
./run_in_window "echo 'Task: Implement login component' | pbcopy" "coding"
./run_in_window "echo 'Task: Write test cases' | pbcopy" "testing"
```

## Agent Communication Patterns

### 1. File-Based Communication
Agents can coordinate through shared files:
```bash
# Agent 1 writes status
echo "auth-research: complete" > /tmp/agent_status.txt

# Agent 2 reads status
cat /tmp/agent_status.txt
```

### 2. Git-Based Coordination
Use git commits/branches for handoffs:
```bash
# Research agent commits findings
git add docs/auth-research.md
git commit -m "AGENT_RESEARCH: OAuth2 flow documented"

# Coding agent pulls and continues
git pull
# Sees research is complete, begins implementation
```

### 3. Shared Task List
Create a central task file agents monitor:
```bash
# Create shared todo list
echo "[ ] Research auth providers
[ ] Design auth flow  
[ ] Implement OAuth2
[ ] Write tests
[ ] Update docs" > /tmp/project_tasks.md

# Agents update their progress
sed -i '' 's/\[ \] Research/\[x\] Research/' /tmp/project_tasks.md
```

### 4. Named Pipes for Real-time Communication
```bash
# Create pipe
mkfifo /tmp/agent_pipe

# Agent 1 sends message
echo "AUTH_MODULE: ready for testing" > /tmp/agent_pipe

# Agent 2 receives
cat /tmp/agent_pipe
```

## Workspace Presets

### Full Stack Development
```bash
# Backend agent
./run_in_terminal "cd ~/Projects/api && claude" "Backend"

# Frontend agent  
./run_in_terminal "cd ~/Projects/web && claude" "Frontend"

# Database agent
./run_in_terminal "cd ~/Projects/db && claude" "Database"

# DevOps agent
./run_in_terminal "cd ~/Projects/infra && claude" "DevOps"
```

### Bug Investigation Team
```bash
# Reproduce bug
./open_terminal "reproduce" "claude"

# Debug and trace
./open_terminal "debug" "claude"

# Fix implementation
./open_terminal "fix" "claude"

# Verify fix
./open_terminal "verify" "claude"
```

### Code Review Pipeline
```bash
# Reviewer 1: Architecture
./run_in_terminal "claude" "Architecture Review"

# Reviewer 2: Security
./run_in_terminal "claude" "Security Review"

# Reviewer 3: Performance
./run_in_terminal "claude" "Performance Review"
```

## Advanced Patterns

### Sequential Task Processing
```bash
# Create task queue
echo "task1: Analyze requirements
task2: Design solution
task3: Implement core
task4: Add tests
task5: Document" > /tmp/task_queue.txt

# Each agent takes next task
head -1 /tmp/task_queue.txt
tail -n +2 /tmp/task_queue.txt > /tmp/task_queue.tmp && mv /tmp/task_queue.tmp /tmp/task_queue.txt
```

### Parallel Task Distribution
```bash
# Split work across agents
./run_in_window "echo 'Process files 1-100' | pbcopy" "agent-1"
./run_in_window "echo 'Process files 101-200' | pbcopy" "agent-2"
./run_in_window "echo 'Process files 201-300' | pbcopy" "agent-3"
```

### Agent Status Dashboard
```bash
# Create status monitor in separate window
./run_in_terminal "watch -n 5 'echo \"=== Agent Status ===\"; ls -la /tmp/agent_*.status 2>/dev/null | tail -5; echo \"\n=== Recent Tasks ===\"; tail -5 /tmp/completed_tasks.log 2>/dev/null'" "Monitor"
```

## Best Practices

1. **Name Your Agents**: Use descriptive names for easy identification
2. **Define Clear Boundaries**: Each agent should have a specific role
3. **Use Absolute Paths**: Always use full paths when sharing file locations
4. **Document Handoffs**: Leave clear notes when passing work between agents
5. **Clean Up**: Remove temporary files and pipes when done

## Quick Commands Reference

```bash
# Spawn agent with custom name
./open_terminal "name" "claude"

# Run command in specific window
./run_in_window "command" "window-name"

# List all windows
./list_terminals

# Project setup with server
./project_workspace_simple /full/path/to/project [start|preview|startProd]

# Execute in new terminal
./run_in_terminal "command" "window-title"
```

## Troubleshooting

- **"Window not found"**: Check exact window name with `./list_terminals`
- **Path issues**: Always use absolute paths (starting with /)
- **Permission denied**: Check file permissions and Terminal accessibility settings
- **Commands not running**: Ensure scripts are executable (`chmod +x script_name`)