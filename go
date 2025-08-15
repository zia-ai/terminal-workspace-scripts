#!/bin/bash

# Quick start script - cleanup and spawn all workspaces

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Branch name (default: greg, or pass as argument)
BRANCH_NAME="${1:-greg}"

clear

echo -e "${BLUE}"
echo "╔══════════════════════════════════════╗"
echo "║     Workspace Quick Start - GO!      ║"
echo "╚══════════════════════════════════════╝"
echo -e "${NC}"

echo -e "${YELLOW}Branch: $BRANCH_NAME${NC}"
echo ""

# Step 1: Cleanup
echo -e "${RED}🧹 Cleaning up existing workspaces...${NC}"
echo "───────────────────────────────────────"

# Kill any existing servers
echo "Stopping servers on ports 8081-8083..."
for port in 8081 8082 8083; do
    PID=$(lsof -ti :$port 2>/dev/null)
    if [ ! -z "$PID" ]; then
        kill -9 $PID 2>/dev/null
        echo "  ✓ Stopped server on port $port"
    fi
done

# Close existing workspace windows
echo "Closing Terminal windows..."
/Users/gregorywhiteside/Projects/.scripts/close_workspaces >/dev/null 2>&1
echo "  ✓ Closed workspace windows"

echo ""
sleep 1

# Step 2: Spawn new workspaces
echo -e "${GREEN}🚀 Spawning fresh workspaces...${NC}"
echo "───────────────────────────────────────"
/Users/gregorywhiteside/Projects/.scripts/spawn_staff_workspaces "$BRANCH_NAME"

echo ""
echo -e "${GREEN}✨ All set! Your workspaces are ready:${NC}"
echo "───────────────────────────────────────"
echo "  📦 BUGS    → Port 8081"
echo "  📦 MEDIUM  → Port 8082"
echo "  📦 QUICK   → Port 8083"
echo ""
echo "Each workspace has:"
echo "  • Claude window (for development)"
echo "  • Server window (running vite)"
echo "  • Git branch: $BRANCH_NAME"
echo ""
echo -e "${BLUE}Happy coding! 🎉${NC}"