package go_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimental(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.1, ExperimentalMin(100.1, 200.2))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Muhammad",
	}

	second := map[string]string{
		"Name": "Muhammad",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlice(t *testing.T) {
	first := []string{"rafli"}
	second := []string{"rafli"}

	assert.True(t, slices.Equal(first, second))
}
