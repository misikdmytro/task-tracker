package database_test

import (
	"testing"

	"github.com/misikdmytro/task-tracker/internal/config"
	"github.com/stretchr/testify/require"
)

func RequireConfig(t *testing.T) *config.Config {
	c, err := config.NewConfig("../../../config/dev.config.yaml")
	require.NoError(t, err)
	return c
}
