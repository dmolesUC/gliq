package util

import "slices"

func Map[T any, U any](s []T, f func(T) U) []U {
	s1 := make([]U, len(s))
	for i, t := range s {
		s1[i] = f(t)
	}
	return s1
}

func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	s2 := S(nil)
	for _, v := range s {
		if f(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

func Intersect[S1 ~[]E, S2 ~[]E, E comparable](s1 S1, s2 S2) S1 {
	return Filter(s1, func(e E) bool {
		return slices.Contains(s2, e)
	})
}
