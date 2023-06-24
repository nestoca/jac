package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const jacrcFile = ".jacrc"
const jacHomeDir = ".jac"

type Config struct {
	Dir  string `yaml:"dir,omitempty"`
	Glob string `yaml:"glob,omitempty"`
}

func LoadConfig(explicitDir string) (*Config, error) {
	var cfg *Config

	// Determine different candidate directories to look for config file
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("getting current directory: %w", err)
	}
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("getting home directory: %w", err)
	}
	jacHomeDir := filepath.Join(userHomeDir, jacHomeDir)
	candidateDirs := []string{
		explicitDir,
		currentDir,
		userHomeDir,
		jacHomeDir,
	}

	// Look in all candidate directories for config file
	for _, dir := range candidateDirs {
		if dir == "" {
			continue
		}
		cfg, err = loadConfig(dir)
		if err != nil {
			return nil, fmt.Errorf("loading config from %q: %w", dir, err)
		}
		if cfg != nil {
			cfg.Dir = filepath.Join(dir, cfg.Dir)
			break
		}
	}
	if cfg == nil {
		return nil, fmt.Errorf("found %s config file in none of expected locations: %v", jacrcFile, candidateDirs)
	}

	// Set defaults
	if cfg.Glob == "" {
		cfg.Glob = "**/*.yaml"
	}

	return cfg, nil
}

func loadConfig(dir string) (*Config, error) {
	// Ensure directory exists
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("directory %q does not exist", dir)
		}
		return nil, fmt.Errorf("checking for %s: %w", dir, err)
	}

	// Check if directory contains a config file
	jacrcPath := filepath.Join(dir, jacrcFile)
	if _, err := os.Stat(jacrcPath); err != nil {
		if os.IsNotExist(err) {
			// We're done digging for config files
			return nil, nil
		}
		return nil, fmt.Errorf("checking for %s: %w", jacrcPath, err)
	}

	// Read the YAML file content
	content, err := os.ReadFile(jacrcPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML data into the Config struct
	var cfg Config
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling yaml config file %q: %w", jacrcPath, err)
	}

	if cfg.Dir == "" {
		// We're done digging for config files
		cfg.Dir = dir
		return &cfg, nil
	}

	// Load alternative config
	altDir := filepath.Join(dir, cfg.Dir)
	altCfg, err := loadConfig(altDir)
	if err != nil {
		return nil, fmt.Errorf("loading config from alternative directory %q: %w", altDir, err)
	}
	if altCfg == nil {
		return &cfg, nil
	}

	// Merge alternative config into this one
	if altCfg.Glob != "" {
		cfg.Glob = altCfg.Glob
	}
	if altCfg.Dir != "" {
		cfg.Dir = altCfg.Dir
	}

	// Resolve dir to absolute path
	cfg.Dir, err = filepath.Abs(altDir)
	if err != nil {
		return nil, fmt.Errorf("resolving absolute path for config directory %q: %w", cfg.Dir, err)
	}

	return &cfg, nil
}
