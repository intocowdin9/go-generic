package go_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, int64(200), Min[int64](int64(2000), int64(200)))
	assert.Equal(t, float64(200), Min[float64](float64(2000), float64(200)))
	// this will be error when using Age type
	// to solve this use type approximation '~' beside type inside. for example, ~int
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(2000)))
}

func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(200), Min(int64(2000), int64(200)))
	assert.Equal(t, float64(200), Min(float64(2000), float64(200)))
	assert.Equal(t, Age(100), Min(Age(100), Age(2000)))
}
