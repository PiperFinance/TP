package utils

func Chunks[T any](array []T, chunkSize int) [][]T {
	steps := len(array) / chunkSize

	c := make([][]T, steps)
	for i := 0; i < steps; i++ {
		c[i] = array[(i * chunkSize):((i + 1) * chunkSize)]
	}
	return c
}
