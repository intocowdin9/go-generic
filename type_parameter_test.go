package go_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T interface{}](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("rafli")
	assert.Equal(t, "rafli", result)

	var resultNumber int = Length[int](212)
	assert.Equal(t, 212, resultNumber)
}
