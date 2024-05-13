/*
typesafe functional programming utils in go cuz i'm crazy
*/
package utils

func Map[T any, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[T any, R any](s []T, initial R, f func(R, T) R) R {
	result := initial
	for _, v := range s {
		result = f(result, v)
	}
	return result
}
