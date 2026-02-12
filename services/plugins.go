package services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Plugin struct {
	Meta   map[string]interface{}
	Name   string
	Path   string
	HTML   string
	JS     string
	CSS    string
	Config map[string]interface{}
}

type PluginService struct {
	rawBaseURL string
}

func NewPluginService() *PluginService {
	return &PluginService{
		rawBaseURL: "https://raw.githubusercontent.com/trybuchet/smash-soda-registry/main/overlay/plugins",
	}
}

// DownloadPlugin downloads a plugin's files from GitHub raw URL into plugins/<pluginID>/ folder
func (s *PluginService) DownloadPlugin(pluginID string, files []string) error {
	execDir, err := GetExecutableDir()
	if err != nil {
		return err
	}

	pluginDir := filepath.Join(execDir, "plugins", pluginID)
	if err := os.MkdirAll(pluginDir, os.ModePerm); err != nil {
		return err
	}

	for _, file := range files {
		url := fmt.Sprintf("%s/%s/%s", s.rawBaseURL, pluginID, file)
		dst := filepath.Join(pluginDir, file)

		if err := downloadFile(url, dst); err != nil {
			if err == errNotFound {
				continue
			}
			return err
		}
	}

	return nil
}

// LoadPlugins loads all locally saved plugins into memory
func (s *PluginService) LoadPlugins() []Plugin {
	plugins := make([]Plugin, 0)

	dir, err := GetExecutableDir()
	if err != nil {
		fmt.Println("Error getting executable directory:", err)
		return plugins
	}

	pluginsDir := filepath.Join(dir, "plugins")
	if _, err := os.Stat(pluginsDir); os.IsNotExist(err) {
		fmt.Println("Plugins directory does not exist:", pluginsDir)
		return plugins
	}

	pluginDirs, err := ioutil.ReadDir(pluginsDir)
	if err != nil {
		fmt.Println("Error reading plugins directory:", err)
		return plugins
	}

	for _, pluginDir := range pluginDirs {
		if pluginDir.IsDir() {
			pluginName := pluginDir.Name()
			pluginPath := filepath.Join(pluginsDir, pluginName)

			htmlPath := filepath.Join(pluginPath, pluginName+".html")
			htmlContent, htmlErr := ioutil.ReadFile(htmlPath)
			if htmlErr != nil {
				htmlContent = []byte("")
			}

			cssPath := filepath.Join(pluginPath, pluginName+".css")
			cssContent, _ := ioutil.ReadFile(cssPath)

			jsPath := filepath.Join(pluginPath, pluginName+".js")
			jsContent, _ := ioutil.ReadFile(jsPath)

			configPath := filepath.Join(pluginPath, "config.json")
			configContent, _ := ioutil.ReadFile(configPath)
			if len(configContent) == 0 {
				configContent = []byte("{}")
			}

			var config map[string]interface{}
			if err := json.Unmarshal(configContent, &config); err != nil {
				config = map[string]interface{}{}
			}

			metaPath := filepath.Join(pluginPath, "meta.json")
			metaContent, _ := ioutil.ReadFile(metaPath)
			if len(metaContent) == 0 {
				metaContent = []byte("{}")
			}

			var meta map[string]interface{}
			if err := json.Unmarshal(metaContent, &meta); err != nil {
				meta = map[string]interface{}{}
			}

			plugin := Plugin{
				Meta:   meta,
				Name:   pluginName,
				Path:   pluginPath,
				HTML:   string(htmlContent),
				CSS:    string(cssContent),
				JS:     string(jsContent),
				Config: config,
			}

			plugins = append(plugins, plugin)
		}
	}

	return plugins
}

func (s *PluginService) DeletePlugin(pluginId string) error {
	dir, err := GetExecutableDir()
	if err != nil {
		return fmt.Errorf("error getting executable directory: %w", err)
	}

	pluginsDir := filepath.Join(dir, "plugins")
	pluginPath := filepath.Join(pluginsDir, pluginId)

	if err := os.RemoveAll(pluginPath); err != nil {
		return fmt.Errorf("error deleting plugin: %w", err)
	}

	return nil
}

var errNotFound = fmt.Errorf("not found")

func downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return errNotFound
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
