package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("positive size", func(t *testing.T) {
		size := 10
		data := generateRandomElements(size)

		assert.Len(t, data, size, "длина слайса должна совпадать с запрошенным размером")

		for _, v := range data {
			assert.Positive(t, v, "каждое число должно быть положительным")
		}
	})

	t.Run("zero size", func(t *testing.T) {
		data := generateRandomElements(0)

		assert.Empty(t, data, "при нулевом размере должен вернуться пустой слайс")
	})

	t.Run("negative size", func(t *testing.T) {
		data := generateRandomElements(-5)

		assert.Empty(t, data, "при отрицательном размере должен вернуться пустой слайс")
	})
}

func TestMaximum(t *testing.T) {
	t.Run("regular slice", func(t *testing.T) {
		data := []int{3, 7, 2, 9, 4}
		want := 9

		got := maximum(data)
		assert.Equal(t, want, got)
	})

	t.Run("single element", func(t *testing.T) {
		data := []int{42}
		want := 42

		got := maximum(data)
		assert.Equal(t, want, got)
	})

	t.Run("empty slice", func(t *testing.T) {
		data := []int{}
		want := 0

		got := maximum(data)
		assert.Equal(t, want, got)
	})

	t.Run("all equal elements", func(t *testing.T) {
		data := []int{5, 5, 5, 5}
		want := 5

		got := maximum(data)
		assert.Equal(t, want, got)
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("regular slice", func(t *testing.T) {
		data := []int{3, 7, 2, 9, 4, 1, 8, 6, 5, 10, 12, 11, 0, 13, 14, 15}
		want := 15

		got := maxChunks(data)
		assert.Equal(t, want, got)
	})

	t.Run("empty slice", func(t *testing.T) {
		data := []int{}
		want := 0

		got := maxChunks(data)
		assert.Equal(t, want, got)
	})

	t.Run("matches single-threaded result", func(t *testing.T) {
		data := generateRandomElements(1000)

		want := maximum(data)
		got := maxChunks(data)

		assert.Equal(t, want, got, "результат многопоточного поиска должен совпадать с однопоточным")
	})
}
