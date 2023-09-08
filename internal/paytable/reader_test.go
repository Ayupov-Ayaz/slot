package paytable

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadPayTable(t *testing.T) {
	// todo: implement me
	_, err := ReadPayTable()
	require.NoError(t, err)
}
