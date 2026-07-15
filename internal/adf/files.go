package adf

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func copyEmbeddedTree(source fs.FS, sourceRoot, destinationRoot string, preserve bool) ([]string, []string, error) {
	var created []string
	var preserved []string

	err := fs.WalkDir(source, sourceRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == sourceRoot {
			return nil
		}

		relative, err := filepath.Rel(filepath.FromSlash(sourceRoot), filepath.FromSlash(path))
		if err != nil {
			return err
		}
		destination := filepath.Join(destinationRoot, relative)

		if entry.IsDir() {
			return os.MkdirAll(destination, 0o755)
		}

		if preserve {
			if _, err := os.Stat(destination); err == nil {
				preserved = append(preserved, relative)
				return nil
			} else if !os.IsNotExist(err) {
				return err
			}
		}

		data, err := fs.ReadFile(source, path)
		if err != nil {
			return fmt.Errorf("ler asset %s: %w", path, err)
		}
		if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(destination, data, 0o644); err != nil {
			return fmt.Errorf("gravar %s: %w", destination, err)
		}
		created = append(created, relative)
		return nil
	})
	return created, preserved, err
}

func copyDirectory(source, destination string) error {
	return filepath.WalkDir(source, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, relative)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
}
