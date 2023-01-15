package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupTeardownConnection(t *testing.T) {
	currDBUser := os.Getenv("DB_USER")
	defer os.Setenv("DB_USER", currDBUser)

	err := SetupConn("api")
	assert.Nil(t, err)

	os.Setenv("DB_USER", "wrong!")
	err = SetupConn("api")
	assert.NotNil(t, err)
}

func TestLastDownMigration(t *testing.T) {
	err := SetupConn("api")
	assert.Nil(t, err)

	// migrate up
	err = MigrateUp()
	assert.Nil(t, err)

	// migrate 1 step down
	err = StepMigrate(-1)
	assert.Nil(t, err)

	// migrate up again
	err = MigrateUp()
	assert.Nil(t, err)
}
