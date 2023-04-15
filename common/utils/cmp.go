package utils

import "golang.org/x/exp/constraints"

// Max returns the larger of x or y with generics.
func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Min returns the smaller of x or y with generics.
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
