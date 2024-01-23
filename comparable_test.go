package go_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1 T, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.Equal(t, true, IsSame[string]("rafli", "rafli"))
	assert.Equal(t, true, IsSame[int](100, 1200))
}
