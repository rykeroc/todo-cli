package internal

import (
	"fmt"
	"os"
	"path"
)

const (
	DefaultLogLevel = 1
	AppName         = "todo"
)

// GetAppConfigDir godoc
//
// Get the path the applications configuration directory.
//
// Creates the app configuration directory if it does not exist.
//
// Returns empty string and error on error.
//
// Returns the app configuration path and nil on success.
func GetAppConfigDir() (string, error) {
	userConfDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("GetAppConfigDir: %v", err)
	}

	// Check if the path already exists
	appConfPath := path.Join(userConfDir, AppName)
	if _, err = os.Stat(appConfPath); err == nil {
		// App conf dir already exists
		return appConfPath, nil
	}

	// App conf path does not exist
	if os.IsNotExist(err) {
		// Create app conf directory
		if err := os.MkdirAll(appConfPath, 0755); err != nil {
			// Error while creating directory
			return "", fmt.Errorf("GetAppConfigDir: %v", err)
		}
		// App conf dir created successfully
		return appConfPath, nil
	}

	// A different error occurred
	return "", fmt.Errorf("GetAppConfigDir: %v", err)
}
