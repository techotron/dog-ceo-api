package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppConfigFromEnv(t *testing.T) {
	// Environment variables sourced from /.vscode/settings.json
	testAppConfig, testError := GetAppConfigFromEnv()
	assert.Equal(t, "app", testAppConfig.DBUser)
	assert.Nil(t, testError)

	os.Unsetenv("DB_USER") // Remove a single env var to test error
	_, isPresent := os.LookupEnv("DB_USER")
	assert.False(t, isPresent)
	testAppConfig, testError = GetAppConfigFromEnv()
	assert.NotNil(t, testError)
}
