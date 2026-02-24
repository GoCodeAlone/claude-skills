package skills

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

// Skill represents an installable Claude skill.
type Skill struct {
	Name        string
	Description string
	EmbeddedDir embed.FS
	DirPath     string
	PreInstall  func() error
}

// Registry holds all registered skills by name.
var Registry = map[string]*Skill{}

// Register adds a skill to the global registry.
func Register(s *Skill) {
	Registry[s.Name] = s
}

// Install runs the full installation process for the skill into the given workspace directory.
func (s *Skill) Install(workspace string) error {
	if s.PreInstall != nil {
		if err := s.PreInstall(); err != nil {
			return fmt.Errorf("pre-install failed for %s: %w", s.Name, err)
		}
	}

	targetDir := filepath.Join(workspace, ".claude", "skills", s.Name)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create target directory %s: %w", targetDir, err)
	}

	srcPath := s.DirPath + "/SKILL.md"
	data, err := fs.ReadFile(s.EmbeddedDir, srcPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded SKILL.md from %s: %w", srcPath, err)
	}

	targetPath := filepath.Join(targetDir, "SKILL.md")
	if err := os.WriteFile(targetPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write SKILL.md to %s: %w", targetPath, err)
	}

	fmt.Printf("  Installed %s -> %s\n", s.Name, targetPath)
	return nil
}

// commandExists reports whether the named command is available in PATH.
func commandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// brewInstall runs `brew install <pkg>` with stdout/stderr forwarded to the terminal.
func brewInstall(pkg string) error {
	cmd := exec.Command("brew", "install", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
