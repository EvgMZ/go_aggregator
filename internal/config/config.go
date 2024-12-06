package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL             string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func (c *Config) SetUser(userName string) error {
	c.Current_user_name = userName
	return write(*c)

}
func Read() (Config, error) {
	filePath, err := getConfigFilerPath()
	// fmt.Println(filePath)
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(filePath)
	// fmt.Println(err)
	if err != nil {
		return Config{}, nil
	}
	defer file.Close()
	// fmt.Println(err)
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	// fmt.Println(cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil

}

func getConfigFilerPath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	file := filepath.Join(path, configFileName)
	return file, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilerPath()
	if err != nil {
		return err
	}
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
