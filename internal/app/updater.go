package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/minio/selfupdate"
)

var Version = "v1.0.0"
var GH_TOKEN = "" // replace with your GitHub token

const URL = "https://api.github.com/repos/zenkiet/window-service-watcher/releases/latest"

type ReleaseInfo struct {
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

type UpdateInfo struct {
	Available      bool   `json:"available"`
	CurrentVersion string `json:"currentVersion,omitempty"`
	LatestVersion  string `json:"latestVersion,omitempty"`
	ReleaseNotes   string `json:"releaseNotes,omitempty"`
	DownloadURL    string `json:"downloadUrl,omitempty"`
}

func (a *App) CheckUpdate() (UpdateInfo, error) {
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Authorization", "token "+GH_TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return UpdateInfo{}, fmt.Errorf("failed to fetch release info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UpdateInfo{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var release ReleaseInfo
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return UpdateInfo{}, fmt.Errorf("failed to decode release info: %w", err)
	}

	current, _ := version.NewVersion(Version)
	latest, err := version.NewVersion(release.TagName)

	if latest.LessThanOrEqual(current) {
		return UpdateInfo{Available: false, CurrentVersion: Version}, nil
	}

	var downloadURL = ""
	targetExt := ".exe"
	for _, asset := range release.Assets {
		if strings.HasSuffix(asset.Name, targetExt) {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		return UpdateInfo{}, fmt.Errorf("no suitable asset found for download")
	}

	return UpdateInfo{
		Available:      true,
		CurrentVersion: Version,
		LatestVersion:  release.TagName,
		ReleaseNotes:   release.Body,
		DownloadURL:    downloadURL,
	}, nil
}

func (a *App) Update(downloadURL string) error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("only supported on windows")
	}

	resp, err := http.Get(downloadURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download update")
	}
	defer resp.Body.Close()

	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	if err != nil {
		return fmt.Errorf("update failed: %v", err)
	}
	return nil
}

func (a *App) RestartApp() {
	selfExecutable, _ := os.Executable()
	cmd := exec.Command(selfExecutable)
	cmd.Start()
	os.Exit(0)
}
