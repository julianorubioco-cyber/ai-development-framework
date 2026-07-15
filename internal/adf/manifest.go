package adf

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type Manifest struct {
	Framework  string   `json:"framework"`
	Version    string   `json:"version"`
	Installed  string   `json:"installed_at"`
	OS         string   `json:"os"`
	Arch       string   `json:"arch"`
	Skills     []string `json:"skills"`
	BackupPath string   `json:"backup_path,omitempty"`
}

func saveManifest(manifest Manifest) error {
	path, err := ManifestPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func loadManifest() (*Manifest, error) {
	path, err := ManifestPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}
	return &manifest, nil
}

func installationTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}
