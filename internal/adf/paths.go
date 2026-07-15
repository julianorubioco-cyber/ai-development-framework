package adf

import (
	"errors"
	"os"
	"path/filepath"
)

func homeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if home == "" {
		return "", errors.New("não foi possível localizar a pasta do usuário")
	}
	return home, nil
}

func ClaudeRoot() (string, error) {
	home, err := homeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".claude"), nil
}

func SkillsTarget() (string, error) {
	root, err := ClaudeRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "skills"), nil
}

func ManifestPath() (string, error) {
	root, err := ClaudeRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "adf-install-manifest.json"), nil
}

func BackupsRoot() (string, error) {
	root, err := ClaudeRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "adf-backups"), nil
}
