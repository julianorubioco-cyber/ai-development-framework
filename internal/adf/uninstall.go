package adf

import (
	"errors"
	"os"
	"path/filepath"
)

type UninstallResult struct {
	Skills     []string `json:"skills"`
	Restored   bool     `json:"restored_backup"`
	BackupPath string   `json:"backup_path,omitempty"`
	DryRun     bool     `json:"dry_run"`
}

func Uninstall(restoreBackup, dryRun bool) (*UninstallResult, error) {
	manifest, err := loadManifest()
	if err != nil {
		return nil, err
	}
	if manifest == nil {
		return nil, errors.New("manifesto do ADF não encontrado; execute 'adf doctor'")
	}
	result := &UninstallResult{
		Skills:     manifest.Skills,
		BackupPath: manifest.BackupPath,
		DryRun:     dryRun,
	}
	if dryRun {
		return result, nil
	}

	target, err := SkillsTarget()
	if err != nil {
		return nil, err
	}
	for _, name := range manifest.Skills {
		if err := os.RemoveAll(filepath.Join(target, name)); err != nil {
			return nil, err
		}
	}

	if restoreBackup && manifest.BackupPath != "" {
		entries, readErr := os.ReadDir(manifest.BackupPath)
		if readErr == nil {
			for _, entry := range entries {
				if !entry.IsDir() {
					continue
				}
				source := filepath.Join(manifest.BackupPath, entry.Name())
				destination := filepath.Join(target, entry.Name())
				if err := os.RemoveAll(destination); err != nil {
					return nil, err
				}
				if err := copyDirectory(source, destination); err != nil {
					return nil, err
				}
			}
			result.Restored = true
		}
	}

	path, err := ManifestPath()
	if err != nil {
		return nil, err
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return result, nil
}
