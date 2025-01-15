package util

func Map[T any, U any](s []T, f func(T) U) []U {
	s1 := make([]U, len(s))
	for i, t := range s {
		s1[i] = f(t)
	}
	return s1
}
