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

func TestDetectMazyOSCompatibility(t *testing.T) {
	project := t.TempDir()
	for _, dir := range []string{"__memoria", "identidade", "marketing"} {
		if err := os.MkdirAll(filepath.Join(project, dir), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(project, "CLAUDE.md"), []byte("# Projeto"), 0o644); err != nil {
		t.Fatal(err)
	}

	profile, err := DetectCompatibility(project)
	if err != nil {
		t.Fatal(err)
	}
	if profile.Mode != "compatibility" {
		t.Fatalf("modo esperado compatibility, recebido %s", profile.Mode)
	}
	found := false
	for _, system := range profile.DetectedSystems {
		if system == "mazyos-compatible" {
			found = true
		}
	}
	if !found {
		t.Fatal("MazyOS não foi detectado")
	}
}

func TestCompatibilityInitDoesNotDuplicateBusinessMemory(t *testing.T) {
	project := t.TempDir()
	if err := os.MkdirAll(filepath.Join(project, "__memoria"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(project, "CLAUDE.md"), []byte("# Projeto"), 0o644); err != nil {
		t.Fatal(err)
	}

	result, err := InitWorkspace(project, false)
	if err != nil {
		t.Fatal(err)
	}
	if result.Compatibility.Mode != "compatibility" {
		t.Fatal("modo compatibilidade não ativado")
	}
	for _, forbidden := range []string{
		"company.md", "context.md", "decisions.md", "architecture.md",
	} {
		if _, err := os.Stat(filepath.Join(project, ".claude", forbidden)); !os.IsNotExist(err) {
			t.Fatalf("arquivo duplicado criado: %s", forbidden)
		}
	}
	for _, required := range []string{"specs", "plans", "reviews", "releases"} {
		if !existsDir(filepath.Join(project, ".claude", required)) {
			t.Fatalf("artefato técnico ausente: %s", required)
		}
	}
}
