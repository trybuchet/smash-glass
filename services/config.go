package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfigService struct {
	config map[string]interface{}
}

func NewConfigService() *ConfigService {
	// Create a new ConfigService instance
	cs := &ConfigService{}

	// Read the config file
	cs.config, _ = cs.ReadConfig()

	// Return the ConfigService instance
	return cs
}

// ReadConfig reads the Smash Soda configuration file from the %appdata% directory
func (cs *ConfigService) ReadConfig() (map[string]interface{}, bool) {
	// Get the %appdata% directory
	appData := os.Getenv("APPDATA")
	if appData == "" {
		fmt.Println("Error: APPDATA environment variable not set")
		return nil, false
	}

	// Construct the path to the config file
	configPath := filepath.Join(appData, "Trybuchet", "Smash Soda", "config.json")

	// Read the file content
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return nil, false
	}

	// Parse the JSON content
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, false
	}

	// Return the JSON data
	return jsonData, true
}

func (cs *ConfigService) GetProp(group, key string) (interface{}, bool) {
	if groupMap, ok := cs.config[group].(map[string]interface{}); ok {
		if value, ok := groupMap[key]; ok {
			return value, true
		}
	}
	return nil, false
}

func (cs *ConfigService) GetExecutableDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exePath), nil
}
