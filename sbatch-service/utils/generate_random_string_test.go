package utils_test

import (
	"math/rand"
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/utils"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomString(t *testing.T) {
	// Arrange
	utils.OverrideRandSource(rand.NewSource(0))

	// Act
	v := utils.GenerateRandomString(5)
	v2 := utils.GenerateRandomString(5)

	// Assert
	require.NotEqual(t, v, v2)
	require.Equal(t, "Rczqb", v)
	require.Equal(t, "bblha", v2)
}
