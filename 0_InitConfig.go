package BetterConfig

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName string = "config.json"

var isInit bool
var projectName string
var configFilePath string

type preferences map[string]any

var valuesVar preferences

func InitBetterConfig(ProjectNameFolder string) error {
	if ProjectNameFolder == "" {
		return fmt.Errorf("variable ProjectName can't be empty")
	}

	if isInit {
		return nil
	}

	cfgFolder, err := provideConfigFolderPath(ProjectNameFolder)
	if err != nil {
		return err
	}

	projectName = ProjectNameFolder
	configFilePath = filepath.Join(cfgFolder, configFileName)

	_, err = os.Stat(configFilePath)
	if errors.Is(err, os.ErrNotExist) {
		// slog.Debug("BetterConfig-init : cfg file not exists. Creating new.")
		err2 := os.WriteFile(configFilePath, []byte("{}"), os.FileMode(0775))
		if err2 != nil {
			return err2
		}
	} else if err != nil {
		return err
	}

	b, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&valuesVar)

	isInit = true

	return err
}

func (p *preferences) Save() error {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, b, 0775)
	return err
}
