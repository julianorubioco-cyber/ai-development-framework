package adf

import (
	"os"
	"path/filepath"

	embedded "github.com/julianorubioco-cyber/ai-development-framework/internal/assets"
)

type InitResult struct {
	ProjectRoot string   `json:"project_root"`
	Workspace   string   `json:"workspace"`
	Created     []string `json:"created"`
	Preserved   []string `json:"preserved"`
	DryRun      bool     `json:"dry_run"`
}

func FindProjectRoot(start string) (string, error) {
	current, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}

	for {
		if info, statErr := os.Stat(filepath.Join(current, ".git")); statErr == nil && info.IsDir() {
			return current, nil
		}
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}
	return filepath.Abs(start)
}

func InitWorkspace(path string, dryRun bool) (*InitResult, error) {
	if path == "" {
		path = "."
	}
	root, err := FindProjectRoot(path)
	if err != nil {
		return nil, err
	}
	workspace := filepath.Join(root, ".claude")
	result := &InitResult{
		ProjectRoot: root,
		Workspace:   workspace,
		DryRun:      dryRun,
	}
	if dryRun {
		return result, nil
	}

	created, preserved, err := copyEmbeddedTree(
		embedded.Files,
		"assets/templates/project-workspace/.claude",
		workspace,
		true,
	)
	if err != nil {
		return nil, err
	}
	result.Created = created
	result.Preserved = preserved
	return result, nil
}
