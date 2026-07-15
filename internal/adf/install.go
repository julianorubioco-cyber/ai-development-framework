package adf

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	embedded "github.com/julianorubioco-cyber/ai-development-framework/internal/assets"
)

type InstallResult struct {
	Target     string   `json:"target"`
	Skills     []string `json:"skills"`
	BackupPath string   `json:"backup_path,omitempty"`
	DryRun     bool     `json:"dry_run"`
}

func embeddedSkillNames() ([]string, error) {
	entries, err := fs.ReadDir(embedded.Files, "assets/skills")
	if err != nil {
		return nil, err
	}
	var names []string
	for _, entry := range entries {
		if entry.IsDir() {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)
	return names, nil
}

func Install(dryRun bool) (*InstallResult, error) {
	target, err := SkillsTarget()
	if err != nil {
		return nil, err
	}
	names, err := embeddedSkillNames()
	if err != nil {
		return nil, err
	}

	var existing []string
	for _, name := range names {
		if info, statErr := os.Stat(filepath.Join(target, name)); statErr == nil && info.IsDir() {
			existing = append(existing, name)
		}
	}

	result := &InstallResult{Target: target, Skills: names, DryRun: dryRun}
	if dryRun {
		return result, nil
	}

	if err := os.MkdirAll(target, 0o755); err != nil {
		return nil, err
	}

	if len(existing) > 0 {
		backups, err := BackupsRoot()
		if err != nil {
			return nil, err
		}
		backup := filepath.Join(backups, time.Now().Format("20060102-150405"))
		if err := os.MkdirAll(backup, 0o755); err != nil {
			return nil, err
		}
		for _, name := range existing {
			if err := copyDirectory(filepath.Join(target, name), filepath.Join(backup, name)); err != nil {
				return nil, fmt.Errorf("backup de %s: %w", name, err)
			}
		}
		result.BackupPath = backup
	}

	for _, name := range names {
		destination := filepath.Join(target, name)
		if err := os.RemoveAll(destination); err != nil {
			return nil, err
		}
		sourceRoot := "assets/skills/" + name
		if _, _, err := copyEmbeddedTree(embedded.Files, sourceRoot, destination, false); err != nil {
			return nil, err
		}
	}

	version := strings.TrimSpace(Version)
	if version == "" {
		version = "dev"
	}
	if err := saveManifest(Manifest{
		Framework:  "AI Development Framework",
		Version:    version,
		Installed:  installationTime(),
		OS:         runtime.GOOS,
		Arch:       runtime.GOARCH,
		Skills:     names,
		BackupPath: result.BackupPath,
	}); err != nil {
		return nil, err
	}
	return result, nil
}
