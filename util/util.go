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

// Abs returns the absolute value of x.
func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Set represents a set structure
// Inspired from https://bitfieldconsulting.com/posts/generic-set
type Set[E comparable] map[E]struct{}

// NewSet returns a new set
func NewSet[E comparable](values ...E) Set[E] {
	s := Set[E]{}
	for _, v := range values {
		s[v] = struct{}{}
	}
	return s
}

// Add adds an element to the set
func (s Set[E]) Add(values ...E) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

// Remove removes an element from the set
func (s Set[E]) Remove(values ...E) {
	for _, v := range values {
		delete(s, v)
	}
}

// Contains returns true if the value is in the set
func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

// Members returns all the members of the set in a slice
func (s Set[E]) Members() []E {
	result := make([]E, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}
