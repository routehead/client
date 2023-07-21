package confs

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	Username string
	Token    string
}

func CreateUserConfigFile(config Config) error {
	userDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(userDir, "routehead")

	// Create the folder if it does not exist
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configDir, os.FileMode(0755))
		if err != nil {
			return err
		}
	}

	configFile := filepath.Join(configDir, "config.json")

	if _, err := os.Stat(configFile); err == nil {
		return errors.New("file already exists")
	}

	configFileWriter, createErr := os.Create(configFile)

	if createErr != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	_, err = configFileWriter.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfigFile() (*Config, error) {
	userDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(userDir, "routehead")

	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	configFile := filepath.Join(configDir, "config.json")
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	config := Config{
		Username: result["Username"],
		Token:    result["Token"],
	}
	return &config, nil
}
