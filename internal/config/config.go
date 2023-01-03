package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ProjectDir string `yaml:"projectDir"`
	Module     string `yaml:"module"`
	Session    string `yaml:"session"`
}

func SetViperConfig() error {
	configHome, err := getAoCGenHomeDir()
	if err != nil {
		return err
	}

	configName := ".aocgen"
	configType := "yaml"

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configHome)

	if err = viper.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); !ok {
			return err
		}
	}

	return nil
}

func GetUserConfig() (*Config, error) {
	configHome, err := getAoCGenHomeDir()
	if err != nil {
		return nil, err
	}

	contents, err := os.ReadFile(fmt.Sprintf("%s/.aocgen.yaml", configHome))
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err = yaml.Unmarshal(contents, c); err != nil {
		return nil, err
	}

	return c, nil
}

func UpdateUserConfig(cfg *Config) error {
	configHome, err := getAoCGenHomeDir()
	if err != nil {
		return err
	}

	cfg.ProjectDir, err = expandHomeDirectory(cfg.ProjectDir)
	if err != nil {
		return err
	}

	contentBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	confFile := fmt.Sprintf("%s/.aocgen.yaml", configHome)

	return os.WriteFile(confFile, contentBytes, os.ModePerm)
}

func expandHomeDirectory(path string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(path, "$HOME/") || strings.HasPrefix(path, "~/") {
		path = filepath.Join(homeDir, path[2:])
	}

	return path, nil
}

func getAoCGenHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	if err := os.MkdirAll(fmt.Sprintf("%s/.config/aocgen", homeDir), os.ModePerm); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.config/aocgen", homeDir), nil
}
