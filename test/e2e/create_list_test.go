package e2e_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateListOK(t *testing.T) {
	c := NewClient(BaseAddr)

	response, err := c.CreateList(uuid.NewString())
	require.NoError(t, err)
	assert.Greater(t, response.ID, 0)
}

func TestCreateListLongNameBadRequest(t *testing.T) {
	c := NewClient(BaseAddr)

	_, err := c.CreateList("SArUgjw0jT2Vpfik1ffidrsB0NopE4yplmv8YUIZmaoCPAQBViJzPmVIPVXcjuPkvIP0eB7TUE2L1uKevPAsou0zf6MMDAvZmtGKURxu9bAkbPxn399xa5heQBt11yk2F7RxVflxc6LvUR7CLZ9uGOkFtq6hgLIaaTCwvKmPt4mWKWQUaoTquMTPgzg4KtQT5HFlJndtHD9b7GCuY3WOzM9ErDFN320I72Hnq2iCj5IpuJOkuSBDUjGTSjSqNmRy1BSzbQkzTDVjYOmkfoNaKC8OSta7soPx87URGUSG5iANbyxDD2XcabEXCcETIHEMK7zAA39g0kBRuWpTfOyl67gbx4OMFvNfFo1aL2d6bAGueeDwN9ubQuHfgQEQeLtdlRtNHtgm7qYK0OKct3EsKPm51uVUfmdCzCSeOEGWBOEzXUZshBUXPS5AeGxLcpbpznhJqGrzNgM5")
	assert.Error(t, err, "400 Bad Request")
}
