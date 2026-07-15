package adf

import (
	"os"
	"path/filepath"

	embedded "github.com/julianorubioco-cyber/ai-development-framework/internal/assets"
)

type InitResult struct {
	ProjectRoot   string                `json:"project_root"`
	Workspace     string                `json:"workspace"`
	Created       []string              `json:"created"`
	Preserved     []string              `json:"preserved"`
	Skipped       []string              `json:"skipped"`
	Compatibility *CompatibilityProfile `json:"compatibility"`
	DryRun        bool                  `json:"dry_run"`
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
	profile, err := DetectCompatibility(root)
	if err != nil {
		return nil, err
	}

	workspace := filepath.Join(root, ".claude")
	result := &InitResult{
		ProjectRoot:   root,
		Workspace:     workspace,
		Compatibility: profile,
		DryRun:        dryRun,
	}
	if dryRun {
		return result, nil
	}

	if err := os.MkdirAll(workspace, 0o755); err != nil {
		return nil, err
	}

	// In compatibility mode, only create technical artifacts.
	if profile.Mode == "compatibility" {
		technicalDirs := []string{
			"specs", "plans", "preflights", "reviews", "releases", "history",
		}
		for _, name := range technicalDirs {
			path := filepath.Join(workspace, name)
			if existsDir(path) {
				result.Preserved = append(result.Preserved, name)
				continue
			}
			if err := os.MkdirAll(path, 0o755); err != nil {
				return nil, err
			}
			if err := os.WriteFile(filepath.Join(path, ".gitkeep"), []byte{}, 0o644); err != nil {
				return nil, err
			}
			result.Created = append(result.Created, name)
		}

		// Do not create parallel business-memory files.
		result.Skipped = append(result.Skipped,
			"CLAUDE.md",
			"context.md",
			"architecture.md",
			"company.md",
			"decisions.md",
			"memory/",
			"knowledge/",
		)

		if err := SaveCompatibilityProfile(profile); err != nil {
			return nil, err
		}
		result.Created = append(result.Created, "compatibility.json")
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

	if err := SaveCompatibilityProfile(profile); err != nil {
		return nil, err
	}
	result.Created = append(result.Created, "compatibility.json")
	return result, nil
}
