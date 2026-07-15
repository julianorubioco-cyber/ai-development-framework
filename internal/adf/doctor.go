package adf

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type DoctorResult struct {
	OS               string   `json:"os"`
	Arch             string   `json:"arch"`
	Version          string   `json:"version"`
	ClaudeFound      bool     `json:"claude_found"`
	GitFound         bool     `json:"git_found"`
	ManifestFound    bool     `json:"manifest_found"`
	InstalledVersion string   `json:"installed_version,omitempty"`
	SkillsTarget     string   `json:"skills_target"`
	MissingSkills    []string `json:"missing_skills"`
	MalformedSkills  []string `json:"malformed_skills"`
	Healthy          bool     `json:"healthy"`
}

func commandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func Doctor() (*DoctorResult, error) {
	target, err := SkillsTarget()
	if err != nil {
		return nil, err
	}
	names, err := embeddedSkillNames()
	if err != nil {
		return nil, err
	}
	manifest, err := loadManifest()
	if err != nil {
		return nil, err
	}

	result := &DoctorResult{
		OS:            runtime.GOOS,
		Arch:          runtime.GOARCH,
		Version:       Version,
		ClaudeFound:   commandExists("claude"),
		GitFound:      commandExists("git"),
		ManifestFound: manifest != nil,
		SkillsTarget:  target,
	}
	if manifest != nil {
		result.InstalledVersion = manifest.Version
	}

	for _, name := range names {
		path := filepath.Join(target, name, "SKILL.md")
		data, readErr := os.ReadFile(path)
		if os.IsNotExist(readErr) {
			result.MissingSkills = append(result.MissingSkills, name)
			continue
		}
		if readErr != nil {
			return nil, readErr
		}
		text := string(data)
		if !strings.HasPrefix(text, "---\n") ||
			!strings.Contains(text, "name:") ||
			!strings.Contains(text, "description:") {
			result.MalformedSkills = append(result.MalformedSkills, name)
		}
	}
	result.Healthy = result.ManifestFound &&
		len(result.MissingSkills) == 0 &&
		len(result.MalformedSkills) == 0
	return result, nil
}
