package util

import (
	"math"
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

func TestMaxShould(t *testing.T) {
	t.Run("work with integer values", func(t *testing.T) {
		assert.Equal(t, 2, Max(1, 2))
	})

	t.Run("work with float values", func(t *testing.T) {
		assert.Equal(t, 2.2, Max(1.1, 2.2))
	})

	t.Run("work with infinite values", func(t *testing.T) {
		assert.Equal(t, math.Inf(0), Max(math.Inf(0), 2))
	})

	t.Run("work with NaN values", func(t *testing.T) {
		assert.Equal(t, 23.0, Max(math.NaN(), 23.0))
	})
}

func TestMinShould(t *testing.T) {
	t.Run("work with integer values", func(t *testing.T) {
		assert.Equal(t, 1, Min(1, 2))
	})

	t.Run("work with float values", func(t *testing.T) {
		assert.Equal(t, 1.1, Min(1.1, 2.2))
	})

	t.Run("work with infinite values", func(t *testing.T) {
		assert.Equal(t, math.Inf(-1), Min(math.Inf(-1), 2))
	})

	t.Run("work with NaN values", func(t *testing.T) {
		assert.Equal(t, 23.0, Min(math.NaN(), 23.0))
	})
}

func TestSetShould(t *testing.T) {
	t.Run("work with integer values", func(t *testing.T) {
		s := NewSet(1, 2)
		s.Add(4)

		assert.True(t, s.Contains(1))
		assert.True(t, s.Contains(2))
		assert.True(t, s.Contains(4))
		assert.False(t, s.Contains(3))
	})

	t.Run("work with string values", func(t *testing.T) {
		s := NewSet("a", "b")

		assert.ElementsMatch(t, []string{"a", "b"}, s.Members())
	})

	t.Run("work with rune values", func(t *testing.T) {
		s := NewSet('a', 'b')
		s.Remove('a')
		s.Remove('c')

		assert.False(t, s.Contains('a'))
		assert.True(t, s.Contains('b'))
	})
}
