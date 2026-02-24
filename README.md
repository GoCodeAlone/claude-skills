# claude-skills

A Go CLI tool that installs predefined Claude Code skill files into workspaces, with automatic dependency management via Homebrew.

## What It Does

`claude-skills` installs `SKILL.md` files into a workspace's `.claude/skills/` directory. These skill files give Claude Code structured, domain-specific knowledge for tools like the GitHub CLI and Playwright. Dependencies are installed automatically via Homebrew when missing.

## Installation

```sh
go install github.com/GoCodeAlone/claude-skills@latest
```

Requires Go 1.26+.

## Usage

```sh
# List all available skills
claude-skills list

# Install a skill into the current workspace
claude-skills install <name>

# Install a skill into a specific workspace
claude-skills install <name> --workspace <path>

# Show help
claude-skills help
```

## Available Skills

### gh-cli

GitHub CLI skill for managing pull requests, issues, repositories, releases, workflows, gists, and the GitHub API.

Automatically installs `gh` via Homebrew if not present.

```sh
claude-skills install gh-cli
```

### playwright-cli

Browser automation skill for web testing, form filling, screenshots, and data extraction.

Automatically installs `playwright-cli` via Homebrew if not present, then runs `playwright-cli install --skills`.

```sh
claude-skills install playwright-cli
```

## Building from Source

```sh
git clone https://github.com/GoCodeAlone/claude-skills.git
cd claude-skills
go build -o claude-skills .
```

## Project Structure

```
├── main.go                     # CLI entrypoint with embedded skill files
├── skills/
│   ├── skills.go               # Skill registry and install logic
│   ├── gh_cli.go               # gh-cli skill registration
│   └── playwright_cli.go       # playwright-cli skill registration
└── embedded/
    ├── gh-cli/SKILL.md         # GitHub CLI skill reference
    └── playwright-cli/SKILL.md # Playwright CLI skill reference
```

## License

MIT
