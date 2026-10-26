package BetterConfig

import (
	"os"
	"path/filepath"
)

func provideConfigFolderPath(filename string) (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir = filepath.Join(dir, filename)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
}

func isInitDone() bool {
	return isInit == true
}
func getThisConfigFolder() string {
	return configFilePath
}
