package services

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/image/font/sfnt"
)

// FontInfo holds information about a system font.
type FontInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// FileService struct
type FileService struct{}

// NewFileService creates a new FileService service
func NewFileService() *FileService {
	return &FileService{}
}

// GetFonts returns a list of unique font family names and their file paths
// installed on the system.
func (s *FileService) GetFonts() ([]FontInfo, error) {
	fontDir := "C:\\Windows\\Fonts"
	fontMap := make(map[string]FontInfo) // Map to ensure unique font names

	err := filepath.Walk(fontDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".ttf" && ext != ".otf" && ext != ".ttc" && ext != ".fon" {
			return nil
		}

		fontBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}

		f, err := sfnt.Parse(fontBytes)
		if err != nil {
			return nil
		}

		// Use sfnt.NameIDFamily (ID 1) for the primary font family name
		fmt.Println(f.Name(nil, sfnt.NameIDFamily))
		fmt.Println(f.Name(nil, sfnt.NameIDCompatibleFull))
		fmt.Println(f.Name(nil, sfnt.NameIDFull))
		fmt.Println(f.Name(nil, sfnt.NameIDPostScript))
		fmt.Println(f.Name(nil, sfnt.NameIDTypographicFamily))
		name, err := f.Name(nil, sfnt.NameIDCompatibleFull)
		if err != nil || name == "" {
			// Fallback to sfnt.NameIDFull (ID 4) if family name is not available
			name, err = f.Name(nil, sfnt.NameIDFull)
			if err != nil || name == "" {
				return nil // Skip if no usable name is found
			}
		}

		// Only add if not already present (prioritize first encountered path for a name)
		if _, exists := fontMap[name]; !exists {
			fontMap[name] = FontInfo{Name: name, Path: path}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking font directory: %w", err)
	}

	fonts := make([]FontInfo, 0, len(fontMap))
	for _, font := range fontMap {
		fonts = append(fonts, font)
	}

	sort.Slice(fonts, func(i, j int) bool {
		return fonts[i].Name < fonts[j].Name
	})

	return fonts, nil
}

// ReadHTML reads an HTML file and returns its content as a string.
func (s *FileService) ReadHTML(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadText reads a text file and returns its content as a string.
func (s *FileService) ReadText(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadImage reads an image file and returns it as a base64 encoded string.
func (s *FileService) ReadImage(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(content)
	encoded := base64.StdEncoding.EncodeToString(content)
	return fmt.Sprintf("data:%s;base64,%s", contentType, encoded), nil
}
