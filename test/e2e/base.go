package e2e_test

import (
	"testing"

	"github.com/misikdmytro/task-tracker/internal/bootstrap"
	"github.com/stretchr/testify/require"
)

func Setup(t *testing.T) (*bootstrap.Server, func() error) {
	s, err := bootstrap.NewServer("../../config/config.yaml")
	require.NoError(t, err)

	go s.ListenAndServe()

	return s, s.Close
}
