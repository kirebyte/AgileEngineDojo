package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	assert.Equal(t, Solve([]int{}), 0)
	assert.Equal(t, Solve([]int{1, 2, 3, 4}), 7)
	assert.Equal(t, Solve([]int{1, 2, 3, 4, 5, 6}), 13)
	assert.Equal(t, Solve([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}), 47)
}
