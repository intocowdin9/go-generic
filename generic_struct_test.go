package go_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "muhammad",
		Second: "rafli",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "muhammad",
		Second: "rafli",
	}

	assert.Equal(t, "muhammad", data.ChangeFirst("muhammad"))
	assert.Equal(t, "hello rafli", data.SayHello("rafli"))
}
