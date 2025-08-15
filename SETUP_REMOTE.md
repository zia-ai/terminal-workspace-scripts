# Setting Up Remote Repository

## GitHub

1. Create a new repository on GitHub:
   - Go to https://github.com/new
   - Name: `terminal-workspace-scripts` (or your preferred name)
   - Description: "AppleScript and Bash utilities for managing Terminal workspaces with Claude AI"
   - Make it public or private as desired
   - Don't initialize with README (we already have one)

2. Add the remote and push:
```bash
git remote add origin https://github.com/YOUR_USERNAME/terminal-workspace-scripts.git
git branch -M main
git push -u origin main
```

## Alternative: GitHub CLI

If you have `gh` installed:
```bash
gh repo create terminal-workspace-scripts --public --source=. --remote=origin --push
```

## GitLab

```bash
git remote add origin https://gitlab.com/YOUR_USERNAME/terminal-workspace-scripts.git
git push -u origin main
```

## After Pushing

1. Update README if needed with the actual clone URL
2. Consider adding:
   - License file
   - Contributing guidelines
   - Issue templates

## Sharing

To use these scripts on another machine:
```bash
git clone https://github.com/YOUR_USERNAME/terminal-workspace-scripts.git ~/.scripts
cd ~/.scripts
chmod +x go stop spawn_staff_workspaces cleanup_workspaces
```