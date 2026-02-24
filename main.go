package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/GoCodeAlone/claude-skills/skills"
)

//go:embed embedded/gh-cli/SKILL.md
var ghCliFS embed.FS

//go:embed embedded/playwright-cli/SKILL.md
var playwrightCliFS embed.FS

func main() {
	skills.RegisterGhCli(ghCliFS)
	skills.RegisterPlaywrightCli(playwrightCliFS)

	args := os.Args[1:]

	if len(args) == 0 || args[0] == "help" {
		printUsage()
		return
	}

	switch args[0] {
	case "list":
		runList()
	case "install":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Error: install requires a skill name")
			fmt.Fprintln(os.Stderr, "Usage: claude-skills install <name> [--workspace <path>]")
			os.Exit(1)
		}
		name := args[1]
		workspace := "."
		for i := 2; i < len(args)-1; i++ {
			if args[i] == "--workspace" {
				workspace = args[i+1]
				i++
			}
		}
		runInstall(name, workspace)
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command %q\n", args[0])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("claude-skills - Install Claude skills into your workspace")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  claude-skills list")
	fmt.Println("  claude-skills install <name> [--workspace <path>]")
	fmt.Println("  claude-skills help")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  list                List all available skills")
	fmt.Println("  install <name>      Install a skill by name")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --workspace <path>  Target workspace directory (default: current directory)")
}

func runList() {
	fmt.Println("Available skills:")
	fmt.Println()
	for name, skill := range skills.Registry {
		fmt.Printf("  %-20s %s\n", name, skill.Description)
	}
}

func runInstall(name, workspace string) {
	skill, ok := skills.Registry[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: unknown skill %q\n", name)
		fmt.Fprintln(os.Stderr, "Run 'claude-skills list' to see available skills.")
		os.Exit(1)
	}

	fmt.Printf("Installing skill: %s\n", name)
	if err := skill.Install(workspace); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully installed skill %q into %s\n", name, workspace)
}
