package test

import (
	"github.com/stretchr/testify/require"
	"testing"
	"verifyMyTest/m/util"
)

func TestDBConnection(t *testing.T) {

	_, err := util.NewDb()
	require.NoError(t, err)
}
