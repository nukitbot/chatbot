package common

import (
	"log"
	"os"
)

// GetEnv gets environment variables from the system with the given key. If it
// fails to get the environment variable then it returns the fallback string
func GetEnv(key, fallback string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return fallback
}

// MustHaveEnv gets the environment variable from the system with the given key.
// If it fails to get the environment variable then it will log the error and
// stops the current process
func MustHaveEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Cannot find '%s' Environment variable.", key)
	}
	return v
}
