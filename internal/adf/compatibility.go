package adf

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
)

type CompatibilityProfile struct {
	Mode                string            `json:"mode"`
	DetectedSystems     []string          `json:"detected_systems"`
	ProjectRoot         string            `json:"project_root"`
	BusinessMemoryRoot  string            `json:"business_memory_root,omitempty"`
	ProjectInstructions string            `json:"project_instructions,omitempty"`
	IdentityRoot        string            `json:"identity_root,omitempty"`
	MarketingRoot       string            `json:"marketing_root,omitempty"`
	DataRoot            string            `json:"data_root,omitempty"`
	OutputRoot          string            `json:"output_root,omitempty"`
	TechnicalArtifacts  map[string]string `json:"technical_artifacts"`
	Conflicts           []string          `json:"conflicts"`
	Recommendations     []string          `json:"recommendations"`
}

func existsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func existsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func detectFirstDir(root string, names ...string) string {
	for _, name := range names {
		path := filepath.Join(root, name)
		if existsDir(path) {
			return path
		}
	}
	return ""
}

func detectFirstFile(root string, names ...string) string {
	for _, name := range names {
		path := filepath.Join(root, name)
		if existsFile(path) {
			return path
		}
	}
	return ""
}

func DetectCompatibility(path string) (*CompatibilityProfile, error) {
	root, err := FindProjectRoot(path)
	if err != nil {
		return nil, err
	}

	profile := &CompatibilityProfile{
		Mode:               "native",
		ProjectRoot:        root,
		TechnicalArtifacts: map[string]string{},
	}

	profile.BusinessMemoryRoot = detectFirstDir(
		root,
		"__memoria",
		"_memoria",
		"memoria",
		"memory",
		"knowledge",
	)
	profile.ProjectInstructions = detectFirstFile(root, "CLAUDE.md")
	profile.IdentityRoot = detectFirstDir(root, "identidade", "brand", "branding")
	profile.MarketingRoot = detectFirstDir(root, "marketing")
	profile.DataRoot = detectFirstDir(root, "dados", "data", "inputs")
	profile.OutputRoot = detectFirstDir(root, "saidas", "outputs", "dist")

	if profile.BusinessMemoryRoot != "" {
		profile.DetectedSystems = append(profile.DetectedSystems, "external-business-memory")
	}
	if profile.ProjectInstructions != "" {
		profile.DetectedSystems = append(profile.DetectedSystems, "root-claude-instructions")
	}
	if profile.IdentityRoot != "" || profile.MarketingRoot != "" {
		profile.DetectedSystems = append(profile.DetectedSystems, "business-workspace")
	}

	// MazyOS heuristic: its commonly observed structure.
	if existsDir(filepath.Join(root, "__memoria")) &&
		existsFile(filepath.Join(root, "CLAUDE.md")) &&
		(existsDir(filepath.Join(root, "identidade")) ||
			existsDir(filepath.Join(root, "marketing"))) {
		profile.DetectedSystems = append(profile.DetectedSystems, "mazyos-compatible")
	}

	if len(profile.DetectedSystems) > 0 {
		profile.Mode = "compatibility"
	}

	adfRoot := filepath.Join(root, ".claude")
	profile.TechnicalArtifacts["specs"] = filepath.Join(adfRoot, "specs")
	profile.TechnicalArtifacts["plans"] = filepath.Join(adfRoot, "plans")
	profile.TechnicalArtifacts["preflights"] = filepath.Join(adfRoot, "preflights")
	profile.TechnicalArtifacts["reviews"] = filepath.Join(adfRoot, "reviews")
	profile.TechnicalArtifacts["releases"] = filepath.Join(adfRoot, "releases")
	profile.TechnicalArtifacts["history"] = filepath.Join(adfRoot, "history")

	if profile.Mode == "compatibility" {
		profile.Recommendations = append(profile.Recommendations,
			"reutilizar a memória de negócio existente como fonte de verdade",
			"não criar company.md, context.md ou memory/ paralelos com conteúdo de negócio",
			"usar o ADF apenas para artefatos técnicos e fluxo de engenharia",
		)
	}

	if profile.BusinessMemoryRoot != "" &&
		existsDir(filepath.Join(adfRoot, "memory")) {
		profile.Conflicts = append(profile.Conflicts,
			"há memória externa e .claude/memory; escolha uma única fonte de verdade",
		)
	}
	if profile.ProjectInstructions != "" &&
		existsFile(filepath.Join(adfRoot, "CLAUDE.md")) {
		profile.Conflicts = append(profile.Conflicts,
			"há CLAUDE.md na raiz e em .claude; mantenha instruções canônicas na raiz",
		)
	}

	sort.Strings(profile.DetectedSystems)
	return profile, nil
}

func SaveCompatibilityProfile(profile *CompatibilityProfile) error {
	target := filepath.Join(profile.ProjectRoot, ".claude", "compatibility.json")
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(target, data, 0o644)
}
