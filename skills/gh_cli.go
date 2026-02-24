package skills

import (
	"embed"
	"fmt"
)

// RegisterGhCli registers the gh-cli skill using the provided embedded filesystem.
func RegisterGhCli(fs embed.FS) {
	Register(&Skill{
		Name:        "gh-cli",
		Description: "GitHub CLI - manage PRs, issues, repos, releases, workflows, and the GitHub API",
		EmbeddedDir: fs,
		DirPath:     "embedded/gh-cli",
		PreInstall: func() error {
			if commandExists("gh") {
				fmt.Println("  gh is already installed")
				return nil
			}
			fmt.Println("  Installing gh via Homebrew...")
			return brewInstall("gh")
		},
	})
}
