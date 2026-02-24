package skills

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
)

// RegisterPlaywrightCli registers the playwright-cli skill using the provided embedded filesystem.
func RegisterPlaywrightCli(fs embed.FS) {
	Register(&Skill{
		Name:        "playwright-cli",
		Description: "Browser automation - web testing, form filling, screenshots, and data extraction",
		EmbeddedDir: fs,
		DirPath:     "embedded/playwright-cli",
		PreInstall: func() error {
			if !commandExists("playwright-cli") {
				fmt.Println("  Installing playwright-cli via Homebrew...")
				if err := brewInstall("playwright-cli"); err != nil {
					return err
				}
			} else {
				fmt.Println("  playwright-cli is already installed")
			}
			fmt.Println("  Running playwright-cli install --skills...")
			cmd := exec.Command("playwright-cli", "install", "--skills")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		},
	})
}
