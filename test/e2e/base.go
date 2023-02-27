package e2e_test

import (
	"testing"

	"github.com/misikdmytro/task-tracker/internal/bootstrap"
	"github.com/stretchr/testify/require"
)

func Setup(t *testing.T) (*bootstrap.Server, func(), func()) {
	s, err := bootstrap.NewServer("../../config/config.yaml")
	require.NoError(t, err)

	start := func() { go s.ListenAndServe() }
	close := func() { s.Close() }

	return s, start, close
}
