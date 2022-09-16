package utils

import "os"

// GetConfig retrieves the value of the config variable named by the key.
// It returns the value, which will be empty if the variable is not present.
func GetConfig(key string) string {
	return os.Getenv(key)
}
