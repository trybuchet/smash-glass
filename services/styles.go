package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Theme struct {
	ID   string
	Meta map[string]interface{}
}

type StyleService struct {
	rawBaseURL string
}

func NewStyleService() *StyleService {
	return &StyleService{
		rawBaseURL: "https://raw.githubusercontent.com/trybuchet/smash-soda-registry/main/overlay/themes",
	}
}

func (s *StyleService) DownloadTheme(themeID string, files []string) error {
	execDir, err := GetExecutableDir()
	if err != nil {
		return err
	}

	themeDir := filepath.Join(execDir, "themes", themeID)
	if err := os.MkdirAll(themeDir, os.ModePerm); err != nil {
		return err
	}

	for _, file := range files {
		url := fmt.Sprintf("%s/%s/%s", s.rawBaseURL, themeID, file)
		dst := filepath.Join(themeDir, file)

		err := downloadThemeFile(url, dst)
		if err != nil {
			if err == errThemeNotFound {
				continue
			}
			return err
		}
	}

	return nil
}

func (s *StyleService) GetOverlayStyles() []Theme {
	themes := make([]Theme, 0)

	dir, err := GetExecutableDir()
	if err != nil {
		fmt.Println("Error getting executable directory:", err)
		return themes
	}

	themesDir := filepath.Join(dir, "themes")
	if _, err := os.Stat(themesDir); os.IsNotExist(err) {
		fmt.Println("Themes directory does not exist:", themesDir)
		return themes
	}

	entries, err := os.ReadDir(themesDir)
	if err != nil {
		fmt.Println("Error reading themes directory:", err)
		return themes
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		themeID := entry.Name()
		themePath := filepath.Join(themesDir, themeID)

		metaPath := filepath.Join(themePath, "meta.json")
		metaContent, metaErr := os.ReadFile(metaPath)
		if metaErr != nil || len(metaContent) == 0 {
			metaContent = []byte("{}")
		}

		var meta map[string]interface{}
		if err := json.Unmarshal(metaContent, &meta); err != nil {
			fmt.Println("Error parsing meta file for theme:", themeID, err)
			continue
		}

		themes = append(themes, Theme{
			ID:   themeID,
			Meta: meta,
		})
	}

	return themes
}

func (s *StyleService) GetOverlayThemeCSS(themeID string) (string, error) {
	if themeID == "" {
		return "", fmt.Errorf("theme ID is required")
	}

	cleanID := filepath.Base(themeID)
	if cleanID != themeID {
		return "", fmt.Errorf("invalid theme ID")
	}

	dir, err := GetExecutableDir()
	if err != nil {
		return "", fmt.Errorf("error getting executable directory: %w", err)
	}

	cssPath := filepath.Join(dir, "themes", cleanID, cleanID+".css")
	cssContent, err := os.ReadFile(cssPath)
	if err != nil {
		return "", fmt.Errorf("error reading theme CSS: %w", err)
	}

	return string(cssContent), nil
}

var errThemeNotFound = fmt.Errorf("not found")

func downloadThemeFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return errThemeNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http %d: %s", resp.StatusCode, url)
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
