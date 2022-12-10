package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseShould(t *testing.T) {
	t.Run("work for empty slice", func(t *testing.T) {
		actual := []int{}
		Reverse(actual)
		assert.Equal(t, []int{}, actual)
	})

	t.Run("work for slice with one element", func(t *testing.T) {
		actual := []int{1}
		Reverse(actual)
		assert.Equal(t, []int{1}, actual)
	})

	t.Run("work for slice with multiple elements", func(t *testing.T) {
		actual := []int{1, 2, 3}
		Reverse(actual)
		assert.Equal(t, []int{3, 2, 1}, actual)
	})

}
