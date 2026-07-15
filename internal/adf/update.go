package adf

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const repository = "julianorubioco-cyber/ai-development-framework"

type githubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

type UpdateResult struct {
	CurrentVersion string `json:"current_version"`
	LatestVersion  string `json:"latest_version"`
	Asset          string `json:"asset"`
	Updated        bool   `json:"updated"`
	Reinstalled    bool   `json:"reinstalled_skills"`
	Message        string `json:"message"`
}

func releaseAssetName() (string, error) {
	osName := runtime.GOOS
	if osName != "windows" && osName != "linux" && osName != "darwin" {
		return "", fmt.Errorf("sistema não suportado: %s", osName)
	}
	if runtime.GOARCH != "amd64" && runtime.GOARCH != "arm64" {
		return "", fmt.Errorf("arquitetura não suportada: %s", runtime.GOARCH)
	}
	name := fmt.Sprintf("adf_%s_%s", osName, runtime.GOARCH)
	if osName == "windows" {
		name += ".exe"
	}
	return name, nil
}

func fetchLatestRelease(ctx context.Context) (*githubRelease, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repository)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "ADF/"+Version)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub respondeu %s", resp.Status)
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}
	if release.TagName == "" {
		return nil, errors.New("release sem tag")
	}
	return &release, nil
}

func versionEqual(a, b string) bool {
	normalize := func(v string) string {
		return strings.TrimPrefix(strings.TrimSpace(v), "v")
	}
	return normalize(a) == normalize(b)
}

func downloadFile(ctx context.Context, url, destination string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "ADF/"+Version)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download respondeu %s", resp.Status)
	}

	f, err := os.OpenFile(destination, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

func fileSHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func verifyChecksum(ctx context.Context, release *githubRelease, assetName, filePath string) error {
	var checksumURL string
	for _, asset := range release.Assets {
		if asset.Name == "checksums.txt" {
			checksumURL = asset.BrowserDownloadURL
			break
		}
	}
	if checksumURL == "" {
		return errors.New("checksums.txt não encontrado")
	}

	tmp := filePath + ".checksums"
	defer os.Remove(tmp)
	if err := downloadFile(ctx, checksumURL, tmp); err != nil {
		return err
	}
	data, err := os.ReadFile(tmp)
	if err != nil {
		return err
	}

	var expected string
	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 && filepath.Base(fields[len(fields)-1]) == assetName {
			expected = fields[0]
			break
		}
	}
	if expected == "" {
		return fmt.Errorf("checksum não encontrado para %s", assetName)
	}
	actual, err := fileSHA256(filePath)
	if err != nil {
		return err
	}
	if !strings.EqualFold(expected, actual) {
		return errors.New("checksum inválido; atualização cancelada")
	}
	return nil
}

func replaceCurrentExecutable(downloaded string) error {
	current, err := os.Executable()
	if err != nil {
		return err
	}
	current, err = filepath.EvalSymlinks(current)
	if err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		pending := current + ".new"
		_ = os.Remove(pending)
		if err := os.Rename(downloaded, pending); err != nil {
			return err
		}

		script := current + ".update.cmd"
		content := fmt.Sprintf(`@echo off
setlocal
set "OLD=%s"
set "NEW=%s"
:wait
ping 127.0.0.1 -n 2 >nul
move /Y "%%NEW%%" "%%OLD%%" >nul 2>&1
if errorlevel 1 goto wait
del "%%~f0"
`, current, pending)
		if err := os.WriteFile(script, []byte(content), 0o644); err != nil {
			return err
		}
		return startDetachedWindows(script)
	}

	backup := current + ".bak"
	_ = os.Remove(backup)
	if err := os.Rename(current, backup); err != nil {
		return err
	}
	if err := os.Rename(downloaded, current); err != nil {
		_ = os.Rename(backup, current)
		return err
	}
	if err := os.Chmod(current, 0o755); err != nil {
		return err
	}
	_ = os.Remove(backup)
	return nil
}

func Update(reinstallSkills bool) (*UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	release, err := fetchLatestRelease(ctx)
	if err != nil {
		return nil, err
	}
	assetName, err := releaseAssetName()
	if err != nil {
		return nil, err
	}

	result := &UpdateResult{
		CurrentVersion: Version,
		LatestVersion:  release.TagName,
		Asset:          assetName,
	}

	if versionEqual(Version, release.TagName) {
		result.Message = "ADF já está na versão mais recente"
		if reinstallSkills {
			if _, err := Install(false); err != nil {
				return nil, err
			}
			result.Reinstalled = true
		}
		return result, nil
	}

	var assetURL string
	for _, asset := range release.Assets {
		if asset.Name == assetName {
			assetURL = asset.BrowserDownloadURL
			break
		}
	}
	if assetURL == "" {
		return nil, fmt.Errorf("asset %s não encontrado em %s", assetName, release.TagName)
	}

	tempDir, err := os.MkdirTemp("", "adf-update-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	downloaded := filepath.Join(tempDir, assetName)
	if err := downloadFile(ctx, assetURL, downloaded); err != nil {
		return nil, err
	}
	if err := verifyChecksum(ctx, release, assetName, downloaded); err != nil {
		return nil, err
	}
	if err := replaceCurrentExecutable(downloaded); err != nil {
		return nil, err
	}

	result.Updated = true
	result.Message = "Atualização preparada; abra um novo terminal"

	if reinstallSkills && runtime.GOOS != "windows" {
		if _, err := Install(false); err != nil {
			return nil, err
		}
		result.Reinstalled = true
	}
	return result, nil
}
