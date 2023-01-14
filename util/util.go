package util

import (
	"golang.org/x/exp/constraints"
)

// Reverse reverses a slice in place
func Reverse[E any](s []E) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Max returns the largest of x or y.
func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Min returns the smallest of x or y.
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
