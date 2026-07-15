package adf

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindProjectRoot(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	nested := filepath.Join(root, "a", "b")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	found, err := FindProjectRoot(nested)
	if err != nil {
		t.Fatal(err)
	}
	if found != root {
		t.Fatalf("esperava %s, recebeu %s", root, found)
	}
}

func TestInitWorkspacePreservesExistingFile(t *testing.T) {
	project := t.TempDir()
	custom := filepath.Join(project, ".claude", "context.md")
	if err := os.MkdirAll(filepath.Dir(custom), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(custom, []byte("custom"), 0o644); err != nil {
		t.Fatal(err)
	}
	result, err := InitWorkspace(project, false)
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(custom)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "custom" {
		t.Fatal("arquivo existente foi sobrescrito")
	}
	if result.Workspace != filepath.Join(project, ".claude") {
		t.Fatal("workspace criado fora do projeto")
	}
}

func TestEmbeddedSkills(t *testing.T) {
	names, err := embeddedSkillNames()
	if err != nil {
		t.Fatal(err)
	}
	if len(names) < 10 {
		t.Fatalf("poucas Skills embutidas: %d", len(names))
	}
}
